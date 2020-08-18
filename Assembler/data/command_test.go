package data

import "testing"

func TestNewCommandOverflowValidation(t *testing.T) {
	var tests = []struct {
		code         int
		param        int
		expectsError bool
	}{
		{0, 0, false},
		{-1, 0, true},
		{0, -1, true},
		{1000, 1000, true},
		{255, 0, false},
		{256, 0, true},
		{0, 255, false},
		{0, 256, true},
	}

	for i, test := range tests {
		_, err := NewCommand(test.code, test.param)
		gotErr := err != nil

		if test.expectsError != gotErr {
			t.Errorf("[%d] Expected error: %t, Got error: %t", i, test.expectsError, gotErr)
		}
	}
}
