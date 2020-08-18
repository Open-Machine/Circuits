package data

import (
	"assembler/config"
	"assembler/errors"
	"assembler/utils"
)

type Command struct {
	commandCode int
	parameter   int
}

func NewCommand(code int, param int) (*Command, *errors.CustomError) {
	if utils.IsOverflow(uint(code), config.AmntBitsCode) {
		err := errors.CommandCodeOverflow(code, config.AmntBitsCode)
		return nil, errors.NewCustomError(err, errors.AssemblerError)
	}
	if utils.IsOverflow(uint(param), config.AmntBitsParam) {
		err := errors.ParamOverflow(param, config.AmntBitsParam)
		return nil, errors.NewCustomError(err, errors.CodeError)
	}

	return &Command{code, param}, nil
}

func (c Command) toExecuter() (string, error) {
	str1, err1 := utils.IntToStrHex(c.commandCode, 2)
	if err1 != nil {
		return "", err1
	}

	str2, err2 := utils.IntToStrHex(c.parameter, 2)
	if err2 != nil {
		return "", err2
	}

	return str1 + str2, nil
}
