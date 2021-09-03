package edit

import (
	"runtime"
	"strings"

	"github.com/bchadwic/yo/pkg/prompt"
	"github.com/bchadwic/yo/yo"
)

type Edit struct {
	Message       string
	DefaultText   string
	DefaultEditor string
	Choices       []string
}

type File struct {
	Path string
	Name string
}

func (e *Edit) Edit(y *yo.Yo) (*File, error) {
	outputEdit(e, y)
	return inputEdit()
}

func outputEdit(e *Edit, y *yo.Yo) (string, error) {
	message := e.Message
	if message == "" {
		message = "Type in an editor to use"
	}
	editor := e.DefaultEditor
	if editor == "" {
		os := runtime.GOOS
		switch os {
		case "windows":
			editor = "notepad.exe"
		default:
			editor = "vi"
		}
	}
	editor, err := (&prompt.Prompt{
		Message: message,
		Default: e.DefaultEditor,
		Choices: e.Choices,
		Validate: func(s string) bool {
			return len(strings.Fields(s)) == 1
		},
	}).Prompt(y)
	return editor, err
}

func inputEdit() (*File, error) {
	return nil, nil
}

// os := runtime.GOOS
// switch os {
// case "windows":
// 	e.Editor = "notepad.exe"
// case "darwin":
// 	e.Editor = "notepad.exe"
// case "linux":
// 	fmt.Println("Linux")
// default:
// 	fmt.Printf("%s.\n", os)
// }
