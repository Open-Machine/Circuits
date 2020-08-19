package core

import (
	"assembler/data"
	"assembler/myerrors"
	"assembler/utils"
	"bufio"
	"fmt"
	"io"
	"os"
)

func AssembleFile(path string) {
	file, err := os.Open(path)
	if err != nil {
		utils.PrintlnErr(fmt.Sprintf("Cannot open file %s . Error: %s", path, err.Error()))
		return
	}
	defer file.Close()

	utils.PrintlnOut(fmt.Sprintf("========= Starting to assemble %s =========", path))

	ptrProgram := programFromFile(file)
	if ptrProgram == nil {
		printFailedToAssemble(path)
		return
	}

	binaryFileName := utils.FilenameWithoutExtension(file.Name())

	binaryFile, err := os.Create(binaryFileName)
	if err != nil {
		utils.PrintlnErr(fmt.Sprintf("Cannot open file %s . Error: %s", binaryFileName, err.Error()))
		return
	}

	resultCode := writeBinaryProgram(*ptrProgram, binaryFileName, binaryFile)
	if resultCode != 0 {
		printFailedToAssemble(path)
		return
	}
	utils.PrintlnOut(fmt.Sprintf("========= Binary file available in %s =========", binaryFileName))
}

func programFromFile(file io.Reader) *data.Program {
	reader := bufio.NewReader(file)
	lineIndex := 1
	program := data.NewProgram(0)

	successful := true

	for {
		line, errRead := reader.ReadString('\n')

		if errRead != nil && errRead != io.EOF {
			utils.PrintlnErr(fmt.Sprintf("Error while reading file. Error: %s", errRead.Error()))
			return nil
		}

		commandPointer, err := assembleEntireLine(line)

		if err != nil {
			successful = false
			utils.PrintlnErr(fmt.Sprintf("[Error] Error on line %d: %s", lineIndex, line))
			utils.PrintlnErr(fmt.Sprintf("\t\tError: %s", err.Error()))
		} else if commandPointer != nil {
			program.AddCommand(*commandPointer)
		}

		if errRead == io.EOF {
			break
		}
		lineIndex++
	}

	if !successful {
		return nil
	}
	return &program
}

func assembleEntireLine(line string) (*data.Command, *myerrors.CustomError) {
	normalizedStr := utils.LineNormalization(line)

	if normalizedStr == "" {
		return nil, nil
	}

	// TODO: gotoLabel, restOfCommandStr := core.AssembleGotoLabel(normalizedStr)
	restOfCommandStr := normalizedStr

	commandPointer, err := AssembleCommand(restOfCommandStr)
	return commandPointer, err
}

func writeBinaryProgram(program data.Program, binaryFileName string, binaryFile io.Writer) int {
	writer := bufio.NewWriter(binaryFile)
	defer writer.Flush()

	binaryStr, errs := program.ToExecuter()

	if errs != nil && len(errs) > 0 {
		utils.PrintlnErr("[Assembler Errors] It was not possible to assemble your code. Encountered errors:")
		for i, err := range errs {
			utils.PrintlnErr(fmt.Sprintf("\t[%d] %s", i+1, err.Error()))
		}
		return 1
	}

	_, err := writer.WriteString(binaryStr)
	if err != nil {
		utils.PrintlnErr(fmt.Sprintf("Could not write to file %s \n", binaryFileName))
		return 2
	}

	return 0
}

func printFailedToAssemble(path string) {
	utils.PrintlnErr(fmt.Sprintf("========= Failed to assembly %s =========", path))
}
