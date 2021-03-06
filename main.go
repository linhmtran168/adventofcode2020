package main

import (
	"fmt"
	"os"
	"strconv"
	"time"

	d1 "github.com/linhmtran168/adventofcode2020/day1"
	d2 "github.com/linhmtran168/adventofcode2020/day2"
	d3 "github.com/linhmtran168/adventofcode2020/day3"
	d4 "github.com/linhmtran168/adventofcode2020/day4"
	d5 "github.com/linhmtran168/adventofcode2020/day5"
	d6 "github.com/linhmtran168/adventofcode2020/day6"
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
	case 4:
		problem = &d4.Problem{}
	case 5:
		problem = &d5.Problem{}
	case 6:
		problem = &d6.Problem{}
	default:
		panic("Invalid problem!")
	}

	defer timeTrack(probDay)()
	problem.Solve()
}

func timeTrack(probDay int) func() {
	start := time.Now()
	return func() {
		fmt.Printf("\n=========\nProblem %d took %v\n", probDay, time.Since(start))
	}
}
