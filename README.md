# yo
yo is a simple io library

## Usage

```go
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
		Preface:    "Writing contents to " + fileName + "\nContents: \n---\n" + fileContents + "\n---",
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
```

```
Type the name of the file you want to make [temp]: 
Enter the contents of the new file 'temp'
return twice to save and quit: 

TODO:
- Create a better README


Writing contents to temp
Contents: 
---
TODO:
- Create a better README
---
Reading contents from temp
file contents matched read contents
```
