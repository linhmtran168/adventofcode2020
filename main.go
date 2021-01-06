package main

import (
	"os"
	"strconv"

	d1 "github.com/linhmtran168/adventofcode2020/day1"
	d2 "github.com/linhmtran168/adventofcode2020/day2"
	d3 "github.com/linhmtran168/adventofcode2020/day3"
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
		problem = &d1.Problem{}
	case 2:
		problem = &d2.Problem{}
	case 3:
		problem = &d3.Problem{}
	default:
		panic("Invalid problem!")
	}

	problem.Solve()
}
