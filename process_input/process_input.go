package processinput

import (
	"errors"
	"fmt"
	"github.com/terminal-shortcut/constants"
	"strings"
)

func ProcessInput(data *string) (res InputData, err error) {
	parts := strings.Split(*data, " ")

	if len(parts) == 0 || len(parts) < constants.MaxCountOfSplitsInCommandInput {
		return InputData{}, errors.New("invalid command")
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
