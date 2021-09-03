package main

import (
	"github.com/bchadwic/yo/pkg/edit"
	"github.com/bchadwic/yo/pkg/prompt"
	"github.com/bchadwic/yo/yo"
)

func main() {
	// yo := yo.ProdYo()
	// // fmt.Println()
	// // (&prompt.Prompt{
	// // 	Message: "Do you like icecream?",
	// // 	Choices: []string{"Yes", "No"},
	// // }).Prompt(yo)
	// // fmt.Println()
	// // (&prompt.Prompt{
	// // 	Message: "How is life?",
	// // 	Choices: []string{"Good", "Bad", "Great"},
	// // 	Default: "Good",
	// // }).Prompt(yo)
	// // fmt.Println()
	// input, err := (&prompt.Prompt{
	// 	Choices:       []string{"Good", "Bad", "Great"},
	// 	Attempts:      2,
	// 	Default:       "Awesome",
	// 	CaseSensitive: true,
	// }).Prompt(yo)
	// if err != nil {
	// 	fmt.Println(err.Error())
	// }

	// fmt.Println(input + "\n\n")
	// input2, err := (&query.Query{
	// 	Message: "Enter a value yo",
	// }).Query(yo)
	// if err != nil {
	// 	fmt.Println(err.Error())
	// }
	// fmt.Println("OUTPUT:\n(" + input2 + ")")
	//exec.Command("which", "vim").Stdout =
	//exec.Cmd.Stdin = ""
	//yo,_,_,_ := yo.TestYo()
	//exec.Command("").Output() = yo.Out
	//fmt.Println(err)
	// cmd := exec.Command("vim", "test.txt")
	// cmd.Stdin = os.Stdin
	// cmd.Stdout = os.Stdout
	// err := cmd.Run()
	// fmt.Println(err)

	// os := runtime.GOOS
	// switch os {
	// case "windows":
	// 	fmt.Println("Windows")
	// case "darwin":
	// 	fmt.Println("MAC operating system")
	// case "linux":
	// 	fmt.Println("Linux")
	// default:
	// 	fmt.Printf("%s.\n", os)
	// }
	(&edit.Edit{}).Edit(yo.ProdYo())
	(&prompt.Prompt{
		Validate: func(s string) bool {
			return s == "hello"
		},
		Attempts: 29,
	}).Prompt(yo.ProdYo())

	//go launch(&Request{}).onSuccess(func() {}).onFailure(func() {})
	launch(&Request{}, onsuccess(onfailure))
}

type Request struct {
}

func onsuccess(f func()) {}

func onfailure() {}

func launch(r *Request, f func()) *Request {
	return &Request{}
}

func (r *Request) onSuccess(validate func()) *Request {
	return &Request{}
}

func (r *Request) onFailure(validate func()) *Request {
	return &Request{}
}
