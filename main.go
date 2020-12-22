package main

import (
	"os"
	"strconv"

	d1 "github.com/linhmtran168/adventofcode2020/day1"
)

type Advent interface {
	Solve()
}

func main() {
	args := os.Args[1:]
	probDay, err := strconv.Atoi(args[0])
	if err != nil {
		panic(err)
	}

	var problem Advent
	switch probDay {
	case 1:
		problem = &(d1.Day1{})
	default:
		panic("Invalid problem!")
	}

	problem.Solve()
}
