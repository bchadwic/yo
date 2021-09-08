package reader

import (
	"fmt"
	"os"

	"github.com/bchadwic/yo/internal/msg"
	"github.com/bchadwic/yo/yo"
)

type Reader struct {
	Path    string
	Preface string
	Output  bool
}

func (r *Reader) Read(y *yo.Yo) (string, error) {
	outputRead(r, y)
	return inputRead(r, y)
}

func outputRead(r *Reader, y *yo.Yo) {
	if r.Preface != "" {
		fmt.Fprintf(y.Out, "%s\n", r.Preface)
	}
}

func inputRead(r *Reader, y *yo.Yo) (string, error) {
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
