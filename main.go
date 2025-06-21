package main

import (
	"fmt"

	"github.com/terminal-shortcut/command"
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
	command.CommandRunner(&inputData)
	fmt.Printf("Command: %s", inputData.Command)

}

//There are 3 commands to support
// 1 rs add
// 2 rs add foldernamd
// 3 rs remove foldername
// 4 rs ls
// 5 rs help
// 6 rs --help
