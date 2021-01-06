package day4

import (
	"bufio"
	"fmt"
	"os"
)

type Problem struct{}

func (adv *Problem) Solve() {
	toboggan, err := readInput("day3/input.txt")
	if err != nil {
		panic(err)
	}
	listSteps := [][]int{{1, 1}, {3, 1}, {5, 1}, {7, 1}, {1, 2}}

	res := 1
	for _, s := range listSteps {
		numTree := walkField(s[0], s[1], toboggan)
		fmt.Printf("Number of tree: %d\n", numTree)
		res *= numTree
	}

	fmt.Printf("Final result: %d", res)
}

func walkField(stepX int, stepY int, fields [][]int) int {
	tHeight := len(fields)
	tWidth := len(fields[0])

	numTree := 0
	posX := 0
	for i := 0; i < tHeight; i += stepY {
		if fields[i][posX] == 1 {
			numTree++
		}

		newPosX := posX + stepX

		if newPosX >= tWidth {
			posX = newPosX - tWidth
		} else {
			posX = newPosX
		}
	}

	return numTree
}

func readInput(filePath string) (toboggan [][]int, err error) {
	file, err := os.Open(filePath)
	if err != nil {
		return
	}
	defer file.Close()

	scanner := bufio.NewScanner(file)
	for scanner.Scan() {
		line := scanner.Text()
		var lineArr []int
		for _, char := range line {
			if char == '.' {
				lineArr = append(lineArr, 0)
			} else {
				lineArr = append(lineArr, 1)
			}
		}

		toboggan = append(toboggan, lineArr)
	}

	return
}
