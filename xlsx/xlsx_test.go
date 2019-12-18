package xlsx

import (
	"os"
	"testing"
)

func TestAlphaIndex(t *testing.T) {
	tests := []struct {
		name  string
		input string
		want  int
	}{
		{
			name:  "single letter index 1",
			input: "A",
			want:  0,
		},
		{
			name:  "single letter index 2",
			input: "Z",
			want:  25,
		},
		{
			name:  "double letter index 1",
			input: "AA",
			want:  26,
		},
		{
			name:  "double letter index 2",
			input: "ZZ",
			want:  701,
		},
		{
			name:  "triple letter index 1",
			input: "AAA",
			want:  702,
		},
		{
			name:  "triple letter index 2",
			input: "AAZ",
			want:  727,
		},
		{
			name:  "triple letter index 3",
			input: "ABA",
			want:  728,
		},
		{
			name:  "triple letter index 4",
			input: "BDP",
			want:  1471,
		},
		{
			name:  "triple letter index 5",
			input: "ZJK",
			want:  17846,
		},
		{
			name:  "triple letter index 6",
			input: "ZZZ",
			want:  18277,
		},
	}

	for _, tt := range tests {
		t.Run(tt.name, func(t *testing.T) {
			index := alphaIndex([]byte(tt.input))
			if index != tt.want {
				t.Fatalf("Test: \"%v\" failed: extected: %v, got: %v", tt.name, tt.want, index)
			}
		})
	}
}

func TestExtract(t *testing.T) {
	tests := []struct {
		name       string
		inputPath  string
		outputPath string
	}{
		{
			name:       "simple spread sheet 1",
			inputPath:  "./testdata/xlsx_files/simple_spread_sheet_1/simple_spread_sheet_1.xlsx",
			outputPath: "./testdata/xlsx_files/simple_spread_sheet_1/simple_spread_sheet_1.txt",
		},
		{
			name:       "simple spread sheet 2",
			inputPath:  "./testdata/xlsx_files/simple_spread_sheet_2/simple_spread_sheet_2.xlsx",
			outputPath: "./testdata/xlsx_files/simple_spread_sheet_2/simple_spread_sheet_2.txt",
		},
		{
			name:       "two page spread sheet 1",
			inputPath:  "./testdata/xlsx_files/two_page_spread_sheet_1/two_page_spread_sheet_1.xlsx",
			outputPath: "./testdata/xlsx_files/two_page_spread_sheet_1/two_page_spread_sheet_1.txt",
		},
	}

	for _, tt := range tests {
		t.Run(tt.name, func(t *testing.T) {
			f, err := os.Open(tt.inputPath)
			if err != nil {
				t.Fatalf("could not open input file: %v", err)
			}

			sheets, err := Extract(f, Config{
				ColumnWidth:  30,
				ColumnMargin: 2,
				RowMargin:    1,
			})
			if err != nil {
				t.Fatalf("could not extract text: %v", err)
			}

			text := ""

			for _, sheet := range sheets {
				text += sheet.SheetName + "\n\n"
				text += sheet.Text + "\n"
			}

			nf, err := os.Create(tt.outputPath)
			if err != nil {
				t.Fatalf("could not create output file: %v", err)
			}

			_, err = nf.Write([]byte(text))
			if err != nil {
				t.Fatalf("could not write to output file: %v", err)
			}

			f.Close()
			nf.Close()
		})
	}

}
