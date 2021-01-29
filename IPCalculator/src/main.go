package main

import (
	"fmt"
	"math"
	"os"
	"strconv"
	"strings"
)

func main() {
	arguments := os.Args
	if len(arguments) == 1 {
		fmt.Println("The program needs to have some arguments!")
		fmt.Println("Arguments: class of IP, count of subnets")
		os.Exit(1)
	}
	ipclass, subnetСount := arguments[1], arguments[2]
	var netbits int
	if ipclass == "A" {
		netbits = 8
	} else if ipclass == "B" {
		netbits = 16
	} else if ipclass == "C" {
		netbits = 24
	} else {
		fmt.Println("Entered class is invalid!")
		os.Exit(1)
	}

	subnet, err := strconv.Atoi(subnetСount)
	if err != nil {
		fmt.Println("Count of subnets is invalid!")
		os.Exit(1)
	}
	var subnetbits float64 = 0
	for math.Pow(2, subnetbits) < float64(subnet) {
		subnetbits++
	}
	hostbits := 32 - int(subnetbits) - netbits

	var totalcount int64 = int64(math.Pow(2, subnetbits+float64(netbits)))

	fmt.Println("Mask:\t", ipToDec(maskGen(int(subnetbits)+netbits)))
	fmt.Println("subnetbits:", subnetbits, "\tnetbits:", netbits, "\thostbits:", hostbits)
	fmt.Println("You can use\t", math.Pow(2, float64(hostbits))-2, "hosts\t", int64(math.Pow(2, float64(netbits))), "networks with", math.Pow(2, subnetbits), "subnets each")
	fmt.Println("\tand total", totalcount, "networks")
	// fmt.Println("Binary mask", maskGen(int(subnetbits)+netbits))
}

func maskGen(bits int) string {
	var mask string
	for i := 1; len(mask) < 35; i++ {
		if bits > 0 {
			mask = mask + "1"
		} else {
			mask = mask + "0"
		}
		if i%8 == 0 && i < 32 {
			mask = mask + "."
		}
		bits--
	}
	return mask
}

func ipToDec(ip string) string {
	var output string
	var ParseErr error
	var i int64
	slices := strings.Split(ip, ".")
	for j, slice := range slices {
		i, ParseErr = strconv.ParseInt(slice, 2, 16)
		if j >= 3 {
			output = output + strconv.Itoa(int(i))
		} else {
			output = output + strconv.Itoa(int(i)) + "."
		}
	}
	if ParseErr != nil {
		fmt.Println(ParseErr)
	}
	return output
}
