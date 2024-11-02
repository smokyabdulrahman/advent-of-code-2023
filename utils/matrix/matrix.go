package matrix

import "unicode"

type Point struct {
	X int
	Y int
}

func HasSymbolNeighboor(x int, y int, matrix [][]rune) bool {
	// To imagine a mini perfect version of the matrix
	//   0 1 2
	// 0[0,0,0]
	// 1[0,0,0]
	// 2[0,0,0]
	neighboorLocations := []Point{
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

	for _, pointDiff := range neighboorLocations {
		neighboor := Point{
			X: pointDiff.X + x,
			Y: pointDiff.Y + y,
		}
		if neighboor.X >= len(matrix) || neighboor.X < 0 {
			continue
		}
		if neighboor.Y >= len(matrix[x]) || neighboor.Y < 0 {
			continue
		}

		neighboorValue := matrix[neighboor.X][neighboor.Y]
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

func ExtractPartNumber(row int, column int, matrix [][]rune) (number int, lastIndex int) {

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
