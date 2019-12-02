package main

import (
	"fmt"
	"io/ioutil"
	"os"
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

func parseInput(dataBin []byte) []int {
	var input = []int{}
	for _, num := range strings.Split(string(dataBin), ",") {
		var numInt, err = strconv.Atoi(strings.Trim(num, "\n"))
		check(err)
		input = append(input, numInt)
	}
	return input
}

func main() {
	var dataBin, err = ioutil.ReadFile("input.txt")
	check(err)

	var initialState = parseInput(dataBin)

	var output = 19690720
	for i := 0; i < 100; i++ {
		for j := 0; j < 100; j++ {
			var input = append([]int(nil), initialState...)
			input[1] = i
			input[2] = j

			intcode(input)
			if input[0] == output {
				fmt.Println(100*i + j)
				os.Exit(0)
			}
		}
	}
}
