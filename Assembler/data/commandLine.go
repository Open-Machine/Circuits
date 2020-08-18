package data

type CommandLine struct {
	lineIndex int
	command   Command
}

func newCommandLine(lineIndex int, command Command) CommandLine {
	return CommandLine{lineIndex: lineIndex, command: command}
}

type LineError struct {
	lineIndex int
	err       error
}

func newLineError(lineIndex int, err error) LineError {
	return LineError{lineIndex: lineIndex, err: err}
}
