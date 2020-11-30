package main

import (
	"./utils"
	"fmt"
	"os"
)

func main() {
	fmt.Println("Starting editorconfig CLI")
	utils.ProcessCommands(os.Args[1:])
}
