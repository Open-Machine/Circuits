package data

type Program struct {
	commands []Command
}

func NewProgram(lines int) Program {
	return Program{commands: make([]Command, 0, lines)}
}

func ProgramFromCommands(commands []Command) Program {
	return Program{commands: commands}
}

func (p *Program) AddCommand(command Command) {
	p.commands = append(p.commands, command)
}

func (p *Program) ToExecuter() (string, []error) {
	str := ""
	errors := make([]error, 0)

	for _, command := range p.commands {
		executerCode, err := command.toExecuter()
		if err != nil {
			errors = append(errors, err)
		}

		str += executerCode
	}

	return str, errors
}
