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

var lowerLimit float64 = 200000000000000
var maxLimit float64 = 400000000000000

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

func parseInputToLinearFunctions(lines []string) []linearFunction {
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
		functions = append(functions, linearFunction{
			m:     m,
			c:     c,
			initX: x,
			initY: y,
			vX:    vX,
			vY:    vY,
		})
	}

	return functions
}

type linearFunction struct {
	m     float64
	c     float64
	initX float64
	initY float64
	vX    float64
	vY    float64
}

func main() {
	file, err := os.Open("input.txt")
	if err != nil {
		log.Fatal(err)
	}
	defer file.Close()

	lines := slices.Collect(Lines(file))
	functions := parseInputToLinearFunctions(lines)

	answer := 0
	// do some nested loops here and add 1 for each unique intersection combo
	comparisonSet := make(map[string]bool)
	for outerIndex, function1 := range functions {
		for innerIndex, function2 := range functions {
			// skip comparing to self
			if outerIndex == innerIndex {
				continue
			}
			// skip comparing if already compared
			index := fmt.Sprintf("%d-%d", outerIndex, innerIndex)
			if _, exist := comparisonSet[index]; exist {
				continue
			}
			index = fmt.Sprintf("%d-%d", innerIndex, outerIndex)
			if _, exist := comparisonSet[index]; exist {
				continue
			}

			// skip if slopes are equal => lines are parallel
			if function1.m == function2.m {
				continue
			}
			// we need to test that the intersection is going to be inside a test area
			// m1x1 + c1 = m2x2 + c2
			// m1x1 - m2x2 = c2 - c1
			// (m1 - m2)x = c2 - c1
			// x = c2 - c1 / m1 - m2
			xIntersect := (function2.c - function1.c) / (function1.m - function2.m)
			// Check that the xIntersect is an intersection that happens in the future
			if function1.vX < 0 && xIntersect > function1.initX {
				continue
			}
			if function1.vX > 0 && xIntersect < function1.initX {
				continue
			}
			if function2.vX < 0 && xIntersect > function2.initX {
				continue
			}
			if function2.vX > 0 && xIntersect < function2.initX {
				continue
			}

			if xIntersect < lowerLimit || xIntersect > maxLimit {
				continue
			}

			yIntersect := function1.m*xIntersect + function1.c
			// Check that the yIntersect is an intersection that happens in the future
			if function1.vY < 0 && yIntersect > function1.initY {
				continue
			}
			if function1.vY > 0 && yIntersect < function1.initY {
				continue
			}
			if function2.vY < 0 && yIntersect > function2.initY {
				continue
			}
			if function2.vY > 0 && yIntersect < function2.initY {
				continue
			}

			if yIntersect < lowerLimit || yIntersect > maxLimit {
				continue
			}

			fmt.Printf(
				"func1 => y = %fx + %f, vX: %f, vY: %f, initX: %f, initY: %f\n",
				function1.m,
				function1.c,
				function1.vX,
				function1.vY,
				function1.initX,
				function1.initY,
			)
			fmt.Printf(
				"func2 => y = %fx + %f, vX: %f, vY: %f, initX: %f, initY: %f\n",
				function2.m,
				function2.c,
				function2.vX,
				function2.vY,
				function2.initX,
				function2.initY,
			)
			fmt.Printf("xIntersect: %f, yIntersect: %f\n", xIntersect, yIntersect)
			fmt.Println()
			comparisonSet[index] = true
			answer += 1
		}
	}

	fmt.Print(answer)
}
