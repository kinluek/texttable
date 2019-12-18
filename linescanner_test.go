package texttable

import (
	"io"
	"testing"
)

func TestLineScanner_Next(t *testing.T) {

	tests := []struct {
		name  string
		input string
		want  string
	}{
		{
			name:  "simple sentence",
			input: "hello there, today is a tuesday, not many people are in work today. It's almost Christmas! Thank god for holidays!",
			want: `  hello there, today is a    
  tuesday, not many people   
  are in work today. It's    
  almost Christmas! Thank    
  god for holidays!          
`,
		},
	}

	for _, tt := range tests {
		t.Run(tt.name, func(t *testing.T) {
			formattedText := ""
			s := NewLineScanner(tt.input, LineScannerConfig{25, 2})
			for {
				line, err := s.Next()
				if err == io.EOF {
					break
				}
				formattedText += line + "\n"
			}
			if formattedText != tt.want {
				t.Fatalf("expected: \n%v\n\ngot: \n%v\n", tt.want, formattedText)
			}
		})
	}
}
