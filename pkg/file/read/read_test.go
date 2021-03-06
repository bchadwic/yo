package read

import (
	"fmt"
	"os"
	"testing"

	"github.com/bchadwic/yo/internal/msg"
	"github.com/bchadwic/yo/pkg/file/write"
	"github.com/bchadwic/yo/yo"
	"github.com/stretchr/testify/assert"
)

const file = "test"

func init() {
	err := (&write.Write{
		File:       file,
		Type:       write.CREATE,
		Permission: 0644,
		Text:       "hello\nworld\nthis is a test",
	}).Write(yo.Yoyo())
	if err != nil {
		panic(err)
	}
}

func Test_outputRead(t *testing.T) {
	tests := []struct {
		name   string
		r      *Read
		output string
	}{
		{
			name:   "empty",
			r:      &Read{},
			output: "",
		},
		{
			name: "preface",
			r: &Read{
				Preface: "Read be reading",
			},
			output: "Read be reading\n",
		},
	}
	for _, tt := range tests {
		t.Run(tt.name, func(t *testing.T) {
			testYo, _, out, _ := yo.TestYo()
			outputRead(tt.r, testYo)
			assert.Equal(t, tt.output, out.String())
		})
	}
}

func Test_inputRead(t *testing.T) {
	tests := []struct {
		name          string
		r             *Read
		output        string
		errWanted     bool
		expectedError error
	}{
		{
			name: "basic",
			r: &Read{
				File: file,
			},
			output:    "hello\nworld\nthis is a test",
			errWanted: false,
		},
		{
			name: "basic invalid",
			r: &Read{
				File: "testt",
			},
			output:        "",
			errWanted:     true,
			expectedError: fmt.Errorf(msg.InvalidFile, "testt"),
		},
	}
	for _, tt := range tests {
		t.Run(tt.name, func(t *testing.T) {
			testYo, _, _, _ := yo.TestYo()
			output, err := inputRead(tt.r, testYo)
			if tt.errWanted {
				assert.EqualError(t, tt.expectedError, err.Error())
				assert.Equal(t, tt.output, "")
			} else {
				assert.NoError(t, err)
			}
			assert.Equal(t, tt.output, output)
		})
	}
}

func Test_Read(t *testing.T) {
	testYo, _, out, _ := yo.TestYo()
	_, err := (&Read{
		Preface: "reading test.txt",
		File:    file,
		Output:  true,
	}).Read(testYo)
	assert.NoError(t, err)
	assert.Equal(t, "reading test.txt\nhello\nworld\nthis is a test\n", out.String())
	os.Remove(file)
}
