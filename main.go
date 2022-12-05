package main

import (
	"advent/day3"
	"advent/utils"
	"os"
)

func main() {
	// Open input file
	input, err := os.ReadFile("day3/3a.txt")
	if err != nil {
		panic(err)
	}

	result := day3.Part1B(string(input))
	utils.Debug(result)
}
