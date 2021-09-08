package read

import (
	"fmt"
	"os"

	"github.com/bchadwic/yo/internal/msg"
	"github.com/bchadwic/yo/yo"
)

type Read struct {
	Path    string
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
	f, err := os.ReadFile(r.Path)
	if err != nil {
		if r.Output {
			fmt.Fprintf(y.Out, msg.InvalidPath+"\n", r.Path)
		}
		return "", fmt.Errorf(msg.InvalidPath, r.Path)
	}
	sf := string(f)
	if r.Output {
		fmt.Fprintf(y.Out, "%s\n", sf)
	}
	return sf, nil
}
