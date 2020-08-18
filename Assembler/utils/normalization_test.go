package utils

import "testing"

func TestLineNormalization(t *testing.T) {
	var tests = []struct {
		param    string
		expected string
	}{
		{"MOV 8", "mov 8"},
		{"MoV 8", "mov 8"},
		{"	 Mov 8	 ", "mov 8"},
		{"Mov 	acc  	8", "mov acc 8"},
	}

	for _, test := range tests {
		got := LineNormalization(test.param)
		if got != test.expected {
			t.Errorf("Expected: '%s', Got: '%s'", test.expected, got)
		}
	}
}

func TestRemoveUnecessarySpaces(t *testing.T) {
	var tests = []struct {
		param    string
		expected string
	}{
		{"	 mov 8	 ", "mov 8"},
		{"mov 	acc  	8", "mov acc 8"},
		{" 	mov 	acc  	8", "mov acc 8"},
	}

	for _, test := range tests {
		got := removeUnecessarySpaces(test.param)

		if got != test.expected {
			t.Errorf("Expected: '%s', Got: '%s'", test.expected, got)
		}
	}
}
