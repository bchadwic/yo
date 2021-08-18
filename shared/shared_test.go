package shared_test

import (
	"fmt"
	"testing"

	"github.com/bchadwic/yo/prompt"
	"github.com/bchadwic/yo/shared"
	"github.com/bchadwic/yo/yo"
	"github.com/stretchr/testify/assert"
)

func Test_PromptQuestion(t *testing.T) {

	tests := []struct {
		name   string
		p      shared.LimitedPrompt
		output string
	}{
		{
			name:   "empty prompt",
			p:      &prompt.Prompt{},
			output: "Enter a value: ",
		},
		{
			name: "messaged prompt",
			p: &prompt.Prompt{
				Message: "What is your favorite type of icecream?",
			},
			output: "What is your favorite type of icecream?: ",
		},
		{
			name: "messaged with default prompt",
			p: &prompt.Prompt{
				Message: "Favorite band?",
				Default: "The Beatles",
			},
			output: "Favorite band? [The Beatles]: ",
		},
		{
			name: "messaged with choices prompt",
			p: &prompt.Prompt{
				Message: "Preferred editor?",
				Choices: []string{"Vim", "Nano", "VS Code"},
			},
			output: "Preferred editor? (Vim, Nano, VS Code): ",
		},
		{
			name: "messaged with choices and default prompt",
			p: &prompt.Prompt{
				Message: "Preferred editor?",
				Choices: []string{"Vim", "Nano", "VS Code"},
				Default: "Vim",
			},
			output: "Preferred editor? (Vim, Nano, VS Code) [Vim]: ",
		},
		{
			name: "empty message with choices prompt",
			p: &prompt.Prompt{
				Choices: []string{"Stop", "Continue"},
			},
			output: "Enter a value (Stop, Continue): ",
		},
		{
			name: "empty message with default prompt",
			p: &prompt.Prompt{
				Default: "Quit",
			},
			output: "Enter a value [Quit]: ",
		},
		{
			name: "empty message with default and choices prompt",
			p: &prompt.Prompt{
				Choices: []string{"Stop", "Continue"},
				Default: "Quit",
			},
			output: "Enter a value (Stop, Continue) [Quit]: ",
		},
	}

	for _, tt := range tests {
		t.Run(tt.name, func(t *testing.T) {
			testYo, _, out, _ := yo.TestYo()
			shared.PromptQuestion(tt.p, testYo)
			assert.Equal(t, tt.output, out.String())
		})
	}
}

func Test_RecieveAnswer(t *testing.T) {
	tests := []struct {
		name          string
		p             shared.LimitedPrompt
		input         string
		output        string
		errWanted     bool
		expectedError error
	}{
		{
			name:      "basic",
			p:         &prompt.Prompt{},
			input:     "yes",
			output:    "yes",
			errWanted: false,
		},
		{
			name: "validation test",
			p: &prompt.Prompt{
				Validate: func(s string) bool {
					return s == "no"
				},
			},
			input:         "yes",
			errWanted:     true,
			expectedError: fmt.Errorf("input is invalid"),
		},
		{
			name: "invalid validation test",
			p: &prompt.Prompt{
				Validate: func(s string) bool {
					return s == "no"
				},
			},
			input:         "yes",
			errWanted:     true,
			expectedError: fmt.Errorf("input is invalid"),
		},
		{
			name: "valid validation test",
			p: &prompt.Prompt{
				Validate: func(s string) bool {
					return s == "no"
				},
			},
			input:  "no",
			output: "no",
		},
		{
			name: "choices test",
			p: &prompt.Prompt{
				Choices: []string{"coca-cola", "pepsi"},
			},
			input:     "coca-cola",
			output:    "coca-cola",
			errWanted: false,
		},
		{
			name: "choices empty test",
			p: &prompt.Prompt{
				Choices: []string{"coca-cola", "pepsi"},
			},
			input:         "",
			errWanted:     true,
			expectedError: fmt.Errorf("enter a choice supplied"),
		},
		{
			name: "invalid choices test",
			p: &prompt.Prompt{
				Choices: []string{"coca-cola", "pepsi"},
			},
			input:         "sprite",
			errWanted:     true,
			expectedError: fmt.Errorf("enter a choice supplied"),
		},
		{
			name: "valid default test",
			p: &prompt.Prompt{
				Default: "quit",
			},
			input:  "",
			output: "quit",
		},
		{
			name: "valid default choices test",
			p: &prompt.Prompt{
				Default: "quit",
				Choices: []string{"continue", "pause"},
			},
			input:  "",
			output: "quit",
		},
		{
			name: "valid default choices user input test",
			p: &prompt.Prompt{
				Default: "quit",
				Choices: []string{"continue", "pause"},
			},
			input:  "continue",
			output: "continue",
		},
	}
	for _, tt := range tests {
		t.Run(tt.name, func(t *testing.T) {
			testYo, in, _, _ := yo.TestYo()
			in.WriteString(tt.input)
			output, err := shared.RecieveAnswer(tt.p, testYo)

			if tt.errWanted {
				assert.EqualError(t, tt.expectedError, err.Error())
			} else {
				assert.NoError(t, err)
			}
			assert.Equal(t, tt.output, output)
		})
	}
}

func Test_ExternalRecieveAnswer(t *testing.T) {
	tests := []struct {
		name            string
		p               shared.LimitedPrompt
		FailureAttempts int
		errWanted       bool
		expectedError   error
	}{
		{
			name: "invalid amount of attempts",
			p: &prompt.Prompt{
				Attempts: 3,
			},
			FailureAttempts: 3,
			errWanted:       true,
			expectedError:   fmt.Errorf("invalid amount of attempts"),
		},
		{
			name: "valid amount of attempts",
			p: &prompt.Prompt{
				Attempts: 3,
			},
			FailureAttempts: 2,
		},
	}
	for _, tt := range tests {
		t.Run(tt.name, func(t *testing.T) {
			testYo, _, _, _ := yo.TestYo()
			testYo.FailureAttempts = tt.FailureAttempts
			_, err := shared.ExternalRecieveAnswer(tt.p, testYo)

			if tt.errWanted {
				assert.EqualError(t, tt.expectedError, err.Error())
			} else {
				assert.NoError(t, err)
			}
		})
	}

}
