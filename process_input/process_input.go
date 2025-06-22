package processinput

import (
	"fmt"
	"strings"

	"github.com/terminal-shortcut/constants"
)

func ProcessInput(data *string) (res InputData, err error) {
	parts := strings.Split(*data, " ")

	if len(parts) == 0 || len(parts) < constants.MaxCountOfSplitsInCommandInput {
		return InputData{}, fmt.Errorf("Invalid Command")
	}

	for i, part := range parts {
		fmt.Printf("%d, %s\n", i, part)
	}
	inputData := InputData{
		Command: parts[1],
	}
	if (len(parts)) > constants.MaxCountOfSplitsInCommandInput {
		inputData.Text = parts[2]
	}

	return inputData, nil
}

//There are 3 commands to support
// 1 rs add
// 2 rs add foldernamd
// 3 rs remove foldername
// 4 rs ls
// 5 rs help
// 6 rs --help
