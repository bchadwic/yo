package write

import (
	"fmt"
	"os"
	"testing"

	"github.com/bchadwic/yo/internal/msg"
	"github.com/bchadwic/yo/pkg/file/read"
	"github.com/bchadwic/yo/yo"
	"github.com/stretchr/testify/assert"
)

const file = "test"

func Test_outputWrite(t *testing.T) {
	tests := []struct {
		name   string
		w      *Write
		output string
	}{
		{
			name:   "empty",
			w:      &Write{},
			output: "",
		},
		{
			name: "preface",
			w: &Write{
				Preface: "Writing just to write",
			},
			output: "Writing just to write\n",
		},
	}
	for _, tt := range tests {
		t.Run(tt.name, func(t *testing.T) {
			testYo, _, out, _ := yo.TestYo()
			outputWrite(tt.w, testYo)
			assert.Equal(t, tt.output, out.String())
		})
	}
}

func Test_inputWrite(t *testing.T) {
	tests := []struct {
		name          string
		w             *Write
		expectedText  string
		expectedError error
		fileWanted    bool
		errWanted     bool
		cleanUp       bool
	}{
		{
			name:          "empty",
			w:             &Write{},
			expectedError: fmt.Errorf(msg.InvalidFile, ""),
			errWanted:     true,
		},
		{
			name: "new file",
			w: &Write{
				File:       file,
				Type:       CREATE,
				Permission: 0777,
			},
			cleanUp: false, // used for next test
		},
		{
			name: "append to file",
			w: &Write{
				File: file,
				Type: APPEND,
				Text: "hello world",
			},
			expectedText: "hello world",
			cleanUp:      false, // used for next test
		},
		{
			name: "replace letters in file",
			w: &Write{
				File: file,
				Type: REPLACE,
				Text: "howdy",
			},
			expectedText: "howdy world",
			cleanUp:      false, // used for next test
		},
		{
			name: "overwrite file",
			w: &Write{
				File:       file,
				Type:       OVERWRITE,
				Text:       "hi",
				Permission: 0777,
			},
			expectedText: "hi",
			cleanUp:      true,
		},
		{
			w: &Write{
				File: file,
				Type: OVERWRITE,
				Text: "hello",
			},
			expectedText:  "",
			expectedError: fmt.Errorf(msg.InvalidType),
			errWanted:     true,
		},
	}
	for _, tt := range tests {
		t.Run(tt.name, func(t *testing.T) {
			testYo, _, _, _ := yo.TestYo()
			err := inputWrite(tt.w, testYo)
			if tt.errWanted {
				assert.EqualError(t, tt.expectedError, err.Error())
			} else {
				assert.NoError(t, err)
			}
			text, _ := (&read.Read{
				File: file,
			}).Read(testYo)
			assert.Equal(t, tt.expectedText, text)
			if tt.w.Type == CREATE {
				assert.True(t, fileFound())
			}
			if tt.cleanUp {
				os.Remove(file)
			}
		})
	}
}

func Test_Write(t *testing.T) {
	yo, _, _, _ := yo.TestYo()
	err := (&Write{}).Write(yo)
	assert.EqualError(t, fmt.Errorf(msg.InvalidFile, ""), err.Error())
	err = (&Write{
		File: file,
		Type: OVERWRITE,
	}).Write(yo)
	assert.EqualError(t, fmt.Errorf(msg.InvalidType), err.Error())
	err = (&Write{
		File: file,
		Type: CREATE,
	}).Write(yo)
	assert.NoError(t, err)
	err = (&Write{
		File: file,
	}).Write(yo)
	assert.EqualError(t, fmt.Errorf(msg.InvalidType), err.Error())
	os.Remove(file)
}

func fileFound() bool {
	_, err := os.Stat(file)
	return err == nil
}
