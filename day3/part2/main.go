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
	var dataBin, err = ioutil.ReadFile("input.txt")
	check(err)
	var paths = strings.Split(string(dataBin), "\n")
	var pathA, pathB = paths[0], paths[1]

	var board = map[int][]int{}
	var lastX, lastY = 0, 0

	var distances = map[int][]int{}
	var lastDist = 0

	var minDist = 1000000

	for _, step := range strings.Split(pathA, ",") {
		var dir, numStepsStr = step[0], step[1:]
		var numSteps, err = strconv.Atoi(numStepsStr)
		check(err)

		for s := 0; s < numSteps; s++ {
			switch dir {
			case 'R':
				lastX++
			case 'U':
				lastY++
			case 'L':
				lastX--
			case 'D':
				lastY--
			}
			board[lastX] = append(board[lastX], lastY)
			lastDist++
			distances[lastX] = append(distances[lastX], lastDist)
		}
	}

	lastX, lastY = 0, 0
	lastDist = 0

	for _, step := range strings.Split(pathB, ",") {
		var dir, numStepsStr = step[0], step[1:]
		var numSteps, err = strconv.Atoi(numStepsStr)
		check(err)

		for s := 0; s < numSteps; s++ {
			switch dir {
			case 'R':
				lastX++
			case 'U':
				lastY++
			case 'L':
				lastX--
			case 'D':
				lastY--
			}
			lastDist++
			if ys, ok := board[lastX]; ok {
				for i, y := range ys {
					if lastY == y {
						var dist = distances[lastX][i] + lastDist
						if dist < minDist {
							minDist = dist
						}
					}
				}
			}
		}
	}
	fmt.Println(minDist)
}
