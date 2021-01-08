package day4

import (
	"bufio"
	"fmt"
	"os"
	"strings"

	u "github.com/linhmtran168/adventofcode2020/utils"
)

type Problem struct{}

func (adv *Problem) Solve() {
	rawPassports, err := readInput("day4/input.txt")
	if err != nil {
		panic(err)
	}

	numValid := 0
	requiredAttrs := []string{"byr", "iyr", "eyr", "hgt", "hcl", "ecl", "pid"}
	numRequiredAttrs := len(requiredAttrs)
	for _, p := range rawPassports {
		pAttrs := strings.Fields(p)
		var validAttrNames []string
		for _, attr := range pAttrs {
			subAttr := strings.Split(attr, ":")
			attrName, _ := subAttr[0], subAttr[1]

			if _, isValid := u.FindString(requiredAttrs, attrName); isValid {
				if _, added := u.FindString(validAttrNames, attrName); !added {
					validAttrNames = append(validAttrNames, attrName)
				}
			}

		}

		if numRequiredAttrs == len(validAttrNames) {
			numValid++
		}
	}

	fmt.Printf("Num passport valid: %d\n", numValid)
}

func readInput(filePath string) (output []string, err error) {
	file, err := os.Open(filePath)
	if err != nil {
		return
	}
	defer file.Close()

	scanDouble := func(data []byte, atEOF bool) (advance int, token []byte, err error) {
		dataLen := len(data)
		for i := 0; i < dataLen-1; i++ {
			if data[i] == '\n' && data[i+1] == '\n' {
				return i + 2, data[:i], nil
			}
		}

		if atEOF && dataLen > 0 {
			return dataLen, data, nil
		}

		return 0, nil, nil
	}

	scanner := bufio.NewScanner(file)
	scanner.Split(scanDouble)
	for scanner.Scan() {
		line := scanner.Text()
		output = append(output, strings.TrimSpace(line))
	}

	return
}
