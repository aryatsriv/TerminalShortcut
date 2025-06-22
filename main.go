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

//commands to support
// 1 tdsc add
// 2 tdsc add foldername
// 3 tdsc remove foldername
// 4 tdsc ls
// 5 tdsc help
// 6 tdsc --help
// 7 tdsc c foldername
