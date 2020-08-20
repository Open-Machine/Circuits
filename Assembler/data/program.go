package data

import (
	"assembler/myerrors"
)

type Program struct {
	commands       []Command
	gotoLabelsDict map[string]int
}

func NewProgram(lines int) Program {
	return Program{commands: make([]Command, 0, lines), gotoLabelsDict: map[string]int{}}
}

func ProgramFromCommandsAndLabels(commands []Command, gotoLabelsDict map[string]int) Program {
	return Program{commands: commands, gotoLabelsDict: gotoLabelsDict}
}

func (p *Program) AddCommand(command Command) {
	p.commands = append(p.commands, command)
}

func (p *Program) LenCommands() int {
	return len(p.commands)
}

func (p *Program) AddGotoLabel(label string, commandIndex int) error {
	_, exists := p.gotoLabelsDict[label]
	if exists {
		return myerrors.GotoLabelAlreadyExistsError(label)
	}

	p.gotoLabelsDict[label] = commandIndex
	return nil
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
