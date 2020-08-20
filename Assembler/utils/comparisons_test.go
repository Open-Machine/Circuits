package utils

import "testing"

func TestSafeStrPointerComparation(t *testing.T) {
	var tests = []struct {
		param1   *string
		param2   *string
		expected bool
	}{
		{nil, nil, true},
		{NewString("Hello"), nil, false},
		{nil, NewString("Hello"), false},
		{NewString("Hello"), NewString("Hello"), true},
		{NewString("Hello"), NewString("Hella"), false},
	}

	for i, test := range tests {
		got := SafeIsEqualStrPointer(test.param1, test.param2)
		if got != test.expected {
			t.Errorf("[%d] Expected: %t, Got: %t", i, test.expected, got)
		}
	}
}
