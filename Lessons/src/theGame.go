package src

import (
	"fmt"
	"os"
)

func theGame() {
	var firstnum int
	fmt.Print("Enter first number: ")
	fmt.Fscan(os.Stdin, &firstnum)
	lastnum := 9 - firstnum

	fmt.Printf("%v9%v", firstnum, lastnum)
}
