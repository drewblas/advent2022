package main

import (
	"advent/day4"
	"advent/utils"
	"os"
)

func main() {
	// Open input file
	input, err := os.ReadFile("day4/4a.txt")
	if err != nil {
		panic(err)
	}

	result := day4.Part1B(string(input))
	utils.Debug(result)
}
