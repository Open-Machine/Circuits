package data

import (
	"assembler/config"
	"testing"
)

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
		_, err := NewCommand(test.code, NewIntParam(test.param))
		gotErr := err != nil

		if test.expectsError != gotErr {
			t.Errorf("[%d] Expected error: %t, Got error: %t", i, test.expectsError, gotErr)
		}
	}
}

func TestNewCommandTest(t *testing.T) {
	oldTesting := config.Testing
	defer func() { config.Testing = oldTesting }()

	config.Testing = false
	ptrCommandNil := NewCommandTest(1, NewIntParam(1))
	if ptrCommandNil != nil {
		t.Errorf("Expected nil command, got not nil command")
	}

	config.Testing = true
	ptrCommandNotNil := NewCommandTest(1, NewIntParam(1))
	if ptrCommandNotNil == nil {
		t.Errorf("Expected nil command, got not nil command")
	}
}
