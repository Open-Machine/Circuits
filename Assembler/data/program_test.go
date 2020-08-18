package data

import "testing"

func TestAddCommand(t *testing.T) {
	program := NewProgram(5)

	if len(program.commandLines) != 0 {
		t.Errorf("Expected length 0, got: %d", len(program.commandLines))
	}

	program.AddCommand(10, Command{0, 0})

	if len(program.commandLines) != 1 {
		t.Errorf("Expected length 1, got: %d", len(program.commandLines))
	}
}

func TestToExecuterSuccess(t *testing.T) {
	program := NewProgram(3)
	program.AddCommand(0, Command{1, 2})
	program.AddCommand(1, Command{15, 7})
	program.AddCommand(2, Command{0, 0})

	got, errors := program.ToExecuter()
	expected := "01020f070000"

	if !(len(errors) == 0 && got == expected) {
		t.Errorf("Expected: '%s', got: '%s'", expected, got)
	}
}

func TestToExecuterFail(t *testing.T) {
	program := NewProgram(3)
	program.AddCommand(0, Command{1, 2})
	program.AddCommand(3, Command{1200, 7})
	program.AddCommand(7, Command{0, 0})

	execCode, errors := program.ToExecuter()

	if !(len(errors) == 1 && errors[0].lineIndex == 3) {
		t.Errorf("Should result in error because of overflow. Executer code: %s // Errors: %v", execCode, errors)
	}
}
