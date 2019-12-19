package texttable

import (
	"io"
	"strings"
)

// TextTable knows how to format a 2D string slice
// into a constant width string table. It should be constructed
// only with the New function.
type TextTable struct {
	scannerMatrix     [][]*LineScanner
	emptyColumnFiller string
	rowSpacing        string
}

// Config should be used to set up configuration
// for a TextTable.
type Config struct {
	ColumnWidth    int
	ColumnMargin   int
	RowMargin      int
	IgnoreNewLines bool
}

// New knows how to create a new TextTable
// from a given text table and optional Config parameters.
// Note: the minimum values for column width, column margin and row
// margin are 1, 0 and 0 respectively.
func New(textTable [][]string, config Config) *TextTable {
	// adjust config to minimum threshold
	// if any parameters go below.
	if config.ColumnWidth < 1 {
		config.ColumnWidth = 1
	}
	if config.ColumnMargin < 0 {
		config.ColumnMargin = 0
	}
	if config.RowMargin < 0 {
		config.RowMargin = 0
	}

	// ensure we have a constant width
	// table by padding any rows with
	// less elements than the max row length.
	var maxRowLength int
	for _, row := range textTable {
		if width := len(row); width > maxRowLength {
			maxRowLength = width
		}
	}
	for i := range textTable {
		diff := maxRowLength - len(textTable[i])
		if diff > 0 {
			extension := make([]string, diff)
			textTable[i] = append(textTable[i], extension...)
		}
	}

	// create line scanner matrix
	height := len(textTable)
	width := maxRowLength
	scannerMatrix := make([][]*LineScanner, height)
	for i := range scannerMatrix {
		scannerMatrix[i] = make([]*LineScanner, width)
	}
	for y := range textTable {
		for x := range textTable[y] {
			scannerMatrix[y][x] = NewLineScanner(textTable[y][x], LineScannerConfig{
				LineWidth:      config.ColumnWidth,
				LineMargin:     config.ColumnMargin,
				IgnoreNewLines: config.IgnoreNewLines,
			})
		}
	}

	return &TextTable{
		scannerMatrix:     scannerMatrix,
		rowSpacing:        strings.Repeat("\n", config.RowMargin),
		emptyColumnFiller: strings.Repeat(" ", config.ColumnWidth+2*config.ColumnMargin),
	}
}

// Output produces the formatted text table as a string.
func (tf *TextTable) Output() (string, error) {
	stringTable := ""

	numRows := len(tf.scannerMatrix)

	for n, scannerRow := range tf.scannerMatrix {
		rowLength := len(scannerRow)
		eofCount := 0
		for i := 0; i < rowLength; i++ {
			line, err := scannerRow[i].Next()
			if err != nil && err != io.EOF {
				return "", err
			}
			if err == io.EOF {
				stringTable += tf.emptyColumnFiller
				eofCount++
			} else {
				stringTable += line
			}
			if eofCount == rowLength {
				break
			}
			if i == rowLength-1 {
				eofCount = 0
				i = -1
				stringTable += "\n"
			}
		}
		if n != numRows-1 {
			stringTable += tf.rowSpacing
		}
	}
	return stringTable, nil
}
