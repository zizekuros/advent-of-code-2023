package day02

import (
	"fmt"
	"os"
	"regexp"
	"strconv"
	"strings"
)

func Solve(part int, filePath string) {

	if part != 1 {
		fmt.Printf("Invalid part number: %v\n", part)
		os.Exit(1)
	}

	content, err := os.ReadFile(filePath)
	if err != nil {
		fmt.Println("Error reading file:", err)
		os.Exit(1)
	}

	redCubes := 12
	greenCubes := 13
	blueCubes := 14

	sumOfPossibleGames := 0

	lines := strings.Split(string(content), "\n")

	colorRe := regexp.MustCompile(`(\d+) (red|green|blue)`)
	gameIDRe := regexp.MustCompile(`Game (\d+):`)

	var gameID int
	for _, line := range lines {
		gameIDMatch := gameIDRe.FindStringSubmatch(line)
		if len(gameIDMatch) == 2 {
			gameID, _ = strconv.Atoi(gameIDMatch[1])
		}

		reveals := strings.Split(line, ";")
		colorCounts := make(map[string]int)

		for _, reveal := range reveals {
			colorMatches := colorRe.FindAllStringSubmatch(reveal, -1)

			for _, match := range colorMatches {
				count, _ := strconv.Atoi(match[1])
				color := match[2]
				if count > colorCounts[color] {
					colorCounts[color] = count
				}
			}
		}

		redCount := colorCounts["red"]
		greenCount := colorCounts["green"]
		blueCount := colorCounts["blue"]

		if redCount > 0 || greenCount > 0 || blueCount > 0 {
			if redCount <= redCubes && greenCount <= greenCubes && blueCount <= blueCubes {
				fmt.Printf("Game %v is possible\n", gameID)
				sumOfPossibleGames += gameID
			}
		}

	}

	fmt.Println(sumOfPossibleGames)
}
