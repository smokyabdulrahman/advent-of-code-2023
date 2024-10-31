package main

import (
	"bufio"
	"fmt"
	"io"
	"log"
	"os"
	"strconv"
	"unicode"
)

func main() {
	file, err := os.Open("./input.txt")
	if err != nil {
		log.Fatal(err)
	}
	defer file.Close()

	reader := bufio.NewReader(file)
	sum := 0
	firstNumber := 'A'
	lastNumber := 'A'
	for {
		if c, _, err := reader.ReadRune(); err != nil {
			if err == io.EOF {
				break
			} else {
				log.Fatal(err)
			}
		} else {
			if unicode.IsDigit(c) {
				if firstNumber == 'A' {
					firstNumber = c
				} else {
					lastNumber = c
				}
			}
			if c == '\n' {
				if lastNumber == 'A' {
					lastNumber = firstNumber
				}
				toInt, _ := strconv.Atoi(string(firstNumber) + string(lastNumber))
				sum += toInt
				firstNumber = 'A'
				lastNumber = 'A'
			}
		}
	}
	fmt.Print(sum)
}
