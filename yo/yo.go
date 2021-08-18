package yo

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
