package texttable

import (
	"bufio"
	"fmt"
	"io"
	"strconv"
	"strings"
	"unicode/utf8"
)

// LineScanner knows how to read a string, line by line
// for a given line width. The lines scanned can only be
// separated by white space, therefore if a word causes the
// line to go over the line width, it will be added to the
// next line. Padding with white space is also taken care
// of, to make sure each line is of equal length.
//
// LineScanner should only be constructed with the NewLineScanner
// function.
type LineScanner struct {
	s             *bufio.Scanner
	lineWidth     int
	lineFormat    string
	overFlow      string
	newLineReturn string
}

// LineScannerConfig should be used to set optional
// configuration parameters for the LineScanner.
type LineScannerConfig struct {
	LineWidth  int
	LineMargin int
}

// NewLineScanner knows how to create a new LineScanner
// from some text and an option LineScannerConfig. A config
// should be passed to set the line width and margin.
// Note: the minimum values for line width and margin are
// 1 and 0, respectively.
func NewLineScanner(text string, config LineScannerConfig) *LineScanner {
	// adjust config to minimum threshold
	// if any parameters go below.
	if config.LineWidth < 1 {
		config.LineWidth = 1
	}
	if config.LineMargin < 0 {
		config.LineMargin = 0
	}

	s := bufio.NewScanner(strings.NewReader(text))
	newLineReturn := []byte(strings.Repeat(" ", config.LineWidth))
	s.Split(scanWordsAndNewLines(newLineReturn))

	margin := strings.Repeat(" ", config.LineMargin)
	lineFormat := margin + "%-" + strconv.Itoa(config.LineWidth) + "v" + margin

	return &LineScanner{
		s:          s,
		lineWidth:  config.LineWidth,
		lineFormat: lineFormat,
		overFlow:   "",
		newLineReturn: string(newLineReturn),
	}
}

// Next returns the next line in the text.
// An io.EOF error is returned to signify that all
// the text has been read and returned.
// If a single word is scanned that exceeds the
// configured line width, the word will be split
// across multiple Next calls.
func (ls *LineScanner) Next() (string, error) {
	line := ls.overFlow

	// split overflown words that exceed
	// line width over multiple lines
	if len(ls.overFlow) > ls.lineWidth {
		newLine := ls.overFlow[:ls.lineWidth]
		ls.overFlow = ls.overFlow[ls.lineWidth:]
		return fmt.Sprintf(ls.lineFormat, newLine), nil
	}
	ls.overFlow = ""

	for ls.s.Scan() {
		if err := ls.s.Err(); err != nil {
			return "", fmt.Errorf("error scanning line: %v", err)
		}

		word := ls.s.Text()
		// do not append straight to global line
		// variable, as we need to be able to backtrack
		// if line width is exceeded
		var newLine string
		if line == "" {
			newLine = word
		} else {
			newLine = line + " " + word
		}
		// split words that exceed line width
		// across multiple lines
		if len(word) > ls.lineWidth {
			ls.overFlow = newLine[ls.lineWidth:]
			return fmt.Sprintf(ls.lineFormat, newLine[:ls.lineWidth]), nil
		}
		// handle new lines
		if word == ls.newLineReturn {
			if len(newLine) > 0 {
				return fmt.Sprintf(ls.lineFormat, line), nil
			}
			return fmt.Sprintf(ls.lineFormat, word), nil
		}

		if len(newLine) > ls.lineWidth {
			ls.overFlow = word
			return fmt.Sprintf(ls.lineFormat, line), nil
		}
		// if line width was not exceeded,
		// set the new line as the global line.
		line = newLine
	}
	if line == "" {
		return "", io.EOF
	}
	return fmt.Sprintf(ls.lineFormat, line), nil
}

// scanWordsAndNewLines returns a split function for a Scanner that
// returns each space-separated word of text, with surrounding spaces
// deleted. It will never return an empty string. The definition of
// space is set by unicode.IsSpace.
//
// This is an extension of Go's standard library bufio.ScanWords
// function, where new lines ('\n') are not ignored by the scanner
// and returns the newLineReturn argument in its place.
// This function is crucial to getting new lines to be wrapped correctly
// when formatting text tables.
func scanWordsAndNewLines(newLineReturn []byte) bufio.SplitFunc {
	return func(data []byte, atEOF bool) (advance int, token []byte, err error) {
		start := 0
		for width := 0; start < len(data); start += width {
			var r rune
			r, width = utf8.DecodeRune(data[start:])
			if r == '\n' {
				return start + width, newLineReturn, nil
			}
			if !isSpace(r) {
				break
			}
		}
		for width, i := 0, start; i < len(data); i += width {
			var r rune
			r, width = utf8.DecodeRune(data[i:])
			if isSpace(r) || r == '\n' {
				return i + width - 1, data[start:i], nil
			}
		}
		// If we're at EOF, we have a final, non-empty, non-terminated word. Return it.
		if atEOF && len(data) > start {
			return len(data), data[start:], nil
		}
		// Request more data.
		return start, nil, nil
	}
}

// isSpace reports whether the character is a Unicode white space character.
// We avoid dependency on the unicode package, but check validity of the implementation
// in the tests.
//
// This is taken from the standard library bufio.isSpace function, with the new line ('\n')
// rune removed from being considered as a space.
func isSpace(r rune) bool {
	if r <= '\u00FF' {
		// Obvious ASCII ones: \t through \r plus space. Plus two Latin-1 oddballs.
		switch r {
		case ' ', '\t', '\v', '\f', '\r':
			return true
		case '\u0085', '\u00A0':
			return true
		}
		return false
	}
	// High-valued ones.
	if '\u2000' <= r && r <= '\u200a' {
		return true
	}
	switch r {
	case '\u1680', '\u2028', '\u2029', '\u202f', '\u205f', '\u3000':
		return true
	}
	return false
}
