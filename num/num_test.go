package num

import "testing"

func TestStr2Float64(t *testing.T) {

	tests := []struct {
		value    string
		expected float64
	}{
		{
			value:    "1",
			expected: 1.00,
		},
		{
			value:    "1.444",
			expected: 1.444,
		},
		{
			value:    "-1.00",
			expected: -1.00,
		},
		{
			value:    "-1.500000000",
			expected: -1.500000000,
		},
	}

	for _, test := range tests {

		rr, err := Str2Float64(test.value)
		if err != nil {
			t.Error(err)
		}

		if rr != test.expected {
			t.Errorf("String to Float64: result '%s' expected '%f'\n", test.value, test.expected)
		}
	}

}
