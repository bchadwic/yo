# yo
yo is a simple io library

## Usage

```go
yo := yo.ProdYo()

input, err := (&prompt.Prompt{
		Message:       "What is your favorite color?",
		Validate:      isValidColor,
		Attempts:      3,
		CaseSensitive: false,
		Choices:       []string{"Red", "Green", "Blue"},
		Default:       "Yellow",
}).Prompt(yo)

fmt.Println(input)
```

```
What is your favorite color? (Red, Green, Blue) [Yellow]: Light Blueish Green
enter a valid choice: I don't know
enter a valid choice: red
Red
```
