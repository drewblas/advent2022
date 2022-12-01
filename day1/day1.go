package day1

import (
	"advent/utils"
	"sort"
	"strconv"
	"strings"

	"github.com/samber/lo"
)

func Part1A(input string) int {
	totals := totalPerElf(input)

	_, max := utils.MinMax(totals)

	return max
}

func Part1B(input string) int {
	totals := totalPerElf(input)
	sort.Ints(totals)

	topThree := lo.Sum(totals[len(totals)-3:])
	return topThree
}

func totalPerElf(input string) []int {
	perElf := utils.SplitEmptyLines(input)

	totalPerElf := make([]int, len(perElf))

	for i, elf := range perElf {
		cals := strings.Split(elf, "\n")
		for _, cal := range cals {
			if cal == "" {
				continue
			}
			val, err := strconv.Atoi(cal)
			if err != nil {
				panic(err)
			}

			totalPerElf[i] += val
		}
	}

	return totalPerElf
}
