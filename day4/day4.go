package day4

import (
	"bufio"
	"fmt"
	"os"
	"regexp"
	"strconv"
	"strings"

	u "github.com/linhmtran168/adventofcode2020/utils"
)

type Problem struct{}

func (adv *Problem) Solve() {
	rawPassports, err := readInput("day4/input.txt")
	if err != nil {
		panic(err)
	}

	requiredAttrs := []string{"byr", "iyr", "eyr", "hgt", "hcl", "ecl", "pid"}
	requiredAttrRegexes := map[string]*regexp.Regexp{
		"hgt": regexp.MustCompile(`^(\d+)(cm|in)$`),
		"hcl": regexp.MustCompile(`^#[0-9a-f]{6}$`),
		"ecl": regexp.MustCompile(`^(amb|blu|brn|gry|grn|hzl|oth)$`),
		"pid": regexp.MustCompile(`^\d{9}$`),
	}

	isValueValid := func(attrName string, attrVal string) bool {
		switch attrName {
		case "byr":
			val, err := strconv.Atoi(attrVal)
			return err == nil && val >= 1920 && val <= 2002
		case "iyr":
			val, err := strconv.Atoi(attrVal)
			return err == nil && val >= 2010 && val <= 2020
		case "eyr":
			val, err := strconv.Atoi(attrVal)
			return err == nil && val >= 2020 && val <= 2030
		case "hgt":
			val := requiredAttrRegexes[attrName].FindStringSubmatch(attrVal)
			if val == nil || len(val) != 3 {
				return false
			}

			hgtVal, err := strconv.Atoi(val[1])
			if err != nil {
				return false
			}

			return (val[2] == "in" && hgtVal >= 59 && hgtVal <= 76) || (val[2] == "cm" && hgtVal >= 150 && hgtVal <= 193)
		case "hcl", "ecl", "pid":
			return requiredAttrRegexes[attrName].Match([]byte(attrVal))
		}

		return false
	}

	numValid := 0
	numRequiredAttrs := len(requiredAttrs)
	for _, p := range rawPassports {
		pAttrs := strings.Fields(p)
		var validAttrNames []string
		// passpost := ""
		for _, attr := range pAttrs {
			subAttr := strings.Split(attr, ":")
			attrName, attrVal := subAttr[0], subAttr[1]

			if _, isNameValid := u.FindString(requiredAttrs, attrName); isNameValid {
				if _, added := u.FindString(validAttrNames, attrName); !added && isValueValid(attrName, attrVal) {
					validAttrNames = append(validAttrNames, attrName)
					// passpost = passpost + fmt.Sprintf("%s - %s ", attrName, attrVal)
				}
			}
		}

		if numRequiredAttrs == len(validAttrNames) {
			numValid++
			// fmt.Printf("%s\n", passpost)
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
