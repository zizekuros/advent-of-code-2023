package main

import (
	"fmt"
	"os"
	"strconv"

	day01 "github.com/zizekuros/advent-of-code-2023/days/day01"
)

func main() {

	if len(os.Args) < 4 {
		fmt.Println("Usage: go run main.go <day number> <part> <input file>")
		os.Exit(1)
	}

	day, err := strconv.Atoi(os.Args[1])
	if err != nil {

	}

	part, err := strconv.Atoi(os.Args[2])
	if err != nil {
		fmt.Printf("Invalid part number: %s\n", os.Args[2])
		os.Exit(1)
	}

	switch day {
	case 1:
		day01.Solve(part, os.Args[3])
	default:
		fmt.Printf("Day %d not implemented yet\n", day)
	}

}