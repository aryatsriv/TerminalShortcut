package command

import process "github.com/terminal-shortcut/process_input"

func CommandRunner(data *process.InputData) (res string) {
	commandMapping := map[string]func(data *process.InputData) (res string){
		"add":    addCommand,
		"remove": removeCommand,
		"ls":     listCommand,
		"help":   helpCommand,
		"c":      changeDirectoryCommand,
	}

	functionToCall, exist := commandMapping[data.Command]
	if !exist {
		return "Command not found. See tdsc help"
	}

	return functionToCall(data)

}

func getCurrentDirectoryOfCommand() {

}

func addCommand(data *process.InputData) (res string) {
	return res
}

func removeCommand(data *process.InputData) (res string) {
	return res
}

func listCommand(data *process.InputData) (res string) {
	return res
}

func helpCommand(data *process.InputData) (res string) {
	return res
}

func changeDirectoryCommand(data *process.InputData) (res string) {
	return res
}
