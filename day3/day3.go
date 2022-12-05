package day3

import (
	"advent/utils"
	"strings"

	"github.com/samber/lo"
)

func Part1A(input string) int {
	total := 0

	for _, line := range strings.Split(input, "\n") {
		if line == "" {
			continue
		}

		mid := len(line) / 2
		one, two := line[:mid], line[mid:]
		s := intersect(one, two)

		total += priorityOf(s)
	}

	return total
}

func intersect(a, b string) (diff rune) {
	m := make(map[rune]bool)

	for _, item := range b {
		m[item] = true
	}

	for _, item := range a {
		if _, ok := m[item]; ok {
			return item
		}
	}
	return '!'
}

func intersect3(a, b, c string) (diff rune) {
	m := make(map[rune]bool)
	n := make(map[rune]bool)

	for _, item := range a {
		m[item] = true
	}

	for _, item := range b {
		if _, ok := m[item]; ok {
			n[item] = true
		}
	}

	for _, item := range c {
		if _, ok := n[item]; ok {
			return item
		}
	}

	return '!'
}

func priorityOf(s rune) int {
	priority := int(s - 'A')

	if priority < 27 {
		priority += 27
	} else {
		priority -= 31
	}

	return priority
}

func Part1B(input string) int {
	total := 0

	lines := strings.Split(input, "\n")
	lines = lines[:len(lines)-1] // Remove blank last line

	groups := lo.Chunk(lines, 3)

	for _, group := range groups {
		s := intersect3(group[0], group[1], group[2])
		priority := priorityOf(s)

		utils.Debug(string(s))
		utils.Debug(priority)

		total += priority
	}

	return total
}
