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

func drawBoard(board map[int][]int) {
	var maxX, maxY = 0, 0
	var minX, minY = 0, 0
	for x, ys := range board {
		if x > maxX {
			maxX = x
		}
		if x < minX {
			minX = x
		}
		for _, y := range ys {
			if y > maxY {
				maxY = y
			}
			if y < minY {
				minY = y
			}
		}
	}
	for y := maxY; y >= minY; y-- {
		for x := minX; x <= maxX; x++ {
			var exists = false
			if ys, ok := board[x]; ok {
				for _, _y := range ys {
					if _y == y {
						fmt.Print("*")
						exists = true
						break
					}
				}
			}
			if !exists {
				fmt.Print(".")
			}
		}
		fmt.Println()
	}
}

func abs(x int) int {
	if x < 0 {
		return -x
	}
	return x
}

func main() {
	var dataBin, err = ioutil.ReadFile("input.txt")
	check(err)
	var paths = strings.Split(string(dataBin), "\n")
	var pathA, pathB = paths[0], paths[1]

	var board = map[int][]int{}
	var lastX, lastY = 0, 0

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
		}
	}

	lastX, lastY = 0, 0
	// drawBoard(board)

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
			if ys, ok := board[lastX]; ok {
				for _, y := range ys {
					if lastY == y {
						var dist = abs(lastX) + abs(lastY)
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
