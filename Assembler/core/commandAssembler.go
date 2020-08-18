package core

import (
	"assembler/config"
	"assembler/data"
	"assembler/utils"
	"strings"
)

type commandConfig struct {
	code     int
	getParam func(words []string) (int, error)
}

func AssembleCommand(line string) (*data.Command, error) {
	arrayWords := strings.Split(line, " ")
	commandName := arrayWords[0]

	commandConfig, exists := commands[commandName]
	if !exists {
		return nil, config.CommandDoesNotExistError
	}

	param, paramErr := commandConfig.getParam(arrayWords)
	if paramErr != nil {
		return nil, paramErr
	}

	commandPointer, err := data.NewCommand(commandConfig.code, param)
	return commandPointer, err
}

func getParamNoParam(words []string) (int, error) {
	if len(words) != 1 {
		return 0, config.TooManyParamsError
	}

	return 0, nil
}

func getSecondWord(words []string) (int, error) {
	if len(words) != 2 {
		if len(words) < 2 {
			return 0, config.MissingParamsError
		}

		return 0, config.TooManyParamsError
	}

	num, err := utils.StrToPositiveInt(words[1])
	if err != nil {
		return 0, err
	}
	return num, err
}

var commands = map[string]commandConfig{
	"nop": commandConfig{
		getParam: getParamNoParam,
		code:     0x0,
	},
	"copy": commandConfig{
		getParam: getSecondWord,
		code:     0x1,
	},
	"store": commandConfig{
		getParam: getSecondWord,
		code:     0x2,
	},
	"add": commandConfig{
		getParam: getSecondWord,
		code:     0x3,
	},
	"sub": commandConfig{
		getParam: getSecondWord,
		code:     0x4,
	},
	"input": commandConfig{
		getParam: getSecondWord,
		code:     0x7,
	},
	"output": commandConfig{
		getParam: getSecondWord,
		code:     0x8,
	},
	"kill": commandConfig{
		getParam: getParamNoParam,
		code:     0x9,
	},
	"jmp": commandConfig{
		getParam: getSecondWord,
		code:     0xA,
	},
	"jg": commandConfig{
		getParam: getSecondWord,
		code:     0xB,
	},
	"je": commandConfig{
		getParam: getSecondWord,
		code:     0xD,
	},
	"jl": commandConfig{
		getParam: getSecondWord,
		code:     0xF,
	},
}
