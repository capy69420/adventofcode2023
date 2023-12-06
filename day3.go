package main

import (
	"bufio"
	"fmt"
	"os"
)

func main() {
	// open file
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

	for _, line := range lines {
		fmt.Println(line)
	}

	return
}
