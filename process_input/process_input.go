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

	return InputData{
		Command: "parts0",
		Text:    "parts1",
	}, nil
}
