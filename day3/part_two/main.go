package main

import (
	"advent-of-code-2023/utils/matrix"
	"bufio"
	"fmt"
	"io"
	"iter"
	"log"
	"os"
	"slices"
	"unicode"
)

func Lines(file *os.File) iter.Seq[string] {
	scanner := bufio.NewScanner(file)
	return func(yield func(string) bool) {
		for scanner.Scan() {
			if err := scanner.Err(); err != nil && err != io.EOF {
				log.Fatal(err)
			}
			if !yield(scanner.Text()) {
				return
			}
		}
	}
}

func getStarNeighboors(x int, y int, matrixx [][]rune) []matrix.Point {
	// To imagine a mini perfect version of the matrix
	//   0 1 2
	// 0[0,0,0]
	// 1[0,0,0]
	// 2[0,0,0]
	neighboorLocations := []matrix.Point{
		// Top row neighboors
		{X: -1, Y: -1},
		{X: -1, Y: 0},
		{X: -1, Y: 1},
		// left neighboor
		{X: 0, Y: -1},
		// right neighboor
		{X: 0, Y: 1},
		// bottom neighboors
		{X: 1, Y: -1},
		{X: 1, Y: 0},
		{X: 1, Y: 1},
	}

	neighboors := make([]matrix.Point, 0, 8)
	for _, pointDiff := range neighboorLocations {
		neighboor := matrix.Point{
			X: pointDiff.X + x,
			Y: pointDiff.Y + y,
		}
		if neighboor.X >= len(matrixx) || neighboor.X < 0 {
			continue
		}
		if neighboor.Y >= len(matrixx[x]) || neighboor.Y < 0 {
			continue
		}

		neighboorValue := matrixx[neighboor.X][neighboor.Y]

		if neighboorValue != 42 {
			continue
		}

		neighboors = append(neighboors, neighboor)
	}

	return neighboors
}

func main() {

	file, err := os.Open("input.txt")
	if err != nil {
		log.Fatal(err)
	}
	defer file.Close()
	lines := slices.Collect(Lines(file))

	// Construct 2d array
	engineMap := make([][]rune, len(lines))
	for i := range engineMap {
		engineMap[i] = []rune(lines[i])
	}
	// loop through array, if digit found, look for neighbour symbol
	answer := 0
	starsMap := make(map[string][]int)
	for row := 0; row < len(engineMap); row++ {
		for column := 0; column < len(engineMap[row]); column++ {
			char := engineMap[row][column]

			// create a map[point][]int <- store * only
			// if char is not number skip
			if !unicode.IsDigit(char) {
				continue
			}
			// get all star neighboors
			starNeighboors := getStarNeighboors(row, column, engineMap)
			if len(starNeighboors) == 0 {
				continue
			}
			// extract number
			partNumber, lastIndex := matrix.ExtractPartNumber(row, column, engineMap)
			// and skip to number end for next iteration
			column = lastIndex + 1

			// push number to each neighboor in the map
			for _, starNeighboor := range starNeighboors {
				index := string(starNeighboor.X) + "," + string(starNeighboor.Y)
				starNumbers, exist := starsMap[index]
				// we don't need to add more, it's already outside of the range we need
				if len(starNumbers) > 2 {
					continue
				}
				if !exist {
					starNumbers = make([]int, 0, 3)
				}
				starNumbers = append(starNumbers, partNumber)
				starsMap[index] = starNumbers
			}

		}

	}
	for _, v := range starsMap {
		if len(v) == 2 {
			gearRatio := 1
			for _, num := range v {
				gearRatio *= num
			}
			answer += gearRatio
		}
	}

	fmt.Print(answer)
}
