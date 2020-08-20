package core

import (
	"assembler/myerrors"
	"assembler/utils"
	"strings"
)

func AssembleGotoLabel(line string) (*string, string, error) {
	indexOfColon := strings.Index(line, ":")

	if indexOfColon < 0 {
		return nil, line, nil
	}

	label := strings.TrimSpace(line[:indexOfColon])
	restOfLine := strings.TrimSpace(line[indexOfColon+1:])

	if !utils.IsValidVarName(label) {
		return nil, restOfLine, myerrors.InvalidLabelParam(label)
	}

	return &label, restOfLine, nil
}
