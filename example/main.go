package main

import (
	"fmt"
	"os"
	"strings"

	"github.com/bchadwic/yo/pkg/file/read"
	"github.com/bchadwic/yo/pkg/file/write"
	"github.com/bchadwic/yo/pkg/input/prompt"
	"github.com/bchadwic/yo/pkg/input/query"
	"github.com/bchadwic/yo/yo"
)

func check(err error) {
	if err != nil {
		panic(err)
	}
}

func main() {
	yo := yo.Yoyo()
	fileName, err := (&prompt.Prompt{
		Message: "Type the name of the file you want to make",
		Default: "temp",
		Validate: func(s string) bool {
			return !strings.ContainsRune(s, ' ')
		},
		Attempts: 0,
	}).Prompt(yo)
	check(err)

	fileContents, err := (&query.Query{
		Message: "Enter the contents of the new file '" + fileName + "'",
	}).Query(yo)
	check(err)

	err = (&write.Write{
		File:       fileName,
		Text:       fileContents,
		Type:       write.CREATE,
		Permission: 0644,
		Preface:    "Writing contents to " + fileName + "\nContents: \n'''\n" + fileContents + "\n'''",
	}).Write(yo)
	check(err)

	readContents, err := (&read.Read{
		File:    fileName,
		Preface: "Reading contents from " + fileName,
	}).Read(yo)
	check(err)
	if fileContents == readContents {
		fmt.Println("file contents matched read contents")
	}
	os.Remove(fileName)
}
