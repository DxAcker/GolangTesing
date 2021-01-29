package main

import (
	"fmt"
	"math"
	"os"
	"regexp"
)

func main() {
	var move string
	fmt.Print("Введите ваш ход: ")
	fmt.Fscan(os.Stdin, &move)

	matched, _ := regexp.MatchString("^[A-H][1-8]\\-[A-H][1-8]$", move)

	if !matched {
		fmt.Println("ERR")
		os.Exit(1)
	}

	razn := int(move[0]) - int(move[3])
	razb := int(move[1]) - int(move[4])

	var square = math.Abs(float64(razn) * float64(razb))

	if square == 2 {
		fmt.Println("YES")
	} else {
		fmt.Println("NO")
	}

}
