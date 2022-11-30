package day5

import (
	"bufio"
	"fmt"
	"os"
)

type Problem struct{}

func (adv *Problem) Solve() {
	answers, err := readInput("day6/input.txt")
	if err != nil {
		panic(err)
	}

	sum1 := partOne(answers)
	sum2 := partTwo(answers)
	fmt.Printf("Sum of yes answers: %d\n", sum1)
	fmt.Printf("Sum of yes answers: %d\n", sum2)
}

func partOne(answers [][]string) int {
	sum := 0
	for _, group := range answers {
		uniqAns := make(map[rune]int)
		for _, people := range group {
			for _, ans := range people {
				uniqAns[ans]++
			}
		}

		sum += len(uniqAns)
	}

	return sum
}

func partTwo(answers [][]string) int {
	sum := 0
	for _, group := range answers {
		uniqAnsGrp := make(map[rune]int)
		for _, people := range group {
			for _, ans := range people {
				uniqAnsGrp[ans]++
			}
		}

		numPeople := len(group)
		for _, v := range uniqAnsGrp {
			if v == numPeople {
				sum += 1
			}
		}
	}

	return sum
}

func readInput(filePath string) (output [][]string, err error) {
	file, err := os.Open(filePath)
	if err != nil {
		return
	}
	defer file.Close()

	scanner := bufio.NewScanner(file)
	var group []string
	for scanner.Scan() {
		line := scanner.Text()
		if line != "" {
			group = append(group, line)

		} else {
			output = append(output, group)
			group = nil
		}
	}

	if group != nil {
		output = append(output, group)
	}

	fmt.Printf("===================\n")
	return
}
