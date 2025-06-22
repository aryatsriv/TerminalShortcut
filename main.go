package main

import (
	"fmt"
	processinput "github.com/terminal-shortcut/process_input"
)

func main() {
	fmt.Println("Hello, Go")
	test := "lst add run"
	inputData, err := processinput.ProcessInput(&test)
	if err != nil {
		fmt.Printf("Error %s", err)
		return
	}
	// command.CommandRunner(&inputData)
	fmt.Printf("Command: %s", inputData.Command)
}
