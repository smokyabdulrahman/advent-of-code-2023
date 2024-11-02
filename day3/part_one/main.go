package main

import (
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

type point struct {
	x int
	y int
}

func hasSymbolNeighboor(x int, y int, matrix [][]rune) bool {
	// To imagine a mini perfect version of the matrix
	//   0 1 2
	// 0[0,0,0]
	// 1[0,0,0]
	// 2[0,0,0]
	neighboorLocations := []point{
		// Top row neighboors
		{x: -1, y: -1},
		{x: -1, y: 0},
		{x: -1, y: 1},
		// left neighboor
		{x: 0, y: -1},
		// right neighboor
		{x: 0, y: 1},
		// bottom neighboors
		{x: 1, y: -1},
		{x: 1, y: 0},
		{x: 1, y: 1},
	}

	for _, pointDiff := range neighboorLocations {
		neighboor := point{
			x: pointDiff.x + x,
			y: pointDiff.y + y,
		}
		if neighboor.x >= len(matrix) || neighboor.x < 0 {
			continue
		}
		if neighboor.y >= len(matrix[x]) || neighboor.y < 0 {
			continue
		}

		neighboorValue := matrix[neighboor.x][neighboor.y]
		// handle symbols
		if neighboorValue >= 33 && neighboorValue <= 45 {
			return true
		}

		// handle / skipping 46 which is \.
		if neighboorValue == 47 {
			return true
		}

		if neighboorValue >= 58 && neighboorValue <= 64 {
			return true
		}
	}

	return false
}

func extractPartNumber(row int, column int, matrix [][]rune) (number int, lastIndex int) {

	for column >= 0 && unicode.IsDigit(matrix[row][column]) {
		column--
	}

	// did we go back too far?!
	if column < 0 || !unicode.IsDigit(matrix[row][column]) {
		column++
	}

	number = int(matrix[row][column] - '0')
	column++

	for column < len(matrix[row]) && unicode.IsDigit(matrix[row][column]) {
		number *= 10 // push numbers 1 place to the left to appen new number
		number += int(matrix[row][column] - '0')
		column++
	}

	// did we go forward too far?!
	if column >= len(matrix[row]) || !unicode.IsDigit(matrix[row][column]) {
		column--
	}

	return number, column
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
	for row := 0; row < len(engineMap); row++ {
		for column := 0; column < len(engineMap[row]); column++ {
			char := engineMap[row][column]

			// ignore none numbers
			if !unicode.IsDigit(char) {
				continue
			}

			// ignore number if it's not a neighboor to a symbol
			if !hasSymbolNeighboor(row, column, engineMap) {
				continue
			}

			partNumber, lastIndex := extractPartNumber(row, column, engineMap)
			// and skip to number end for next iteration
			column = lastIndex + 1

			// add number to sum
			answer += partNumber
		}
	}
	// if symbol found
	// look for number beginning

	fmt.Print(answer)
}
