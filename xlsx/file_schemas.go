package xlsx

import "encoding/xml"

// ContentTypes represents a Content-Type
// XML file, it contains the file descriptions.
type ContentTypes struct {
	XMLName   xml.Name   `xml:"Types"`
	Text      string     `xml:",chardata"`
	Xmlns     string     `xml:"xmlns,attr"`
	Overrides []override `xml:"Override"`
}

// override contains contains information
// on the content type of a file, identified
// by the PartName.
type override struct {
	XMLName     xml.Name `xml:"Override"`
	ContentType string   `xml:"ContentType,attr"`
	PartName    string   `xml:"PartName,attr"`
}

// WorkSheet represents an XLSX worksheet.
type WorkSheet struct {
	XMLName   xml.Name  `xml:"worksheet"`
	Text      string    `xml:",chardata"`
	Xmlns     string    `xml:"xmlns,attr"`
	R         string    `xml:"r,attr"`
	SheetData sheetData `xml:"sheetData"`
}

// sheetData represents the sheet data of
// an XLSX worksheet.
type sheetData struct {
	XMLName xml.Name `xml:"sheetData"`
	Text    string   `xml:",chardata"`
	Rows    []row    `xml:"row"`
}

// row represents a worksheet row.
type row struct {
	Text         string `xml:",chardata"`
	R            string `xml:"r,attr"`
	Spans        string `xml:"spans,attr"`
	Ht           string `xml:"ht,attr"`
	ThickBot     string `xml:"thickBot,attr"`
	CustomHeight string `xml:"customHeight,attr"`
	Cells        []cell `xml:"c"`
}

// cell represents a worksheet cell.
type cell struct {
	Text       string `xml:",chardata"`
	Coordinate string `xml:"r,attr"`
	S          string `xml:"s,attr"`
	Type       string `xml:"t,attr"`
	Value      string `xml:"v"`
	F          string `xml:"f"`
}

// SharedStrings represents a shared strings
// XML file for an XLSX document.
type SharedStrings struct {
	XMLName     xml.Name     `xml:"sst"`
	Text        string       `xml:",chardata"`
	Xmlns       string       `xml:"xmlns,attr"`
	Count       string       `xml:"count,attr"`
	UniqueCount string       `xml:"uniqueCount,attr"`
	StringItem  []stringItem `xml:"si"`
}

// stringItem contains the shared string
// text element
type stringItem struct {
	Text         ssText         `xml:"t"`
	RichTextRuns []richTextRuns `xml:"r"`
}

// ssText represents a shared string
// text element contained in a shared strings
// XML file.
type ssText struct {
	Text  string `xml:",chardata"`
	Space string `xml:"space,attr"`
}

// richTextRuns represents a rich text run
// element use to hold text with with formatting
// at the character level.
type richTextRuns struct {
	Text ssText `xml:"t"`
}
