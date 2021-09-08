package reader

import (
	"fmt"
	"testing"

	"github.com/bchadwic/yo/internal/msg"
	"github.com/bchadwic/yo/yo"
	"github.com/stretchr/testify/assert"
)

func Test_outputRead(t *testing.T) {

	tests := []struct {
		name   string
		r      *Reader
		output string
	}{
		{
			name:   "empty",
			r:      &Reader{},
			output: "",
		},
		{
			name: "preface",
			r: &Reader{
				Preface: "Reader be reading",
			},
			output: "Reader be reading\n",
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
		r             *Reader
		output        string
		errWanted     bool
		expectedError error
	}{
		{
			name: "basic",
			r: &Reader{
				Path: "test.txt",
			},
			output:    "hello\nworld\nthis is a test",
			errWanted: false,
		},
		{
			name: "basic invalid",
			r: &Reader{
				Path: "test.tx",
			},
			output:        "",
			errWanted:     true,
			expectedError: fmt.Errorf(msg.InvalidPath, "test.tx"),
		},
		{
			name: "basic",
			r: &Reader{
				Path:   "test.txt",
				Output: true,
			},
			output:    "hello\nworld\nthis is a test",
			errWanted: false,
		},
		{
			name: "basic invalid",
			r: &Reader{
				Path:   "test.tx",
				Output: true,
			},
			output:        "",
			errWanted:     true,
			expectedError: fmt.Errorf(msg.InvalidPath, "test.tx"),
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
	_, err := (&Reader{
		Preface: "reading test.txt",
		Path:    "test.txt",
		Output:  true,
	}).Read(testYo)
	assert.NoError(t, err)
	assert.Equal(t, "reading test.txt\nhello\nworld\nthis is a test\n", out.String())
}
