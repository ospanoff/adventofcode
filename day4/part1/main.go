package main

import "fmt"

// It is the easy direct solution

func check(err error) {
	if err != nil {
		panic(err)
	}
}

const (
	valid = iota
	noDouble
	decreasingPair
)

func checkCriterias(num int) int {
	var metDouble = false
	var lastDigit = num % 10
	num /= 10
	for num > 0 {
		var digit = num % 10
		if digit == lastDigit {
			metDouble = true
		}
		if digit > lastDigit {
			return decreasingPair
		}
		num /= 10
		lastDigit = digit
	}
	if !metDouble {
		return noDouble
	}
	return valid
}

func main() {
	var min, max = 254032, 789860
	var validCombs = 0
	for n := min; n <= max; n++ {
		if checkCriterias(n) == valid {
			validCombs++
		}
	}
	fmt.Println(validCombs)
}
