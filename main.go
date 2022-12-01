package main

import (
	"advent/day1"
	"advent/utils"
	"os"
)

func main() {
	// Open input file
	input, err := os.ReadFile("day1/1a.txt")
	if err != nil {
		panic(err)
	}

	result := day1.Part1B(string(input))
	utils.Debug(result)
}
