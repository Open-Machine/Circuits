package myerrors

import (
	"fmt"
)

func CommandCodeOverflow(commandCode int, amntBits int) error {
	return fmt.Errorf("Command code '%b' overflows %d bits", commandCode, amntBits)
}

func ParamOverflow(param int, amntBits int) error {
	return fmt.Errorf("Param '%b' overflow %d bits", param, amntBits)
}

func WrongNumberOfParamsError(command string, amntExpected int, amntReceived int, params []string) error {
	strParameters := ""
	if len(params) == 0 {
		strParameters = "no params"
	} else {
		for i, param := range params {
			strParameters += fmt.Sprintf("'%s'", param)
			if i != len(params)-1 {
				strParameters += ", "
			}
		}
	}

	return fmt.Errorf("The command '%s' requires %d parameters, but received %d parameters (parameters: %s)", command, amntExpected, amntReceived, strParameters)
}

func CommandDoesNotExistError(commandStr string) error {
	return fmt.Errorf("Command '%s' does not exist", commandStr)
}

func InvalidParamError(num int, strLength int, hexStr string) error {
	return fmt.Errorf("Number %d cannot be converted to hexadecimal string of length %d. Got: '%s'", num, strLength, hexStr)
}
