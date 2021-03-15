package main

import (
	"io"
	"os/exec"
	"strconv"
)

func runProgram(programPath string) (io.ReadCloser, error) {
	stdout, err := runCommand("java", "-jar", "../logisim-evolution.jar", "../main.circ", "-load", programPath, "-tty", "table")
	return stdout, err
}

func binaryStringToNumber(s string) (uint64, error) {
	ui, err := strconv.ParseUint(s, 2, 64)
	if err != nil {
		return 0, err
	}
	return ui, nil
}

func hexStringToNumber(s string) (uint64, error) {
	ui, err := strconv.ParseUint(s, 16, 64)
	if err != nil {
		return 0, err
	}
	return ui, nil
}

func runCommand(name string, arg ...string) (io.ReadCloser, error) {
	cmd := exec.Command(name, arg...)

	stdout, err := cmd.StdoutPipe()
	if err != nil {
		return nil, err
	}

	err = cmd.Start()
	if err != nil {
		return nil, err
	}

	return stdout, nil
}
