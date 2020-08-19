package data

import (
	"assembler/config"
	"assembler/myerrors"
	"assembler/utils"
	"errors"
)

type Command struct {
	commandCode int
	parameter   CommandParameter
}

func NewCommand(code int, param CommandParameter) (*Command, *myerrors.CustomError) {
	if utils.IsOverflow(uint(code), config.AmntBitsCode) {
		err := myerrors.CommandCodeOverflow(code, config.AmntBitsCode)
		return nil, myerrors.NewAssemblerError(err)
	}
	if !param.IsStr && utils.IsOverflow(uint(param.Num), config.AmntBitsParam) {
		err := myerrors.ParamOverflow(param.Num, config.AmntBitsParam)
		return nil, myerrors.NewCodeError(err)
	}

	return &Command{code, param}, nil
}

func NewCommandTest(code int, param CommandParameter) *Command {
	if !config.Testing {
		return nil
	}
	return &Command{code, param}
}

func (c Command) toExecuter() (string, error) {
	if c.parameter.IsStr {
		// TODO: error or custom error?
		return "", errors.New("Cannot transform command to executer while parameter is still a label")
	}

	str1, err1 := utils.IntToStrHex(c.commandCode, 2)
	if err1 != nil {
		return "", err1
	}

	str2, err2 := utils.IntToStrHex(c.parameter.Num, 2)
	if err2 != nil {
		return "", err2
	}

	return str1 + str2, nil
}
