package query

import (
	"bufio"
	"fmt"

	"github.com/bchadwic/yo/internal/msg"
	"github.com/bchadwic/yo/yo"
)

type Query struct {
	Message string
}

func (q *Query) Query(y *yo.Yo) (string, error) {
	outputQuery(q, y)
	return inputQuery(q, y)
}

func outputQuery(q *Query, y *yo.Yo) {
	message := q.Message

	if message == "" {
		message = msg.EnterValue
	}
	fmt.Fprintf(y.Out, message+"\n"+msg.ReturnValue+": \n\n")
}

func inputQuery(q *Query, y *yo.Yo) (string, error) {
	var input string
	s := bufio.NewScanner(y.In)

	emptyLine := false
	for s.Scan() {
		line := s.Text()
		input += line + "\n"
		if line == "" {
			if emptyLine {
				break
			}
			emptyLine = true
		} else {
			emptyLine = false
		}

	}
	return input[:len(input)-2], nil
}
