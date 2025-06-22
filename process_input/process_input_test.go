package processinput

import (
	"fmt"
	"testing"
)

func TestProcessInput(t *testing.T) {
	tests := []struct {
		name    string
		input   string
		wantCmd string
		wantTxt string
		wantErr bool
	}{
		{"Valid add", "tdsc add", "add", "", false},
		{"Valid add with folder", "tdsc add foldername", "add", "foldername", false},
		{"Valid remove", "tdsc remove foldername", "remove", "foldername", false},
		{"Valid ls", "tdsc ls", "ls", "", false},
		{"Valid help", "tdsc help", "help", "", false},
		{"Invalid empty", "", "", "", true},
		{"Invalid too short", "tdsc", "", "", true},
	}

	for _, tt := range tests {
		t.Run(tt.name, func(t *testing.T) {
			input := tt.input
			got, err := ProcessInput(&input)
			fmt.Printf("%s, %s", got.Command, got.Text)
			if (err != nil) != tt.wantErr {
				t.Errorf("ProcessInput() error = %v, wantErr %v", err, tt.wantErr)
				return
			}
			if !tt.wantErr {
				if got.Command != tt.wantCmd {
					t.Errorf("Command = %v, want %v", got.Command, tt.wantCmd)
				}
				if got.Text != tt.wantTxt {
					t.Errorf("Text = %v, want %v", got.Text, tt.wantTxt)
				}
			}
		})
	}
}
