package parser

import (
	"bufio"
	"strings"
	"testing"
	"time"
)

func TestReadLine(t *testing.T) {
	tests := []struct {
		input        string
		expectedLine string
		expectedEOF  bool
	}{
		{
			input:        "line1\nline2\nline3\n",
			expectedLine: "line1",
			expectedEOF:  false,
		},
		{
			input:        "line1\n",
			expectedLine: "line1",
			expectedEOF:  false,
		},
		{
			input:        "",
			expectedLine: "",
			expectedEOF:  true,
		},
	}

	for _, test := range tests {
		reader := bufio.NewReader(strings.NewReader(test.input))
		line, eof := ReadLine(reader)
		if line != test.expectedLine {
			t.Errorf("expected line %q, got %q", test.expectedLine, line)
		}
		if eof != test.expectedEOF {
			t.Errorf("expected EOF %v, got %v", test.expectedEOF, eof)
		}
	}
}

func TestReadLineEOF(t *testing.T) {
	input := "line1\n"
	reader := bufio.NewReader(strings.NewReader(input))

	// Read first line
	line, eof := ReadLine(reader)
	if line != "line1" || eof {
		t.Errorf("expected line1 and false, got %q and %v", line, eof)
	}

	// Try to read past the end
	line, eof = ReadLine(reader)
	if line != "" || !eof {
		t.Errorf("expected empty string and true, got %q and %v", line, eof)
	}
}

func TestParseLine(t *testing.T) {
	TFMT := time.RFC3339
	test := struct {
		line         string
		expectedTime time.Time
		expectedNum  float64
	}{
		line:         "2021-03-04T23:03:26Z 107.9864",
		expectedTime: time.Date(2021, 3, 4, 23, 3, 26, 0, time.UTC),
		expectedNum:  107.9864,
	}

	gotTime, gotNum := ParseLine(test.line, TFMT)
	if !gotTime.Equal(test.expectedTime) {
		t.Errorf("expected time %v, got %v", test.expectedTime, gotTime)
	}
	if gotNum != test.expectedNum {
		t.Errorf("expected number %f, got %f", test.expectedNum, gotNum)
	}

}
