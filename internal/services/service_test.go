package services

import "testing"

func TestArrayToString(t *testing.T) {
	tests := []struct {
		input    []string
		expected string
	}{
		{[]string{"banana", "maçã", "laranja"}, "banana,maçã,laranja"},
		{[]string{"a", "b", "c"}, "a,b,c"},
		{[]string{}, ""},
		{[]string{"um"}, "um"},
	}

	for _, test := range tests {
		result := ArrayToString(test.input)
		if result != test.expected {
			t.Errorf("arrayToString(%v) = %v; want %v", test.input, result, test.expected)
		}
	}
}
