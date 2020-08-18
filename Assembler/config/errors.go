package config

import "errors"

var (
	// Commands
	CommandDoesNotExistError = errors.New("Command does not exist")
	MissingBitsCmdCodeError  = errors.New("Invalid command code: not enought bits (overflow)")
	CmdCodeOverflowError     = errors.New("Invalid command param: not enought bits (overflow)")
	// Parameters
	TooManyParamsError = errors.New("Too many parameters (this command only requires one parameter)")
	MissingParamsError = errors.New("Missing parameters")
)
