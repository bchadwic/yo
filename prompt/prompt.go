package prompt

import (
	"github.com/bchadwic/yo/shared"
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
	shared.PromptQuestion(p, y)
	return shared.ExternalRecieveAnswer(p, y)
}

func (p *Prompt) GetMessage() string { return p.Message }

func (p *Prompt) GetDefault() string { return p.Default }

func (p *Prompt) GetChoices() []string { return p.Choices }

func (p *Prompt) GetAttempts() int { return p.Attempts }

func (p *Prompt) GetValidate() func(s string) bool { return p.Validate }

func (p *Prompt) GetCaseSensitive() bool { return p.CaseSensitive }
