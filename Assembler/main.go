package main

import (
	"assembler/core"
	"assembler/helper"
	"os"
)

func main() {
	argsWithoutBinaryName := os.Args[1:]
	for i, file := range argsWithoutBinaryName {
		core.AssembleFile(file)
		if i != len(argsWithoutBinaryName)-1 {
			helper.PrintlnOut("")
			helper.PrintlnOut("")
		}
	}
}
