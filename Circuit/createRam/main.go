package main

import (
	"fmt"
	"io/ioutil"
	"log"
	"math"
)

const ramFileName = "../RAM.circ"
const orPortFileName = "../OR_PORT.circ"
const fileMode = 777
const newLine = "\n"

const amntBitsComputer = 8 //this will not work "out of the box" if you change this

func getMaxBytesRAM() int {
	return int(math.Pow(2, amntBitsComputer))
}

func main() {
	fmt.Println("MENU: ")
	fmt.Println("1. Generate RAM")
	fmt.Println("2. Create PortOr")

	fmt.Print("Menu item choosen: ")
	var item int
	_, err := fmt.Scanf("%d", &item)
	if err != nil {
		log.Fatal(err)
	}

	switch item {
	case 1:
		generateRAMCircuitWithUser()
		break
	case 2:
		generateOrPortWithUser()
		break
	}

	fmt.Println()
	fmt.Println("End of execution")
}

func generateOrPortWithUser() {
	fmt.Print("Amount of inputs: ")
	var amntInputs int
	_, err := fmt.Scanf("%d", &amntInputs)
	if err != nil {
		log.Fatal(err)
	}

	orPortCircuitStr, errOrPort := generateOrPort(amntInputs)
	if errOrPort != nil {
		log.Fatal(errOrPort)
	}

	errWrite := ioutil.WriteFile(orPortFileName, []byte(orPortCircuitStr), fileMode)
	if errWrite != nil {
		log.Fatal(errWrite)
	}
}

func generateRAMCircuitWithUser() {
	fmt.Print("Amount of bytes in the RAM: ")
	var amntBytes int
	_, err := fmt.Scanf("%d", &amntBytes)
	if err != nil {
		log.Fatal(err)
	}
	if amntBytes <= 0 {
		amntBytes = getMaxBytesRAM() - 1
	}

	ramCircuitStr := generateRAMCircuit(amntBytes)
	errWrite := ioutil.WriteFile(ramFileName, []byte(ramCircuitStr), fileMode)
	if errWrite != nil {
		log.Fatal(err)
	}
}
