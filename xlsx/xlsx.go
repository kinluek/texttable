package xlsx

import (
	"archive/zip"
	"bytes"
	"encoding/xml"
	"errors"
	"fmt"
	"github.com/kinluek/texttable"
	"math"
	"os"
	"regexp"
	"strconv"
)

var (
	errMissingContentTypes  = errors.New("missing content types file")
	errMissingSharedStrings = errors.New("missing shared strings file")
)

const (
	fileTypeSharedStrings = "application/vnd.openxmlformats-officedocument.spreadsheetml.sharedStrings+xml"
	fileTypeWorkSheet     = "application/vnd.openxmlformats-officedocument.spreadsheetml.worksheet+xml"

	fileNameContentTypes = "[Content_Types].xml"
)

// WorkSheetExtract contains a formatted work sheet string
// and its corresponding sheet number.
type WorkSheetExtract struct {
	SheetName string
	Text      string
}

// Config should be used to specify output table
// formatting of the extracted XLSX text content.
type Config struct {
	ColumnWidth  int
	ColumnMargin int
	RowMargin    int
}

// Extract takes an *os.File which should contain the zipped
// XLSX data and returns the extracted content as a formatted
// string table. It takes a second parameter which sets the column
// width and column and row margins of the generated string tables.
func Extract(file *os.File, config Config) ([]WorkSheetExtract, error) {
	fi, err := file.Stat()
	if err != nil {
		return nil, err
	}
	zr, err := zip.NewReader(file, fi.Size())
	if err != nil {
		return nil, err
	}

	zipFiles := mapZipFiles(zr.File)
	descFile, ok := zipFiles[fileNameContentTypes]
	if !ok {
		return nil, errMissingContentTypes
	}
	var fileDesc ContentTypes
	if err := decodeZipFile(descFile, &fileDesc); err != nil {
		return nil, err
	}

	var sharedStrings SharedStrings

	for _, or := range fileDesc.Overrides {
		switch or.ContentType {
		case fileTypeSharedStrings:
			sharedStringsFile, ok := zipFiles[or.PartName]
			if !ok {
				return nil, errMissingSharedStrings
			}
			if err := decodeZipFile(sharedStringsFile, &sharedStrings); err != nil {
				return nil, fmt.Errorf("could not decode file %v: %v", or.PartName, err)
			}
			break
		}
	}

	workSheetExtracts := make([]WorkSheetExtract, 0)
	stringLookup := MakeStringLookup(sharedStrings)

	for _, or := range fileDesc.Overrides {
		switch or.ContentType {
		case fileTypeWorkSheet:
			workSheet, ok := zipFiles[or.PartName]
			if !ok {
				return nil, fmt.Errorf("missing work sheet %v", or.PartName)
			}
			var sheet WorkSheet
			if err := decodeZipFile(workSheet, &sheet); err != nil {
				return nil, fmt.Errorf("could not decode file %v: %v", or.PartName, err)
			}

			textMatrix, err := MakeTextMatrix(sheet, stringLookup)
			if err != nil {
				return nil, fmt.Errorf("could not create text table: %v", err)
			}

			ttf := texttable.New(textMatrix, texttable.Config{
				ColumnMargin: config.ColumnMargin,
				ColumnWidth:  config.ColumnWidth,
				RowMargin:    config.RowMargin,
			})
			stringTable, err := ttf.Output()
			if err != nil {
				return nil, fmt.Errorf("could not format text table into string: %v", err)
			}

			workSheetExtracts = append(workSheetExtracts, WorkSheetExtract{
				SheetName: workSheet.Name,
				Text:      stringTable,
			})
		}
	}

	return workSheetExtracts, nil
}

// SharedStringLookup is a slice of strings, it should be created
// from a SharedStrings struct.
type SharedStringLookup []string

// MakeStringLookup creates a SharedStringLookup from the
// shared strings XML struct.
func MakeStringLookup(sharedStrings SharedStrings) SharedStringLookup {
	lookupSlice := make(SharedStringLookup, len(sharedStrings.StringItem))
	for i, si := range sharedStrings.StringItem {
		if len(si.RichTextRuns) > 0 {
			text := ""
			for _, run := range si.RichTextRuns {
				text += run.Text.Text
			}
			lookupSlice[i] += text
		} else {
			lookupSlice[i] = si.Text.Text
		}
	}
	return lookupSlice
}

