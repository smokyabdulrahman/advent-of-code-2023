package main

import (
	"bufio"
	"fmt"
	"log"
	"os"
	"strings"
)

func extract(line string) (int, bool) {
	if line[0] >= '0' && line[0] <= '9' {
		return int(line[0] - '0'), true
	}
	if strings.HasPrefix(line, "zero") {
		return 0, true
	}
	if strings.HasPrefix(line, "one") {
		return 1, true
	}
	if strings.HasPrefix(line, "two") {
		return 2, true
	}
	if strings.HasPrefix(line, "three") {
		return 3, true
	}
	if strings.HasPrefix(line, "four") {
		return 4, true
	}
	if strings.HasPrefix(line, "five") {
		return 5, true
	}
	if strings.HasPrefix(line, "six") {
		return 6, true
	}
	if strings.HasPrefix(line, "seven") {
		return 7, true
	}
	if strings.HasPrefix(line, "eight") {
		return 8, true
	}
	if strings.HasPrefix(line, "nine") {
		return 9, true
	}
	return 0, false
}

func main() {
	file, err := os.Open("./input.txt")
	if err != nil {
		log.Fatal(err)
	}
	defer file.Close()

	scanner := bufio.NewScanner(file)
	var answer int

	for scanner.Scan() {
		line := scanner.Text()
		var num int
		for i := 0; i < len(line); i++ {
			if value, ok := extract(line[i:]); ok {
				num = value * 10
				break
			}
		}

		for i := len(line) - 1; i >= 0; i-- {
			if value, ok := extract(line[i:]); ok {
				num += value
				break
			}
		}

		answer += num
	}
	fmt.Print(answer)
}
