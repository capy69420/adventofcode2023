package main

import (
	"bufio"
	"fmt"
	"os"
	"strconv"
	"strings"
)

type ColorSet map[string]int

func (cs ColorSet) Add(color string, value int) {
	cs[color] = value
}

func main() {
	// open file
	file, err := os.Open("day2_input.txt")
	if err != nil {
		fmt.Printf("Error opening day2_input.txt", err)
		return
	}
	defer file.Close()

	// read input file line by line
	scanner := bufio.NewScanner(file)

	baseColors := map[string]int{"red": 12, "green": 13, "blue": 14}
	idsSum := 0
	colorsSum := 0
	// parse each line
	for scanner.Scan() {
		line := scanner.Text()
		// part 2 minimum required colors
		minColors := map[string]int{"red": 0, "green": 0, "blue": 0}
		// parse line
		parts := strings.Split(line, ":")
		gameId, err := strconv.Atoi(strings.Split(parts[0], " ")[1])
		if err != nil {
			fmt.Println("Error parsing game id: ", err)
			return
		}
		sets := strings.Split(parts[1], ";")

		validGame := true
		product := 1 // part 2
		for _, s := range sets {
			// after the set is done the colors are put back into the bag
			colorSet := make(ColorSet)
			colors := strings.Split(s, ",")
			for _, item := range colors {
				pair := strings.Split(item, " ")
				color := pair[2]
				val := pair[1]
				// add the color value to the set
				if intValue, err := strconv.Atoi(val); err == nil {
					if existingValue, ok := colorSet[color]; ok {
						colorSet[color] = existingValue + intValue
					} else {
						colorSet.Add(color, intValue)
					}
				} else {
					fmt.Println("Invalid value for color: ", val)
				}
			}
			for color, val := range colorSet {
				// part 2 get the max color value
				if minColors[color] < val {
					minColors[color] = val
				}
				if baseColors[color] < val {
					// not valid id
					validGame = false
				}
			}
		}

		for _, color := range minColors {
			product = product * color
		}
		if validGame {
			idsSum += gameId
		}
		colorsSum += product
	}

	if err := scanner.Err(); err != nil {
		fmt.Println("Error during scanning", err)
	}

	fmt.Println("Part 1, id sum:", idsSum)
	fmt.Println("Part 2, sum of colors product:", colorsSum)
}
