package edit

import (
	"fmt"

	"github.com/bchadwic/yo/yo"
)

type Edit struct {
	Message     string
	Editor      string
	DefaultText string
	FileName    string
}

type File struct {
	date Date
	name string
	path string
}

type Date struct {
	Day   int
	Month int
	Year  int
}

func (e *Edit) Edit(y *yo.Yo) (*File, error) {
	outputPrompt(e, y)
	return inputPrompt(e, y)
}

func outputPrompt(e *Edit, y *yo.Yo) {

	fmt.Fprintf(y.Out, message+choices+def+": ")
}

func inputPrompt(e *Edit, y *yo.Yo) (*File, error) {
	return nil, nil
}
