package src

import (
	"fmt"
	"os"
)

func sumtoN() {
	var N int
	fmt.Fscan(os.Stdin, &N)

	i := 1
	sum := 0
	for i <= N {
		sum += i
		i++
	}
	fmt.Println(sum)
}
