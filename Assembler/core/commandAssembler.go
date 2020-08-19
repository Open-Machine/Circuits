package core

import (
	"assembler/data"
	"assembler/errors"
	"assembler/utils"
	"strings"
)

type commandConfig struct {
	code     int
	getParam func(commandName string, words []string) (int, *errors.CustomError)
}

func AssembleCommand(line string) (*data.Command, *errors.CustomError) {
	arrayWords := strings.Split(line, " ")
	commandName := arrayWords[0]

	commandConfig, exists := commands[commandName]
	if !exists {
		err := errors.CommandDoesNotExistError(commandName)
		return nil, errors.NewCodeError(err)
	}

	param, paramErr := commandConfig.getParam(commandName, arrayWords)
	if paramErr != nil {
		return nil, paramErr
	}

	commandPointer, customErr := data.NewCommand(commandConfig.code, param)
	return commandPointer, customErr
}

func getParamNoParam(commandName string, words []string) (int, *errors.CustomError) {
	if len(words) != 1 {
		remainingParams := getCommandParams(words)
		err := errors.WrongNumberOfParamsError(commandName, 0, len(remainingParams), remainingParams)
		return 0, errors.NewCodeError(err)
	}

	return 0, nil
}

func getSecondWord(commandName string, words []string) (int, *errors.CustomError) {
	if len(words) != 2 {
		if len(words) < 2 {
			err := errors.WrongNumberOfParamsError(commandName, 1, 0, []string{})
			return 0, errors.NewCodeError(err)
		}

		remainingParams := getCommandParams(words)
		err := errors.WrongNumberOfParamsError(commandName, 1, len(remainingParams), remainingParams)
		return 0, errors.NewCodeError(err)
	}

	num, err := utils.StrToPositiveInt(words[1])
	if err != nil {
		return 0, errors.NewCodeError(err)
	}
	return num, nil
}

func getCommandParams(words []string) []string {
	return words[1:]
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
