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

func getFuel(mass int) int {
	var fuel = mass
	var totalFuel int
	for ; fuel > 0; fuel = fuel/3 - 2 {
		totalFuel += fuel
	}
	return totalFuel - mass
}

func main() {
	var dataB, err = ioutil.ReadFile("input.txt")
	check(err)
	var data = string(dataB)
	var mass, totalFuel int
	for _, line := range strings.Split(data, "\n") {
		if len(line) > 0 {
			mass, err = strconv.Atoi(line)
			check(err)
			totalFuel += getFuel(mass)
		}
	}
	fmt.Println(totalFuel)
}
