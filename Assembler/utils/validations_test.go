package utils

import "testing"

func TestOverflow(t *testing.T) {
	var tests = []struct {
		num           uint
		availableBits int
		isOverflow    bool
	}{
		{0, 1, false},
		{1, 1, false},
		{2, 1, true},
		{2, 2, false},
		{3, 2, false},
		{4, 2, true},
		{7, 3, false},
		{8, 3, true},
	}

	for _, test := range tests {
		gotIsOverflow := IsOverflow(test.num, test.availableBits)

		if test.isOverflow != gotIsOverflow {
			t.Errorf("Expected overflow: %t, Got overflow: %t // Binary number %d: %b // Available bits: %d", test.isOverflow, gotIsOverflow, test.num, test.num, test.availableBits)
		}
	}
}
