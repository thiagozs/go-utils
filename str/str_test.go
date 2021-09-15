package str

import (
	"testing"
)

func TestTrimBreakLines(t *testing.T) {

	tests := []struct {
		value    string
		expected string
	}{
		{
			value:    "\n\n\ntest",
			expected: "test",
		},
		{
			value:    "\t\ntest\t\n",
			expected: "test",
		},
		{
			value:    "        test",
			expected: "test",
		},
	}

	for _, test := range tests {

		rr := TrimBreakLines(test.value)

		if rr != test.expected {
			t.Errorf("Trim break lines: result '%s' expected '%s'\n", test.value, test.expected)
		}
	}

}

func TestFloat642Str(t *testing.T) {

	tests := []struct {
		value    float64
		expected string
	}{
		{
			value:    1,
			expected: "1.00",
		},
		{
			value:    1.444,
			expected: "1.44",
		},
		{
			value:    -1.00,
			expected: "-1.00",
		},
	}

	for _, test := range tests {

		rr := Float642Str(test.value)

		if rr != test.expected {
			t.Errorf("Float64 to string: result '%f' expected '%s'\n", test.value, test.expected)
		}
	}

}

func TestFloat642StrPrec(t *testing.T) {

	tests := []struct {
		value     float64
		precision int
		expected  string
	}{
		{
			value:     1,
			precision: 8,
			expected:  "1.00000000",
		},
		{
			value:     1.444,
			precision: 3,
			expected:  "1.444",
		},
		{
			value:     -1.00,
			precision: 2,
			expected:  "-1.00",
		},
		{
			value:     -1.500000000,
			precision: 1,
			expected:  "-1.5",
		},
	}

	for _, test := range tests {

		rr := Float642StrPrec(test.value, test.precision)

		if rr != test.expected {
			t.Errorf("Float64 to string precision: result '%f' expected '%s'\n", test.value, test.expected)
		}
	}

}

func TestSubStr(t *testing.T) {

	tests := []struct {
		value    string
		substr   string
		expected string
	}{
		{
			value:    "The quick brown fox jumps over the lazy dog",
			substr:   "the lazy dog",
			expected: "the lazy dog",
		},
		{
			value:    "The quick brown fox jumps over the lazy dog",
			substr:   "jumps",
			expected: "jumps",
		},
		{
			value:    "The quick brown fox jumps over the lazy dog",
			substr:   "brown fox jumps",
			expected: "brown fox jumps",
		},
	}

	for _, test := range tests {

		rr := SubStr(test.value, test.substr)

		if rr != test.expected {
			t.Errorf("SubString: result '%s' expected '%s'\n", test.value, test.expected)
		}
	}

}
