package main

import (
	"bufio"
	"fmt"
	"io"
	"iter"
	"log"
	"os"
	"slices"
	"strconv"
	"strings"
)

func calculateSlope(x float64, y float64, vX float64, vY float64) float64 {
	return ((y + vY) - y) / ((x + vX) - x)
}

func calculateYIntercept(x float64, y float64, m float64) float64 {
	return y - m*x
}

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

type linearFunction struct {
	m float64
	c float64
}

func main() {
	file, err := os.Open("input-test.txt")
	if err != nil {
		log.Fatal(err)
	}
	defer file.Close()

	lines := slices.Collect(Lines(file))
	answer := 0
	functions := make([]linearFunction, 0, len(lines))
	for _, line := range lines {
		// Positions
		firstComma := strings.Index(line, ",")
		x, _ := strconv.ParseFloat(strings.TrimSpace(line[:firstComma]), 10)

		secondComma := strings.Index(line[firstComma+2:], ",") + 2 + firstComma
		y, _ := strconv.ParseFloat(strings.TrimSpace(line[firstComma+2:secondComma]), 10)

		firstAt := strings.Index(line[secondComma+2:], "@") + 2 + secondComma
		//z, _ := strconv.ParseFloat(strings.TrimSpace(line[secondComma+2:firstAt]), 10)

		// Velocity
		vFirstComma := strings.Index(line[firstAt+2:], ",") + 2 + firstAt
		vX, _ := strconv.ParseFloat(strings.TrimSpace(line[firstAt+2:vFirstComma]), 10)

		vSecondComma := strings.Index(line[vFirstComma+2:], ",") + 2 + vFirstComma
		vY, _ := strconv.ParseFloat(strings.TrimSpace(line[vFirstComma+2:vSecondComma]), 10)

		m := calculateSlope(x, y, vX, vY)
		c := calculateYIntercept(x, y, m)
		functions = append(functions, linearFunction{m: m, c: c})
	}
	// do some nested loops here and add 1 for each unique intersection combo

	fmt.Print(answer)
}
