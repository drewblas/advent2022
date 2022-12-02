package day2

import "strings"

var partAResults = map[string]int{
	"A X": 4, // Rock Rock = 3+1
	"A Y": 8, // Rock Paper = 6+2
	"A Z": 3, // Rock Sci = 0+3
	"B X": 1, // Paper Rock = 0+1
	"B Y": 5, // Paper Paper = 3+2
	"B Z": 9, // Paper Sci = 6+3
	"C X": 7, // Sci Rock = 6+1
	"C Y": 2, // Sci Paper = 0+2
	"C Z": 6, // Sci Sci = 3+3
}

var partBResults = map[string]int{
	"A X": 3, // Rock Lose Sci = 0+3
	"A Y": 4, // Rock Draw Rock = 3+1
	"A Z": 8, // Rock Win Paper = 6+2
	"B X": 1, // Paper Lose Rock = 0+1
	"B Y": 5, // Paper Draw Paper = 3+2
	"B Z": 9, // Paper Win Sci = 6+3
	"C X": 2, // Sci Lose Paper = 0+2
	"C Y": 6, // Sci Draw Sci = 3+3
	"C Z": 7, // Sci Win Rock = 6+1
}

func Part1A(input string) int {
	total := 0

	for _, line := range strings.Split(input, "\n") {
		result := partAResults[line]

		// Add the result to the total
		total += result
	}

	return total
}

func Part1B(input string) int {
	total := 0

	for _, line := range strings.Split(input, "\n") {
		result := partBResults[line]

		// Add the result to the total
		total += result
	}

	return total
}
