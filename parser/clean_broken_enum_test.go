package parser

import (
	"testing"
)

func TestFormatOneOfEnum(t *testing.T) {
	tests := []struct {
		input    string
		expected string
	}{
		{"one of (`,left,right`)", "one of (left, right)"},
		{"one of (`US`, `UK`, `NZ`)", "one of (`US`, `UK`, `NZ`)"},
		{"Values are one of (`,apple,banana,grape`)", "Values are one of (apple, banana, grape)"},
		{"one of (no change here)", "one of (no change here)"},
	}

	for _, test := range tests {
		output := FormatOneOfEnum(test.input)
		if output != test.expected {
			t.Errorf("For input '%s', expected '%s' but got '%s'", test.input, test.expected, output)
		}
	}
}
