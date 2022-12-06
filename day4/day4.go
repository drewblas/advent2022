package day4

import (
	"regexp"
	"strconv"
	"strings"

	"github.com/samber/lo"
)

func Part1A(input string) int {
	total := 0

	r := regexp.MustCompile(`(?P<OneStart>\d+)-(?P<OneEnd>\d+),(?P<TwoStart>\d+)-(?P<TwoEnd>\d+)`)
	lines := strings.Split(input, "\n")
	lines = lines[:len(lines)-1] // Remove blank last line

	for _, line := range lines {
		matches := r.FindStringSubmatch(line)
		if len(matches) == 0 {
			continue
		}

		oneStart, _ := strconv.Atoi(matches[1])
		oneEnd, _ := strconv.Atoi(matches[2])
		twoStart, _ := strconv.Atoi(matches[3])
		twoEnd, _ := strconv.Atoi(matches[4])

		if fullyContains(oneStart, oneEnd, twoStart, twoEnd) {
			total += 1
		}
	}

	return total
}

func fullyContains(oneStart, oneEnd, twoStart, twoEnd int) bool {
	return (oneStart <= twoStart && oneEnd >= twoEnd) || (twoStart <= oneStart && twoEnd >= oneEnd)
}

// return 0 if no overlap, return 1 if overlap
func overlaps(ints []int) int {
	if (ints[2] > ints[1] && ints[0] < ints[2]) || (ints[0] > ints[3] && ints[2] < ints[0]) {
		return 0
	} else {
		return 1
	}
}

func Part1B(input string) int {
	r := regexp.MustCompile(`(?P<OneStart>\d+)-(?P<OneEnd>\d+),(?P<TwoStart>\d+)-(?P<TwoEnd>\d+)`)

	lines := strings.Split(input, "\n")
	lines = lines[:len(lines)-1] // Remove blank last line

	// Match "1-2,3-4"
	matches := lo.Map(lines, func(line string, _ int) []int {
		match := r.FindStringSubmatch(line)

		// Convert matches to []int{1,2,3,4}
		ints := lo.Map(match[1:], func(s string, _ int) int {
			i, _ := strconv.Atoi(s)
			return i
		})

		return ints
	})

	total := lo.SumBy(matches, overlaps)

	return total
}
