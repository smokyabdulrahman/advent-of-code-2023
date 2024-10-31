package part_one

import (
	"advent-of-code-2023/day2/parser"
	"bufio"
	"fmt"
	"log"
	"os"
	"strings"
)

var maxCubes = map[string]int{
	"red":   12,
	"green": 13,
	"blue":  14,
}

func isAPossibleGame(gameRuns string) bool {
	cubeGroups := parser.ParseRuns(gameRuns)
	for _, cubeGroup := range cubeGroups {
		if !isAPossibleRun(strings.TrimSpace(cubeGroup)) {
			return false
		}
	}

	return true
}

func isAPossibleRun(cubeGroup string) bool {
	cubes := parser.ParseCubeGroup(cubeGroup)

	for _, cube := range cubes {
		number, color := parser.ParseCube(cube)
		if number > maxCubes[color] {
			return false
		}
	}

	return true
}

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
		gameId, gameRuns := parser.ParseGame(line)

		if isAPossibleGame(gameRuns) {
			answer += gameId
		}
	}

	fmt.Print(answer)
}
