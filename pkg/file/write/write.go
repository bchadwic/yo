package write

import (
	"fmt"
	"os"

	"github.com/bchadwic/yo/internal/msg"
	"github.com/bchadwic/yo/yo"
)

type Write struct {
	Permission os.FileMode
	Type       int
	File       string
	Text       string
	Preface    string
}

const (
	REPLACE   = 1
	APPEND    = os.O_APPEND
	OVERWRITE = os.O_TRUNC
	CREATE    = os.O_CREATE
)

func (w *Write) Write(y *yo.Yo) error {
	outputWrite(w, y)
	return inputWrite(w, y)
}

func outputWrite(w *Write, y *yo.Yo) {
	if w.Preface != "" {
		fmt.Fprintf(y.Out, "%s\n", w.Preface)
	}
}

func inputWrite(w *Write, y *yo.Yo) error {
	if w.File == "" {
		return fmt.Errorf(msg.InvalidFile, w.File)
	}
	if w.Type == 0 {
		return fmt.Errorf(msg.InvalidType)
	}
	_, err := os.Stat(w.File)
	if w.Type != CREATE && err != nil {
		return fmt.Errorf(msg.InvalidType)
	}
	ty := w.Type
	if w.Type == REPLACE {
		ty = CREATE
	}
	permission := w.Permission
	if permission >= 0777 {
		permission = 0644
	}
	var f *os.File
	f, err = os.OpenFile(w.File, os.O_WRONLY|ty, permission)
	if err != nil {
		return fmt.Errorf(msg.InvalidPermission)
	}
	_, err = f.Write([]byte(w.Text))
	return err
}
