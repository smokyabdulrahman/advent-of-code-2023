package main

import (
	"advent-of-code-2023/day2/parser"
	"bufio"
	"fmt"
	"log"
	"os"
)

func main() {
	file, err := os.Open("input.txt")
	if err != nil {
		log.Fatal(err)
	}
	defer file.Close()

	scanner := bufio.NewScanner(file)
	var answer = 0
	for scanner.Scan() {
		line := scanner.Text()
		_, gameRuns := parser.ParseGame(line)
		maxRed := 0
		maxGreen := 0
		maxBlue := 0
		for _, cubeGroups := range parser.ParseRuns(gameRuns) {
			cubes := parser.ParseCubeGroup(cubeGroups)
			for _, cube := range cubes {
				number, color := parser.ParseCube(cube)

				switch color {
				case "red":
					if number > maxRed {
						maxRed = number
					}
				case "green":
					if number > maxGreen {
						maxGreen = number
					}
				case "blue":
					if number > maxBlue {
						maxBlue = number
					}
				}
			}

		}
		power := (maxRed * maxGreen * maxBlue)
		answer += power
	}

	fmt.Print(answer)
}