// MakeTextMatrix takes a WorkSheet and a SharedStringLookup
// to create a 2D string slice from. The 2D slice will be populated
// according to the layout of the WorkSheet.
func MakeTextMatrix(sheet WorkSheet, lookup SharedStringLookup) ([][]string, error) {
	rows := sheet.SheetData.Rows

	width, height, err := GetSheetSize(sheet)
	if err != nil {
		return nil, err
	}
	textTable := makeStringMatrix(width, height)
	for _, row := range rows {
		for _, cell := range row.Cells {
			x, y, err := parseXYCoordinate(cell.Coordinate)
			if err != nil {
				return nil, fmt.Errorf("could not parse x y index: %v", err)
			}

			var text string
			if cell.Type == "s" {
				stringIdx, err := strconv.Atoi(cell.Value)
				if err != nil {
					return nil, fmt.Errorf("could not parse string index %v: %v", cell.Value, err)
				}
				text = lookup[stringIdx]
			} else {
				text = cell.Value
			}
			textTable[y][x] = text
		}
	}

	return textTable, nil
}

// GetSheetSize takes sheetData and returns
// the max width and height of the spread sheet
// cells.
func GetSheetSize(workSheet WorkSheet) (width, height int, err error) {
	workSheetRows := workSheet.SheetData.Rows

	// get height
	lastRow := workSheetRows[len(workSheetRows)-1]
	pos := lastRow.Cells[0].Coordinate

	rowIndexRegex := regexp.MustCompile(`\d+`)
	matched := rowIndexRegex.Find([]byte(pos))
	height, err = strconv.Atoi(string(matched))
	if err != nil {
		return 0, 0, fmt.Errorf("could not parse height: %v", err)
	}

	// get width
	colIndexRegex := regexp.MustCompile(`[a-zA-Z]+`)
	maxColIndex := 0

	for _, row := range workSheetRows {
		if rowLength := len(row.Cells); rowLength > 0 {
			lastCellPos := row.Cells[rowLength-1].Coordinate
			matched := colIndexRegex.Find([]byte(lastCellPos))

			colIndex := alphaIndex(matched)
			if colIndex > maxColIndex {
				maxColIndex = colIndex
			}
		}
	}

	return maxColIndex + 1, height, nil
}

// parseXYCoordinate takes a XLSX coordinate of the form:
// AADD, where AA are alphabetical characters and DD are digits.
// and parses the corresponding int x, y indexes.
// Example: A1 -> (0, 0) C10 -> (2, 9)
func parseXYCoordinate(coordinate string) (x, y int, err error) {
	xyRegex := regexp.MustCompile(`([a-zA-Z]+)(\d+)`)
	matches := xyRegex.FindSubmatch([]byte(coordinate))
	if len(matches) != 3 {
		return 0, 0, fmt.Errorf("invalid coordinate: %v", coordinate)
	}
	x = alphaIndex(matches[1])
	y, err = strconv.Atoi(string(matches[2]))
	if err != nil {
		return 0, 0, err
	}
	y = y - 1
	return x, y, nil
}

// alphaPositions should be used to set the position values
// of the alphabetical characters
var alphaPositions = []byte("ABCDEFGHIJKLMNOPQRSTUVWXYZ")

// alphaIndex takes an alphabetical index (used as the horizontal
// index of an XLSX spread sheet) and converts it to the numerical
// index.
// Example:
// * A -> 0
// * C -> 2
// * AB -> 27
// * BDP -> 1471
func alphaIndex(charIdx []byte) int {
	index := 0
	powerFactor := 0

	for i := len(charIdx) - 1; i >= 0; i-- {
		multiplier := bytes.Index(alphaPositions, []byte{charIdx[i]}) + 1
		columnOffset := int(math.Pow(26, float64(powerFactor)))
		offset := multiplier * columnOffset
		index += offset
		powerFactor++
	}

	return index - 1
}

// makeStringMatrix creates a 2D string array using
// a given width and height.
func makeStringMatrix(width, height int) [][]string {
	matrix := make([][]string, height)
	for i := range matrix {
		matrix[i] = make([]string, width)
	}
	return matrix
}

// decodeZipFile takes a zip file and opens it and
// reads it into the given data structure.
func decodeZipFile(f *zip.File, v interface{}) error {
	ff, err := f.Open()
	if err != nil {
		return err
	}
	defer ff.Close()
	dec := xml.NewDecoder(ff)
	if err := dec.Decode(v); err != nil {
		return err
	}
	return nil
}

// mapZipFiles takes a slice of *zip.Files and turns
// it into a map, using the file name as the key.
func mapZipFiles(files []*zip.File) map[string]*zip.File {
	zipMap := make(map[string]*zip.File)
	for _, f := range files {
		zipMap["/"+f.Name] = f
		zipMap[f.Name] = f
	}
	return zipMap
}
