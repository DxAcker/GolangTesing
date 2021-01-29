package src

import (
	"fmt"
	"regexp"
	"os"
)

func check(err error) {
	if err !=nil {
		fmt.Println(err.Error())
		os.Exit(1)
	}
}

func match(pattern string, text string) {
	matched, _ := regexp.Match(pattern, []byte(text))
	return matched
}

func main() {
	pattern := '^[A-Z][0-9]-[A-Z][0-9]$'
	fmt.Print("Enter your move: ")
	var move string
	fmt.Fscan(os.Stdin, &move)
	
}
