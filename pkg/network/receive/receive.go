package receive

import (
	"fmt"

	"github.com/bchadwic/yo/yo"
)

type Recieve struct {
	Endpoint *Endpoint
	Preface  string
	Output   bool
}

func (r *Recieve) Recieve(y *yo.Yo) (string, error) {
	outputRecieve(r, y)
	return "", nil
}

func outputRecieve(r *Recieve, y *yo.Yo) {
	if r.Preface != "" {
		fmt.Fprintf(y.Out, "%s\n", r.Preface)
	}
}

func inputRecieve() {

}
