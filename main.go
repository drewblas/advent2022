package main

import (
	"advent/day2"
	"advent/utils"
	"os"
)

func main() {
	// Open input file
	input, err := os.ReadFile("day2/2a.txt")
	if err != nil {
		panic(err)
	}

	result := day2.Part1B(string(input))
	utils.Debug(result)
}
