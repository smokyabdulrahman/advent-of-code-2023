package main

import (
	"bufio"
	"fmt"
	"log"
	"os"
	"strconv"
	"strings"
)

var maxCubes = map[string]int{
	"red":   12,
	"green": 13,
	"blue":  14,
}

func parseGame(line string) (gameId int, gameRuns string) {
	game := strings.Split(line, ":")
	gameIdPart := strings.Split(game[0], " ")

	gameId, _ = strconv.Atoi(gameIdPart[1])
	gameRuns = strings.TrimSpace(game[1])
	return
}

func isAPossibleGame(gameRuns string) bool {
	// 3 blue, 4 red; 1 red, 2 green, 6 blue; 2 green
	runs := strings.Split(gameRuns, ";")
	for _, run := range runs {
		if !isAPossibleRun(strings.TrimSpace(run)) {
			return false
		}
	}

	return true
}

func isAPossibleRun(run string) bool {
	// 1 red, 2 green, 6 blue
	cubes := strings.Split(run, ", ")

	for _, cube := range cubes {
		cubeInfo := strings.Split(cube, " ")
		number, _ := strconv.Atoi(cubeInfo[0])
		color := cubeInfo[1]
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
		gameId, gameRuns := parseGame(line)

		if isAPossibleGame(gameRuns) {
			answer += gameId
		}
	}

	fmt.Print(answer)
}
