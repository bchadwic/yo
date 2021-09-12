package read

import (
	"fmt"
	"os"

	"github.com/bchadwic/yo/internal/msg"
	"github.com/bchadwic/yo/yo"
)

type Read struct {
	File    string
	Preface string
	Output  bool
}

func (r *Read) Read(y *yo.Yo) (string, error) {
	outputRead(r, y)
	return inputRead(r, y)
}

func outputRead(r *Read, y *yo.Yo) {
	if r.Preface != "" {
		fmt.Fprintf(y.Out, "%s\n", r.Preface)
	}
}

func inputRead(r *Read, y *yo.Yo) (string, error) {
	f, err := os.ReadFile(r.File)
	if err != nil {
		if r.Output {
			fmt.Fprintf(y.Out, msg.InvalidFile+"\n", r.File)
		}
		return "", fmt.Errorf(msg.InvalidFile, r.File)
	}
	sf := string(f)
	if r.Output {
		fmt.Fprintf(y.Out, "%s\n", sf)
	}
	return sf, nil
}
