package main

import (
	"bufio"
	"fmt"
	"os"
	"strconv"
)

// day 1
func main() {
	// open file
	file, err := os.Open("day1_input.txt")
	if err != nil {
		fmt.Printf("Error opening:", err)
		return
	}
	defer file.Close()

	// read the file line by line
	scanner := bufio.NewScanner(file)

	sumCounter := 0
	// read each line
	for scanner.Scan() {
		line := scanner.Text()
		// parse the line
		var ldigit rune
		var rdigit rune
		digits := []string{"one", "two", "three", "four", "five", "six", "seven", "eight", "nine"}
		// left digit
		// check for a digit or a spelled digit whatever coms first
		for i := 0; i < len(line); i++ {
			digitFound := false
			if line[i] >= '0' && line[i] <= '9' {
				ldigit = rune(line[i])
				break
			} else {
				// check if it is spelled out with letters, iterate over possible digits
				for j := 0; j < len(digits); j++ {
					if i+len(digits[j]) <= len(line) && line[i:i+len(digits[j])] == digits[j] {
						// found a spelled digit, turn it into a number
						// add '0' to turn it into a string
						ldigit = rune(j + 1 + '0')
						digitFound = true
						break
					}
					if digitFound {
						break
					}
				}
				if digitFound {
					break
				}
			}
		}
		// right digit
		for i := len(line) - 1; i >= 0; i-- {
			digitFound := false
			if line[i] >= '0' && line[i] <= '9' {
				rdigit = rune(line[i])
				break
			} else {
				for j := 0; j < len(digits); j++ {
					if i+len(digits[j]) <= len(line) && line[i:i+len(digits[j])] == digits[j] {
						rdigit = rune(j + 1 + '0')
						digitFound = true
						break
					}
				}
				if digitFound {
					break
				}
			}
			if digitFound {
				break
			}
		}
		num, err := strconv.Atoi(string(ldigit) + string(rdigit))

		if err != nil {
			fmt.Println("Error no digits: ", err)
		}
		fmt.Println(num)
		sumCounter += int(num)
	}
	fmt.Println(sumCounter)

	// check for scanner errors
	if err := scanner.Err(); err != nil {
		fmt.Printf("Error parsing:", err)
		return
	}
	// find the first digit
	// find the last digit
	// per line calibration value = first_digit last_digit
	// sum all values
}
