package prompt

import (
	"bufio"
	"fmt"
	"strings"

	"github.com/bchadwic/yo/yo"
)

type Prompt struct {
	// prompt properties
	Message string
	Default string
	Choices []string

	// validation
	Attempts      int
	CaseSensitive bool
	Validate      func(s string) bool
}

func (p *Prompt) Prompt(y *yo.Yo) (string, error) {
	y.FailureAttempts = 0
	PromptQuestion(p, y)
	return ExternalRecieveAnswer(p, y)
}

func PromptQuestion(p *Prompt, y *yo.Yo) {
	message := p.Message
	var choices string
	var def string
	if message == "" {
		message = "Enter a value"
	}
	if d := p.Default; d != "" {
		def = " [" + d + "]"
	}
	if len(p.Choices) > 0 {
		choices = " (" + strings.Join(p.Choices, ", ") + ")"
	}
	fmt.Fprintf(y.Out, message+choices+def+": ")
}

func ExternalRecieveAnswer(p *Prompt, y *yo.Yo) (string, error) {
	s, err := RecieveAnswer(p, y)
	if y.FailureAttempts >= p.Attempts && p.Attempts != 0 {
		return "", fmt.Errorf("invalid amount of attempts")
	}
	if err != nil {
		if err.Error() == "enter a valid input" || err.Error() == "enter a choice supplied" {
			y.FailureAttempts++
			fmt.Fprint(y.Err, err.Error()+": ")
			return ExternalRecieveAnswer(p, y)
		}
	}
	return s, err
}

func RecieveAnswer(p *Prompt, y *yo.Yo) (string, error) {
	s := bufio.NewScanner(y.In)
	s.Scan()
	input := s.Text()
	orgInput := input

	// Write a function to compare regardless of case sensitive
	if !p.CaseSensitive {
		input = strings.ToLower(input)
	}

	for _, e := range p.Choices {
		if p.CaseSensitive {
			if input == e {
				return e, nil
			}
		} else {
			if input == strings.ToLower(e) {
				return e, nil
			}
		}
	}
	if p.Default != "" {
		if p.CaseSensitive {
			if input == p.Default {
				return p.Default, nil
			}
		} else {
			if input == strings.ToLower(p.Default) {
				return p.Default, nil
			}
		}
		if input == "" {
			return p.Default, nil
		}
	}
	if len(p.Choices) > 0 {
		return "", fmt.Errorf("enter a choice supplied")
	}
	if p.Validate != nil && !p.Validate(orgInput) {
		return "", fmt.Errorf("input is invalid")
	}
	return orgInput, nil
}