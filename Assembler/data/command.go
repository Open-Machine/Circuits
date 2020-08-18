package data

import (
	"assembler/config"
	"assembler/utils"
)

type Command struct {
	commandCode int
	parameter   int
}

func NewCommand(code int, param int) (*Command, error) {
	if utils.IsOverflow(uint(code), config.AmntBitsCode) {
		return nil, config.MissingBitsCmdCodeError
	}
	if utils.IsOverflow(uint(param), config.AmntBitsParam) {
		return nil, config.CmdCodeOverflowError
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
