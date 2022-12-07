package utils

import (
	"fmt"
	"regexp"
	"strconv"
	"strings"

	"github.com/samber/lo"
)

func Debug(v interface{}) {
	fmt.Printf("%#v\n", v)
	// fmt.Printf("%+v\n", v)
}

func SplitEmptyLines(input string) []string {
	re := regexp.MustCompile("\n\n")
	return re.Split(input, -1)
}

func MinMax(array []int) (int, int) {
	var max int = array[0]
	var min int = array[0]
	for _, value := range array {
		if max < value {
			max = value
		}
		if min > value {
			min = value
		}
	}
	return min, max
}

func MinMaxIndex(array []int) (int, int) {
	var max int = array[0]
	var maxI int = 0
	var min int = array[0]
	var minI int = 0
	for i, value := range array {
		if max < value {
			max = value
			maxI = i
		}
		if min > value {
			min = value
			minI = i
		}
	}
	return minI, maxI
}

// MatchLinesToMap takes an input of lines
// and a regex with named capture groups
// and returns a map[string]int for each line
func MatchLinesToMap(input string, re *regexp.Regexp) []map[string]int {
	lines := strings.Split(input, "\n")
	// Remove any blank lines
	lines = lo.Reject(lines, func(line string, _ int) bool {
		return strings.TrimSpace(line) == ""
	})

	matches := lo.Map(lines, func(line string, _ int) map[string]int {
		result := make(map[string]int)
		match := re.FindStringSubmatch(line)

		for i, name := range re.SubexpNames()[1:] {
			result[name] = lo.Must(strconv.Atoi(match[i+1]))
		}

		return result
	})

	return matches
}
