package day5

import (
	"bufio"
	"fmt"
	"os"
)

type Problem struct{}

type seatCode struct {
	row    string
	column string
}

func (adv *Problem) Solve() {
	seatCodes, err := readInput("day5/input.txt")
	if err != nil {
		panic(err)
	}

	var findValue func(byte, string, int, int) int
	findValue = func(lowerCode byte, code string, lower int, upper int) int {
		if code[0] == lowerCode {
			if len(code) == 1 {
				return lower
			}

			return findValue(lowerCode, code[1:], lower, (lower+upper+1)/2-1)
		} else {
			if len(code) == 1 {
				return upper
			}

			return findValue(lowerCode, code[1:], (lower+upper+1)/2, upper)
		}
	}

	maxSeatID := 0
	seatHotcodes := make([]int, 128*8)
	for _, code := range seatCodes {
		row := findValue('F', code.row, 0, 127)
		col := findValue('L', code.column, 0, 7)
		seatID := row*8 + col

		seatHotcodes[seatID] = 1
		if seatID > maxSeatID {
			maxSeatID = seatID
		}

		fmt.Printf("%s%s:row %d, column %d, seat ID %d\n", code.row, code.column, row, col, seatID)
	}

	validSeatID := 0
	for idx, val := range seatHotcodes {
		if idx == 0 || idx == len(seatHotcodes)-1 {
			continue
		}

		if seatHotcodes[idx-1] != 0 && seatHotcodes[idx+1] != 0 && val == 0 {
			validSeatID = idx
		}
	}

	fmt.Printf("Max Seat ID: %d\n", maxSeatID)
	fmt.Printf("Valid Sead ID: %d\n", validSeatID)
}

func readInput(filePath string) (output []*seatCode, err error) {
	file, err := os.Open(filePath)
	if err != nil {
		return
	}
	defer file.Close()

	scanner := bufio.NewScanner(file)
	for scanner.Scan() {
		line := scanner.Text()
		code := &seatCode{line[0:7], line[7:]}
		output = append(output, code)
	}

	return
}
