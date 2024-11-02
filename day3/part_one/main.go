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
			if !matrix.HasSymbolNeighboor(row, column, engineMap) {
				continue
			}

			partNumber, lastIndex := matrix.ExtractPartNumber(row, column, engineMap)
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
