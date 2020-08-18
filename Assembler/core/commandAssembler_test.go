package core

import (
	"assembler/data"
	"reflect"
	"strings"
	"testing"
)

func TestGetNoParam(t *testing.T) {
	got, err := getParamNoParam([]string{"nop"})
	if !(err == nil && got == 0) {
		t.Errorf("Wrong")
	}
}

func TestGetSecondParam(t *testing.T) {
	var tests = []struct {
		line       string
		expected   int
		expectsErr bool
	}{
		{"mov 1", 1, false},
		{"mov 0x1a", 26, false},
		{"mov 0x001", 1, false},
		// Conversion
		{"mov 0x0f", 15, false},
		{"mov 0xff", 255, false},
		{"mov 0x0ff", 255, false},
		{"mov 0xx0ff", 0, true},
		{"mov x1", 0, true},
		{"mov 0x1g", 0, true},
		{"mov 1a", 0, true},
		// Words
		{"mov", 0, true},
		{"mov a b", 0, true},
		{"mov a b c", 0, true},
	}

	for _, test := range tests {
		arrayWords := strings.Split(test.line, " ")
		got, err := getSecondWord(arrayWords)
		gotError := err != nil

		if !(test.expected == got && test.expectsErr == gotError) {
			t.Errorf("Expected int: %d, Got int: %d // Expected error: %t, Got error: %t", test.expected, got, test.expectsErr, gotError)
		}
	}
}

func TestAssembleCommand(t *testing.T) {
	if len(commands) != 12 {
		t.Errorf("Tests were not updated")
	}

	var tests = []struct {
		line       string
		expected   *data.Command
		expectsErr bool
	}{
		// Success
		{"nop", getCommand(0x0, 0), false},
		{"copy 0x10", getCommand(0x1, 16), false},
		{"store 0x10", getCommand(0x2, 16), false},
		{"add 10", getCommand(0x3, 10), false},
		{"sub 10", getCommand(0x4, 10), false},
		{"input 7", getCommand(0x7, 7), false},
		{"output 8", getCommand(0x8, 8), false},
		{"kill", getCommand(0x9, 0), false},
		{"jmp 0x8", getCommand(0xA, 8), false},
		{"jg 0x8", getCommand(0xB, 8), false},
		{"je 0x8", getCommand(0xD, 8), false},
		{"jl 0x8", getCommand(0xF, 8), false},
		// Fail
		{"nope", nil, true},
		{"add x10", nil, true},
		{"kill 0", nil, true},
	}

	for i, test := range tests {
		got, err := AssembleCommand(test.line)
		gotError := err != nil

		if test.expectsErr != gotError {
			t.Errorf("[%d] Expected error: %t, Got error: %t", i, test.expectsErr, gotError)
		} else {
			someErrorIsNil := test.expected == nil || got == nil
			bothErrorsAreNil := test.expected == nil && got == nil
			if someErrorIsNil && !bothErrorsAreNil {
				t.Errorf("Command expected is: %v, Got expected is: %v", test.expected, got)
			}

			if !someErrorIsNil {
				if !reflect.DeepEqual(*test.expected, *got) {
					t.Errorf("Command expected is: %v, Got expected is: %v", *test.expected, *got)
				}
			}
		}
	}
}
func getCommand(code int, param int) *data.Command {
	cmd, _ := data.NewCommand(code, param)
	return cmd
}
