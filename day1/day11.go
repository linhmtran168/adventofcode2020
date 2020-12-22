package day1

import (
	"bufio"
	"fmt"
	"os"
	"strconv"
)

type Day1 struct{}

func readInput(filePath string) (numLst []int, err error) {
	file, err := os.Open(filePath)
	if err != nil {
		return
	}
	defer file.Close()

	scanner := bufio.NewScanner(file)
	for scanner.Scan() {
		inputNum, innerErr := strconv.Atoi(scanner.Text())
		if innerErr != nil {
			return nil, innerErr
		}

		numLst = append(numLst, inputNum)
	}

	return
}

func (adv *Day1) Solve() {
	numbers, err := readInput("day1/input.txt")

	if err != nil {
		panic(err)
	}

	var checkMap = make(map[int]int)

	for _, num := range numbers {
		if val, ok := checkMap[num]; ok {
			fmt.Printf("%d * %d: %d\n", num, val, num*val)
			break
		} else {
			checkMap[2020-num] = num
		}
	}

	var checkMap2 = make(map[int][]int)
	var numLen = len(numbers)
out:
	for i := 0; i < numLen; i++ {
		num1 := numbers[i]
		for j := i + 1; j < numLen; j++ {
			num2 := numbers[j]
			neededNum := 2020 - num1 - num2
			if neededNum >= 2020 {
				continue
			}

			if val, ok := checkMap2[num2]; ok {
				fmt.Printf("%d * %d * %d: %d\n", num2, val[0], val[1], num2*val[0]*val[1])
				break out
			} else {
				checkMap2[neededNum] = []int{num1, num2}
			}
		}
	}
}
