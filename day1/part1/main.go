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

func main() {
	var dataB, err = ioutil.ReadFile("input.txt")
	check(err)
	var data = string(dataB)
	var mass, totalFuel int
	for _, line := range strings.Split(data, "\n") {
		if len(line) > 0 {
			mass, err = strconv.Atoi(line)
			check(err)
			totalFuel += mass/3 - 2
		}
	}
	fmt.Println(totalFuel)
}
