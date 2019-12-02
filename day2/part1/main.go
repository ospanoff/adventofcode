package main

import (
	"fmt"
	"io/ioutil"
	"strconv"
	"strings"
)

func check(err error) {
	if err != nil {
		panic(err)
	}
}

func intcode(input []int) {
	var halt = false
	for i := 0; i < len(input); i += 4 {
		switch input[i] {
		case 1:
			input[input[i+3]] = input[input[i+1]] + input[input[i+2]]
		case 2:
			input[input[i+3]] = input[input[i+1]] * input[input[i+2]]
		case 99:
			halt = true
		default:
			panic("Unexpected code")
		}
		if halt {
			break
		}
	}
}

func main() {
	var dataBin, err = ioutil.ReadFile("input.txt")
	check(err)
	var input = []int{}
	var numInt int
	for _, num := range strings.Split(string(dataBin), ",") {
		numInt, err = strconv.Atoi(strings.Trim(num, "\n"))
		check(err)
		input = append(input, numInt)
	}
	input[1] = 12
	input[2] = 2
	intcode(input)
	fmt.Println(input[0])
}
