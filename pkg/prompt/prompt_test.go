package prompt

import (
	"fmt"
	"testing"

	"github.com/bchadwic/yo/internal/msg"
	"github.com/bchadwic/yo/yo"
	"github.com/stretchr/testify/assert"
)

func Test_PromptQuestion(t *testing.T) {

	tests := []struct {
		name   string
		p      *Prompt
		output string
	}{
		{
			name:   "empty prompt",
			p:      &Prompt{},
			output: msg.EnterValue + ": ",
		},
		{
			name: "messaged prompt",
			p: &Prompt{
				Message: "What is your favorite type of icecream?",
			},
			output: "What is your favorite type of icecream?: ",
		},
		{
			name: "messaged with default prompt",
			p: &Prompt{
				Message: "Favorite band?",
				Default: "The Beatles",
			},
			output: "Favorite band? [The Beatles]: ",
		},
		{
			name: "messaged with choices prompt",
			p: &Prompt{
				Message: "Preferred editor?",
				Choices: []string{"Vim", "Nano", "VS Code"},
			},
			output: "Preferred editor? (Vim, Nano, VS Code): ",
		},
		{
			name: "messaged with choices and default prompt",
			p: &Prompt{
				Message: "Preferred editor?",
				Choices: []string{"Vim", "Nano", "VS Code"},
				Default: "Vim",
			},
			output: "Preferred editor? (Vim, Nano, VS Code) [Vim]: ",
		},
		{
			name: "empty message with choices prompt",
			p: &Prompt{
				Choices: []string{"Stop", "Continue"},
			},
			output: msg.EnterValue + " (Stop, Continue): ",
		},
		{
			name: "empty message with default prompt",
			p: &Prompt{
				Default: "Quit",
			},
			output: msg.EnterValue + " [Quit]: ",
		},
		{
			name: "empty message with default and choices prompt",
			p: &Prompt{
				Choices: []string{"Stop", "Continue"},
				Default: "Quit",
			},
			output: msg.EnterValue + " (Stop, Continue) [Quit]: ",
		},
	}

	for _, tt := range tests {
		t.Run(tt.name, func(t *testing.T) {
			testYo, _, out, _ := yo.TestYo()
			outputPrompt(tt.p, testYo)
			assert.Equal(t, tt.output, out.String())
		})
	}
}

func Test_RecieveAnswer(t *testing.T) {
	tests := []struct {
		name          string
		p             *Prompt
		input         string
		output        string
		errWanted     bool
		expectedError error
	}{
		{
			name:      "basic",
			p:         &Prompt{},
			input:     "yes",
			output:    "yes",
			errWanted: false,
		},
		{
			name: "validation test",
			p: &Prompt{
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
			p: &Prompt{
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
			p: &Prompt{
				Validate: func(s string) bool {
					return s == "no"
				},
			},
			input:  "no",
			output: "no",
		},
		{
			name: "valid validation test with choices",
			p: &Prompt{
				Validate: func(s string) bool {
					return s == "no"
				},
				Choices: []string{"not no", "anything other than no"},
			},
			input:  "no",
			output: "no",
		},
		{
			name: "choices test",
			p: &Prompt{
				Choices: []string{"coca-cola", "pepsi"},
			},
			input:     "coca-cola",
			output:    "coca-cola",
			errWanted: false,
		},
		{
			name: "choices empty test",
			p: &Prompt{
				Choices: []string{"coca-cola", "pepsi"},
			},
			input:         "",
			errWanted:     true,
			expectedError: fmt.Errorf(msg.InvalidChoice),
		},
		{
			name: "invalid choices test",
			p: &Prompt{
				Choices: []string{"coca-cola", "pepsi"},
			},
			input:         "sprite",
			errWanted:     true,
			expectedError: fmt.Errorf(msg.InvalidChoice),
		},
		{
			name: "valid default test",
			p: &Prompt{
				Default: "quit",
			},
			input:  "",
			output: "quit",
		},
		{
			name: "valid default choices test",
			p: &Prompt{
				Default: "quit",
				Choices: []string{"continue", "pause"},
			},
			input:  "",
			output: "quit",
		},
		{
			name: "valid default choices user input test",
			p: &Prompt{
				Default: "quit",
				Choices: []string{"continue", "pause"},
			},
			input:  "continue",
			output: "continue",
		},
		{
			name: "valid case sensitive test",
			p: &Prompt{
				Choices: []string{"continue", "pause"},
			},
			input:  "CoNtiNuE",
			output: "continue",
		},
		{
			name: "invalid case sensitive test",
			p: &Prompt{
				Choices:       []string{"continue", "pause"},
				CaseSensitive: true,
			},
			input:         "ConTiNUe",
			output:        "",
			errWanted:     true,
			expectedError: fmt.Errorf(msg.InvalidChoice),
		},
		{
			name: "valid default choices case sensitive test 1",
			p: &Prompt{
				Default: "quit",
				Choices: []string{"continue", "pause"},
			},
			input:  "quIt",
			output: "quit",
		},
		{
			name: "invalid default choices case sensitive test 2",
			p: &Prompt{
				Default:       "quit",
				Choices:       []string{"continue", "pause"},
				CaseSensitive: true,
			},
			input:         "qUiT",
			output:        "",
			errWanted:     true,
			expectedError: fmt.Errorf(msg.InvalidChoice),
		},
		{
			name: "valid default choices case sensitive test 3",
			p: &Prompt{
				Default:       "quit",
				Choices:       []string{"continue", "pause"},
				CaseSensitive: true,
			},
			input:  "quit",
			output: "quit",
		},
		{
			name: "valid default choices case sensitive test 4",
			p: &Prompt{
				Default:       "quit",
				Choices:       []string{"continue", "pause"},
				CaseSensitive: true,
			},
			input:  "pause",
			output: "pause",
		},
	}
	for _, tt := range tests {
		t.Run(tt.name, func(t *testing.T) {
			testYo, in, _, _ := yo.TestYo()
			in.WriteString(tt.input)
			output, err := internalInputPrompt(tt.p, testYo)

			if tt.errWanted {
				assert.EqualError(t, tt.expectedError, err.Error())
				assert.Equal(t, tt.output, "")
			} else {
				assert.NoError(t, err)
			}
			assert.Equal(t, tt.output, output)
		})
	}
}

func Test_externalRecieveAnswer(t *testing.T) {
	tests := []struct {
		name            string
		p               *Prompt
		FailureAttempts int
		errWanted       bool
		expectedError   error
	}{
		{
			name: "invalid amount of attempts",
			p: &Prompt{
				Attempts: 3,
			},
			FailureAttempts: 3,
			errWanted:       true,
			expectedError:   fmt.Errorf("invalid amount of attempts"),
		},
		{
			name: "valid amount of attempts",
			p: &Prompt{
				Attempts: 3,
			},
			FailureAttempts: 2,
		},
	}
	for _, tt := range tests {
		t.Run(tt.name, func(t *testing.T) {
			testYo, _, _, _ := yo.TestYo()
			testYo.FailureAttempts = tt.FailureAttempts
			_, err := inputPrompt(tt.p, testYo)

			if tt.errWanted {
				assert.EqualError(t, tt.expectedError, err.Error())
			} else {
				assert.NoError(t, err)
			}
		})
	}

}
