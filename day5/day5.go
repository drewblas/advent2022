package day5

import (
	"advent/utils"
	"regexp"
	"strings"

	"github.com/samber/lo"
)

func PartA(input string) string {
	stackStr, instStr, found := strings.Cut(input, "\n\n")
	if !found {
		panic("No blank line found")
	}

	stacks := loadStacks(stackStr)
	instructions := loadInstructions(instStr)
	stacks = applyInstructions9000(stacks, instructions)

	result := lo.Map(stacks[1:], func(stack []string, _ int) string {
		return stack[len(stack)-1]
	})

	return strings.Join(result, "")
}

func PartB(input string) string {
	stackStr, instStr, found := strings.Cut(input, "\n\n")
	if !found {
		panic("No blank line found")
	}

	stacks := loadStacks(stackStr)
	instructions := loadInstructions(instStr)
	stacks = applyInstructions9001(stacks, instructions)

	result := lo.Map(stacks[1:], func(stack []string, _ int) string {
		return stack[len(stack)-1]
	})

	return strings.Join(result, "")
}

func loadStacks(input string) [][]string {
	lines := strings.Split(input, "\n")
	lines = lines[:len(lines)-1] // Remove blank last line
	lines = lo.Reverse(lines)

	boxes := lo.Map(lines, func(line string, _ int) []string {
		boxStr := lo.ChunkString(line, 4)
		return lo.Map(boxStr, func(box string, _ int) string {
			return string(box[1])
		})
	})

	// The box map is now a 2D array of strings
	// [][]string{[]string{"Z", "M", "P"}, []string{"N", "C"}, []string{" ", "D"}}
	// Now we rotate 90 degrees so each stack is a row
	stacks := make([][]string, 0)
	stacks = append(stacks, []string{}) // Empty zero row to make the indexes match

	for _, row := range boxes {
		for coli, box := range row {
			if box == " " {
				continue
			}
			i := coli + 1 // All stacks are 1-based
			if len(stacks) <= i {
				stacks = append(stacks, []string{})
			}
			stacks[i] = append(stacks[i], box)
		}
	}

	return stacks
}

func loadInstructions(input string) []map[string]int {
	re := regexp.MustCompile(`move (?P<times>\d+) from (?P<from>\d+) to (?P<to>\d+)`)
	return utils.MatchLinesToMap(input, re)
}

func applyInstructions9000(stacks [][]string, instructions []map[string]int) [][]string {
	for _, inst := range instructions {
		times := inst["times"]
		from := inst["from"]
		to := inst["to"]
		for i := 0; i < times; i++ {
			stacks[to] = append(stacks[to], stacks[from][len(stacks[from])-1])
			stacks[from] = stacks[from][:len(stacks[from])-1]
		}
	}
	return stacks
}

func applyInstructions9001(stacks [][]string, instructions []map[string]int) [][]string {
	for _, inst := range instructions {
		times := inst["times"]
		from := inst["from"]
		to := inst["to"]

		rangeStart := len(stacks[from]) - times
		rangeEnd := len(stacks[from])

		goingToMove := stacks[from][rangeStart:rangeEnd]

		stacks[to] = append(stacks[to], goingToMove...)
		stacks[from] = stacks[from][:rangeStart]
	}
	return stacks
}
