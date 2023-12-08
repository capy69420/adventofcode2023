package main

import (
	"bufio"
	"fmt"
	"os"
)

func isDigit(char byte) bool {
	return char >= '0' && char <= '9'
}

func isSymbol(char byte) bool {
	return !isDigit(char) && char != '.'
}

func main() {
	// Open file
	file, err := os.Open("day3_input.txt")
	if err != nil {
		fmt.Printf("Error opening the file", err)
		return
	}
	defer file.Close()

	var lines []string
	scanner := bufio.NewScanner(file)
	for scanner.Scan() {
		line := scanner.Text()
		lines = append(lines, line)
	}

	if err := scanner.Err(); err != nil {
		fmt.Println("Error scanning:", err)
		return
	}

	numbersWithSymbol := make([]string, 0)
	sum := 0

	for y := 0; y < len(lines); y++ {
		line := lines[y]
		x := 0
		for x < len(line) {
			if isDigit(lines[y][x]) {

				// Find the full number
				num := findNumber(lines, x, y)
				fmt.Println("Current number:", string(num)) // Print line[x] when a digit is found
				if num != "" {
					// foundSymbol := false
					// Check surrounding characters, including diagonals, for symbols
					x += len(num)
				}
			} else {
				x++
			}
		}
	}

	fmt.Println("Numbers with symbols nearby:")
	for _, num := range numbersWithSymbol {
		fmt.Println(num)
	}

	fmt.Println("Sum of numbers with symbols nearby:", sum)
}

func findNumber(lines []string, startX, startY int) string {
	var number string
	for y := startY; y < len(lines); y++ {
		line := lines[y]
		for x := startX; x < len(line); x++ {
			if isDigit(lines[y][x]) {
				number += string(lines[y][x])
			} else if number != "" {
				return number
			}
		}
		startX = 0 // Reset startX for subsequent lines
	}
	return number
}
