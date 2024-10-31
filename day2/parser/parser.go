package parser

import (
	"strconv"
	"strings"
)

var MaxCubes = map[string]int{
	"red":   12,
	"green": 13,
	"blue":  14,
}

func ParseGame(line string) (gameId int, gameRuns string) {
	game := strings.Split(line, ": ")
	gameIdPart := strings.Split(game[0], " ")

	gameId, _ = strconv.Atoi(gameIdPart[1])
	gameRuns = strings.TrimSpace(game[1])
	return
}

func ParseRuns(run string) (cubeGroups []string) {
	cubeGroups = strings.Split(run, "; ")
	return
}

func ParseCubeGroup(cubeGroup string) (cubes []string) {
	cubes = strings.Split(cubeGroup, ", ")
	return
}

func ParseCube(cube string) (number int, color string) {
	cubeInfo := strings.Split(cube, " ")
	number, _ = strconv.Atoi(cubeInfo[0])
	color = cubeInfo[1]
	return
}
