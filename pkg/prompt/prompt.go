package prompt

import (
	"bufio"
	"fmt"
	"strings"

	"github.com/bchadwic/yo/internal/msg"
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
	outputPrompt(p, y)
	return inputPrompt(p, y)
}

func outputPrompt(p *Prompt, y *yo.Yo) {
	message := p.Message
	var choices string
	var def string
	if message == "" {
		message = msg.EnterValue
	}
	if d := p.Default; d != "" {
		def = " [" + d + "]"
	}
	if len(p.Choices) > 0 {
		choices = " (" + strings.Join(p.Choices, ", ") + ")"
	}
	fmt.Fprintf(y.Out, message+choices+def+": ")
}

func inputPrompt(p *Prompt, y *yo.Yo) (string, error) {
	attempts := 0
	for attempts < p.Attempts || p.Attempts == 0 {
		s, err := internalInputPrompt(p, y)
		if err != nil {
			attempts++
			if attempts >= p.Attempts {
				break
			}
			fmt.Fprint(y.Err, err.Error()+": ")
		} else {
			return s, nil
		}
	}
	return "", fmt.Errorf(msg.InvalidAttempts)
}

func internalInputPrompt(p *Prompt, y *yo.Yo) (string, error) {
	s := bufio.NewScanner(y.In)
	s.Scan()
	input := s.Text()
	orgInput := input

	if !p.CaseSensitive {
		input = strings.ToLower(input)
	}
	if p.Validate != nil && p.Validate(orgInput) {
		return orgInput, nil
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
		return "", fmt.Errorf(msg.InvalidChoice)
	}
	if p.Validate != nil && !p.Validate(orgInput) {
		return "", fmt.Errorf(msg.InvalidValue)
	}
	return orgInput, nil
}
