package core

import (
	"assembler/data"
	"assembler/utils"
	"fmt"
)

type Assembler struct {
	path    *string
	program *data.Program
}

func NewAssembler() Assembler {
	return Assembler{path: nil, program: nil}
}

func AssembleFile(path string) string {
	// TODO
	panic("not done")
}

func (a *Assembler) assembleEntireLine(line string, lineIndex int) {
	normalizedStr := utils.LineNormalization(line)

	// gotoLabel, restOfCommandStr := core.AssembleGotoLabel(normalizedStr)
	restOfCommandStr := normalizedStr

	commandPointer, err := AssembleCommand(restOfCommandStr)

	if err != nil {
		fmt.Printf("[Assembler Error] Error on line %d: %s\n", lineIndex, line)
		fmt.Printf("\t\tError: %s\n", err.Error())
	} else {
		a.program.AddCommand(*commandPointer)
	}
}
