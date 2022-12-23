package main

import (
	// Pick a day
	anyday "advent/day11"
	"advent/utils"
	"fmt"
	"os"
)

func main() {
	day := "day11"

	run("Test Part A:", day+"/test.txt", anyday.PartA)
	run("Full Part A:", day+"/input.txt", anyday.PartA)
	run("Test Part B:", day+"/test.txt", anyday.PartB)
	run("Full Part B:", day+"/input.txt", anyday.PartB)
}

func run(title, filename string, f func(string) string) {
	input, err := os.ReadFile(filename)
	if err == nil {
		fmt.Println(title)
		result := f(string(input))
		utils.Debug(result)
	}
}
