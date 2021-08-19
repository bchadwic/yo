package multi

import (
	"bufio"
	"fmt"

	"github.com/bchadwic/yo/yo"
)

type Multi struct {
	Message string
	Save    *escapeKeyValue
}

type escapeKeyValue struct {
	escapeSequence byte
	escapeMessage  string
}

func EscKeyValue(escapeSequence byte, escapeMessage string) *escapeKeyValue {
	return &escapeKeyValue{
		escapeSequence: escapeSequence,
		escapeMessage:  escapeMessage,
	}
}

func (m *Multi) Multi(y *yo.Yo) (string, error) {
	multiQuery(m, y)
	return recieveAnswer(m, y)
}

func multiQuery(m *Multi, y *yo.Yo) {
	message := m.Message
	saveKV := m.Save

	if message == "" {
		message = "Type in a value"
	}
	if saveKV == nil {
		saveKV = EscKeyValue('*', "return (*) to save and quit")
	}
	fmt.Fprintf(y.Out, message+"\n"+saveKV.escapeMessage+": \n")
}

func recieveAnswer(m *Multi, y *yo.Yo) (string, error) {
	r := bufio.NewReader(y.In)

	var input string
	i := 0
	for i < 2 {
		curr, err := r.ReadString('\x1b')
		if err != nil {
			return "", err
		}
		if len(curr) <= 1 {
			i++
		}
		input += curr
	}
	fmt.Println(input)

	return input, nil
}
