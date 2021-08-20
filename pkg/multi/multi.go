package multi

import (
	"bufio"
	"fmt"

	"github.com/bchadwic/yo/yo"
)

type Multi struct {
	Message string
}

func (m *Multi) Multi(y *yo.Yo) (string, error) {
	multiQuery(m, y)
	return recieveAnswer(m, y)
}

func multiQuery(m *Multi, y *yo.Yo) {
	message := m.Message

	if message == "" {
		message = "Type in a value"
	}
	fmt.Fprintf(y.Out, message+"\nReturn twice to save and quit: \n\n")
}

func recieveAnswer(m *Multi, y *yo.Yo) (string, error) {
	r := bufio.NewReader(y.In)

	var input string
	i := 0
	for i < 2 {
		curr, err := r.ReadString('\n')
		if err != nil {
			return "", err
		}
		if len(curr) <= 1 {
			i++
		}
		input += curr
	}
	return input, nil
}
