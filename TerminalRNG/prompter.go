package main

import (
	"fmt"
	"github.com/c-bata/go-prompt"
	"time"
)

const INTERNET string = "internet"
const CPU string = "cpu"
const KEYBOARD string = "keyboard"
const DATA string = "data"

func completer(d prompt.Document) []prompt.Suggest {
	s := []prompt.Suggest{
		{Text: INTERNET, Description: "Use Website Latency as source of Randomness"},
		{Text: CPU, Description: "Use CPU Usage as source of Randomness"},
		{Text: KEYBOARD, Description: "Use Keyboard Input as source of Randomness"},
		{Text: DATA, Description: "Use data processing as source Randomness"},
	}
	return prompt.FilterHasPrefix(s, d.GetWordBeforeCursor(), true)
}

func main() {
	startTime := time.Now().UnixMicro()
	fmt.Println("Please select an option.")
	t := prompt.Input("> ", completer)
	fmt.Println("You selected " + t)
	duration := time.Now().UnixMicro() - startTime
	switch t {
	case INTERNET:
		for i := 0; i < 100; i++ {
			fmt.Println("Internet gives a random time of: ", Internet(duration))
		}
	}

}
