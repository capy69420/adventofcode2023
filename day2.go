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
	// parse each line
	for scanner.Scan() {
		line := scanner.Text()
		// parse line
		parts := strings.Split(line, ":")
		gameId, err := strconv.Atoi(strings.Split(parts[0], " ")[1])
		if err != nil {
			fmt.Println("Error parsing game id: ", err)
			return
		}
		fmt.Println("gameId: ", gameId)
		sets := strings.Split(parts[1], ";")

		validGame := true
		for _, s := range sets {
			// after the set is done the colors are put back into the bag
			colorSet := make(ColorSet)
			//fmt.Println("sets: ", s)
			colors := strings.Split(s, ",")
			for _, item := range colors {
				pair := strings.Split(item, " ")
				color := pair[2]
				val := pair[1]
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
			fmt.Println("colorSet:", colorSet)
			for color, val := range colorSet {
				if baseColors[color] < val {
					// not valid id
					validGame = false
				}
			}
		}
		if validGame {
			idsSum += gameId
		}
	}

	if err := scanner.Err(); err != nil {
		fmt.Println("Error during scanning", err)
	}

	fmt.Println("id sum:", idsSum)
}
