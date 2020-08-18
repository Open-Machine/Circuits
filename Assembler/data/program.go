package data

type Program struct {
	commandLines []CommandLine
}

func NewProgram(lines int) Program {
	return Program{commandLines: make([]CommandLine, 0, lines)}
}

func (p *Program) AddCommand(line int, command Command) {
	p.commandLines = append(p.commandLines, newCommandLine(line, command))
}

func (p *Program) ToExecuter() (string, []LineError) {
	str := ""
	errors := make([]LineError, 0)

	for _, commandLine := range p.commandLines {
		executerCode, err := commandLine.command.toExecuter()
		if err != nil {
			lineErr := newLineError(commandLine.lineIndex, err)
			errors = append(errors, lineErr)
		}

		str += executerCode
	}

	return str, errors
}
