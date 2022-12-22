package main

import (
	. "advent/day10" // Pick a day
	"advent/utils"
	"fmt"
	"os"
)

func main() {
	day := "day10"

	run("Test Part A:", day+"/test.txt", PartA)
	run("Full Part A:", day+"/input.txt", PartA)
	run("Test Part B:", day+"/test.txt", PartB)
	run("Full Part B:", day+"/input.txt", PartB)
}

func run(title, filename string, f func(string) string) {
	input, err := os.ReadFile(filename)
	if err == nil {
		fmt.Println(title)
		result := f(string(input))
		utils.Debug(result)
	}
}
