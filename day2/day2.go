package day2

import (
	"bufio"
	"errors"
	"fmt"
	"os"
	"strconv"
	"strings"
)

type Day2 struct{}

type Password struct {
	Min  int
	Max  int
	Char string
	Pass string
}

func (adv *Day2) Solve() {
	passwords, err := readInput("day2/input.txt")
	if err != nil {
		panic(err)
	}

	numCorrect := 0
	numCorrect2 := 0
	for _, v := range passwords {
		if isValid(v) {
			numCorrect += 1
		}

		if isValid2(v) {
			numCorrect2 += 1
		}
	}

	fmt.Printf("%d\n", numCorrect)
	fmt.Printf("%d\n", numCorrect2)
}

func readInput(filePath string) (passwords []Password, err error) {
	file, err := os.Open(filePath)
	if err != nil {
		return
	}
	defer file.Close()

	scanner := bufio.NewScanner(file)
	for scanner.Scan() {
		line := scanner.Text()
		password, innerErr := parsePassword(line)

		if innerErr != nil {
			return nil, innerErr
		}

		passwords = append(passwords, password)
	}

	return
}

func parsePassword(str string) (Password, error) {
	arr := strings.Split(str, " ")
	if len(arr) != 3 {
		return Password{}, errors.New("Invalid data")
	}

	limits := strings.Split(arr[0], "-")
	minLimit, err := strconv.Atoi(limits[0])
	if err != nil {
		return Password{}, errors.New("Invalid data")
	}
	maxLimit, err := strconv.Atoi(limits[1])
	if err != nil {
		return Password{}, errors.New("Invalid data")
	}

	return Password{minLimit, maxLimit, arr[1][:1], arr[2]}, nil
}

func isValid(pass Password) bool {
	numOccurrent := strings.Count(pass.Pass, pass.Char)

	if numOccurrent >= pass.Min && numOccurrent <= pass.Max {
		return true
	}

	return false
}

func isValid2(pass Password) bool {
	if string(pass.Pass[pass.Min-1]) == pass.Char && string(pass.Pass[pass.Max-1]) == pass.Char {
		return false
	}

	if string(pass.Pass[pass.Min-1]) != pass.Char && string(pass.Pass[pass.Max-1]) != pass.Char {
		return false
	}

	return true
}
