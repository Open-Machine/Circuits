package main

import (
	"fmt"
	"log"
)

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
		break
	case 2:
		break
	}

	fmt.Println()
	fmt.Println("End of execution")
}
