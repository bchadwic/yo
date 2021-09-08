package query

import (
	"testing"

	"github.com/bchadwic/yo/internal/msg"
	"github.com/bchadwic/yo/yo"
	"github.com/stretchr/testify/assert"
)

func Test_outputQuery(t *testing.T) {
	tests := []struct {
		name   string
		q      *Query
		output string
	}{
		{
			name:   "empty multi",
			q:      &Query{},
			output: msg.EnterValue + "\n" + msg.ReturnValue + ": \n\n",
		},
		{
			name: "empty multi",
			q: &Query{
				Message: "Type out the message you would like to send",
			},
			output: "Type out the message you would like to send\n" + msg.ReturnValue + ": \n\n",
		},
	}

	for _, tt := range tests {
		t.Run(tt.name, func(t *testing.T) {
			testYo, _, out, _ := yo.TestYo()
			outputQuery(tt.q, testYo)
			assert.Equal(t, tt.output, out.String())
		})
	}
}

func Test_inputQuery(t *testing.T) {
	tests := []struct {
		name        string
		q           *Query
		input       string
		output      string
		errWanted   bool
		errExpected error
	}{
		{
			name:   "empty input",
			input:  "\n\n",
			output: "",
		},
		{
			name:   "single line input",
			input:  "hi\n\n",
			output: "hi",
		},
		{
			name:   "double line input",
			input:  "hi\nhello\n\n",
			output: "hi\nhello",
		},
		{
			name:   "tab input",
			input:  "hello \tworld\n\n",
			output: "hello \tworld",
		},
	}
	for _, tt := range tests {
		t.Run(tt.name, func(t *testing.T) {
			testYo, in, _, _ := yo.TestYo()
			in.Write([]byte(tt.input))
			output, err := inputQuery(tt.q, testYo)

			if tt.errWanted {
				assert.Error(t, tt.errExpected, err)
			} else {
				assert.NoError(t, err)
			}
			assert.Equal(t, tt.output, output)
		})
	}
}
