package main

import (
	"assembler/core"
	"assembler/utils"
	"os"
)

func main() {
	argsWithoutBinaryName := os.Args[1:]
	for i, file := range argsWithoutBinaryName {
		core.AssembleFile(file)
		if i != len(argsWithoutBinaryName)-1 {
			utils.PrintlnOut("")
			utils.PrintlnOut("")
		}
	}
}
