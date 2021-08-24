package main

import (
	"fmt"

	"github.com/bchadwic/yo/pkg/prompt"
	"github.com/bchadwic/yo/pkg/query"
	"github.com/bchadwic/yo/yo"
)

func main() {
	yo := yo.ProdYo()
	// fmt.Println()
	// (&prompt.Prompt{
	// 	Message: "Do you like icecream?",
	// 	Choices: []string{"Yes", "No"},
	// }).Prompt(yo)
	// fmt.Println()
	// (&prompt.Prompt{
	// 	Message: "How is life?",
	// 	Choices: []string{"Good", "Bad", "Great"},
	// 	Default: "Good",
	// }).Prompt(yo)
	// fmt.Println()
	input, err := (&prompt.Prompt{
		Choices:       []string{"Good", "Bad", "Great"},
		Attempts:      2,
		Default:       "Awesome",
		CaseSensitive: true,
	}).Prompt(yo)
	if err != nil {
		fmt.Println(err.Error())
	}

	fmt.Println(input + "\n\n")
	input2, err := (&query.Query{
		Message: "Enter a value yo",
	}).Query(yo)
	if err != nil {
		fmt.Println(err.Error())
	}
	fmt.Println("OUTPUT:\n(" + input2 + ")")
}
