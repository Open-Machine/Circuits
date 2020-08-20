package core

import (
	"assembler/config"
	"assembler/data"
	"assembler/utils"
	"bytes"
	"reflect"
	"strings"
	"testing"
)

func TestAssembleFile(t *testing.T) {
	// TODO
}

func TestProgramFromFile(t *testing.T) {
	stdout := utils.Out
	utils.Out = new(bytes.Buffer)
	defer func() { utils.Out = stdout }()

	stderr := utils.Err
	utils.Err = new(bytes.Buffer)
	defer func() { utils.Err = stderr }()

	var tests = []struct {
		lines      []string
		expected   *data.Program
		expectsErr bool
	}{
		// blank lines
		{
			[]string{
				"       ",
				" ",
				"    ",
			},
			newProgramPointer([]data.Command{}),
			false,
		},
		// blank lines and right methods
		{
			[]string{
				"",
				"  ",
				"store 1",
				"store 0xC",
			},
			newProgramPointer([]data.Command{
				*newCommand(0x2, 1),
				*newCommand(0x2, 12),
			}),
			false,
		},
		// wrong method
		{
			[]string{
				"mov 1",
			},
			nil,
			true,
		},
		// wrong params
		{
			[]string{
				"store",
			},
			nil,
			true,
		},
	}

	for i, test := range tests {
		str := ""
		for i, line := range test.lines {
			str += line
			if i != len(test.lines)-1 {
				str += "\n"
			}
		}

		got := programFromFile(strings.NewReader(str))

		bothNil := test.expected == nil && got == nil
		someNil := test.expected == nil || got == nil

		if !bothNil {
			if someNil || !reflect.DeepEqual(*test.expected, *got) {
				t.Errorf("[%d] Expected: %v, Got: %v", i, test.expected, got)
			}
		}

		stderrStr := utils.Err.(*bytes.Buffer).String()
		gotErr := stderrStr != ""
		if test.expectsErr != gotErr {
			t.Errorf("[%d] Expected error: %t, Got error: %t // ", i, test.expectsErr, gotErr)
			t.Errorf("\t\t StdErr: %s", stderrStr)
		}
	}
}
func newProgramPointer(commands []data.Command) *data.Program {
	prog := data.ProgramFromCommands(commands)
	return &prog
}

func TestAssembleEntireLine(t *testing.T) {
	var tests = []struct {
		param            string
		expectedLabel    *string
		expectedCmd      *data.Command
		amntErrsExpected int
	}{
		// Success without label
		{"", nil, nil, 0},
		{" 	 ", nil, nil, 0},
		{"	 	 	", nil, nil, 0},
		{"	input 	1 ", nil, newCommand(7, 1), 0},
		{"	input 	0x1 ", nil, newCommand(7, 1), 0},
		// Fail without label
		{"	inputa 	0x1 ", nil, nil, 1},
		{"	inputa 	a0x1 ", nil, nil, 1},
		// Success with label
		{"	label: input 	0x1 ", utils.NewString("label"), newCommand(7, 1), 0},
		// Fail with label
		{"	label: inputa 	0x1 ", utils.NewString("label"), nil, 1},
		{"	1label: inputa 	0x1 ", nil, nil, 2},
		{"	1label: input 	0x1 ", nil, newCommand(7, 1), 1},
	}

	for i, test := range tests {
		gotLabel, gotCmd, errs := assembleEntireLine(test.param)

		bothNil := test.expectedCmd == nil && gotCmd == nil
		someNil := test.expectedCmd == nil || gotCmd == nil

		if !bothNil {
			if someNil || *test.expectedCmd != *gotCmd {
				t.Errorf("[%d] Expected: %v, Got: %v", i, test.expectedCmd, gotCmd)
			}
		}

		if !utils.SafeIsEqualStrPointer(gotLabel, test.expectedLabel) {
			t.Errorf("[%d] Expected: %v, Got: %v", i, test.expectedLabel, gotLabel)
		}

		if len(errs) != test.amntErrsExpected {
			t.Errorf("[%d] Amnt errors expected: %d, Errors: %v", i, test.amntErrsExpected, errs)
		}
	}
}

func TestWriteAssembledFile(t *testing.T) {
	stdout := utils.Out
	utils.Out = new(bytes.Buffer)
	defer func() { utils.Out = stdout }()

	stderr := utils.Err
	utils.Err = new(bytes.Buffer)
	defer func() { utils.Err = stderr }()

	oldTesting := config.Testing
	config.Testing = true
	defer func() { config.Testing = oldTesting }()

	var tests = []struct {
		param           data.Program
		expectedFileStr string
		expectedCode    int
		expectsErr      bool
	}{
		{
			data.ProgramFromCommands([]data.Command{
				*newCommand(0x2, 1),
				*newCommand(0x2, 12),
			}),
			"0201020c",
			0,
			false,
		},
		{
			data.ProgramFromCommands([]data.Command{
				*newCommand(0x2, 1),
				*newCommand(0x2, 12),
				*newCommand(0xD, 15),
			}),
			"0201020c0d0f",
			0,
			false,
		},
		{
			data.ProgramFromCommands([]data.Command{
				*newCommand(0x2, 1),
				*data.NewCommandTest(0x222, data.NewIntParam(12)),
				*newCommand(0xD, 115),
			}),
			"",
			1,
			true,
		},
		{
			data.ProgramFromCommands([]data.Command{
				*newCommand(0x2, 1),
				*data.NewCommandTest(0x2, data.NewStringParam("a")),
			}),
			"",
			1,
			true,
		},
	}

	for i, test := range tests {
		fileWriter := new(bytes.Buffer)
		got := writeBinaryProgram(test.param, "File", fileWriter)

		if !reflect.DeepEqual(test.expectedCode, got) {
			t.Errorf("[%d] Expected: %v, Got: %v", i, test.expectedCode, got)
		}

		gotFileStr := fileWriter.String()
		if gotFileStr != test.expectedFileStr {
			t.Errorf("[%d] Expected file str: %v, Got file str: %v", i, test.expectedFileStr, gotFileStr)
		}

		stderrStr := utils.Err.(*bytes.Buffer).String()
		gotErr := stderrStr != ""
		if test.expectsErr != gotErr {
			t.Errorf("[%d] Expected error: %t, Got error: %t // ", i, test.expectsErr, gotErr)
			t.Errorf("\t\t StdErr: %s", stderrStr)
		}
	}
}

func newCommand(cmdCode int, param int) *data.Command {
	got, _ := data.NewCommand(cmdCode, data.NewIntParam(param))
	return got
}
