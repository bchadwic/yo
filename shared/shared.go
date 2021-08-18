package shared

import (
	"bufio"
	"fmt"
	"strings"

	"github.com/bchadwic/yo/yo"
)

type LimitedPrompt interface {
	GetMessage() string
	GetDefault() string
	GetChoices() []string
	GetAttempts() int
	GetValidate() func(s string) bool
	GetCaseSensitive() bool
}

func PromptQuestion(p LimitedPrompt, y *yo.Yo) {
	message := p.GetMessage()
	var choices string
	var def string
	if message == "" {
		message = "Enter a value"
	}
	if d := p.GetDefault(); d != "" {
		def = " [" + d + "]"
	}
	if len(p.GetChoices()) > 0 {
		choices = " (" + strings.Join(p.GetChoices(), ", ") + ")"
	}
	fmt.Fprintf(y.Out, message+choices+def+": ")
}

func ExternalRecieveAnswer(p LimitedPrompt, y *yo.Yo) (string, error) {
	s, err := RecieveAnswer(p, y)
	if y.FailureAttempts >= p.GetAttempts() && p.GetAttempts() != 0 {
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

func RecieveAnswer(p LimitedPrompt, y *yo.Yo) (string, error) {
	s := bufio.NewScanner(y.In)
	s.Scan()
	input := s.Text()
	orgInput := input

	// Write a function to compare regardless of case sensitive
	if !p.GetCaseSensitive() {
		input = strings.ToLower(input)
	}
	for _, e := range p.GetChoices() {
		if input == strings.ToLower(e) {
			return e, nil
		}
	}
	if p.GetDefault() != "" {
		if input == strings.ToLower(p.GetDefault()) || input == "" {
			return p.GetDefault(), nil
		}
	}
	if len(p.GetChoices()) > 0 {
		return "", fmt.Errorf("enter a choice supplied")
	}
	if p.GetValidate() != nil && !p.GetValidate()(orgInput) {
		return "", fmt.Errorf("input is invalid")
	}
	return orgInput, nil
}
