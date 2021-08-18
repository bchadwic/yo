package yo

/*
select - menu style prompt that allows users to navigate up and down
checklist - like select but multiple selected with space
X prompt - close ended questions (Ex. (Y/N), (Start/Stop))
query - open ended fill in the blank (Ex. Create a commit message) possible character counter?
editor - open editor, return file as variable out of function
slider - move left and right to increase and decrease
*/

import (
	"bytes"
	"io"
	"os"
)

type Yo struct {
	// reader / writer
	In  io.ReadCloser
	Out io.Writer
	Err io.Writer

	// validation
	FailureAttempts int
}

func ProdYo() *Yo {
	return &Yo{
		In:  os.Stdin,
		Out: os.Stdout,
		Err: os.Stderr,
	}
}

func TestYo() (*Yo, *bytes.Buffer, *bytes.Buffer, *bytes.Buffer) {
	in := &bytes.Buffer{}
	out := &bytes.Buffer{}
	err := &bytes.Buffer{}
	return &Yo{
		In:  io.NopCloser(in),
		Out: out,
		Err: err,
	}, in, out, err
}
