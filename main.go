package main

import (
	"fmt"
	"os"
	"strconv"

	day01 "github.com/zizekuros/advent-of-code-2023/days/day01"
)

func main() {

	if len(os.Args) < 2 {
		fmt.Println("Usage: go run main.go <day number>")
		os.Exit(1)
	}

	day, err := strconv.Atoi(os.Args[1])
	if err != nil {
		fmt.Printf("Invalid day number: %s\n", os.Args[1])
		os.Exit(1)
	}

	switch day {
	case 1:
		day01.Solve()
	default:
		fmt.Printf("Day %d not implemented yet\n", day)
	}

}
