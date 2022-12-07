package day6

import (
	"fmt"
	"strings"

	"github.com/samber/lo"
)

func PartA(input string) string {
	lines := strings.Split(input, "\n")
	lines = lines[:len(lines)-1] // Remove blank last line

	for _, line := range lines {
		for i := 4; i < len(line); i++ {
			window := strings.Split(line[i-4:i], "")
			dups := lo.FindDuplicates(window)
			if len(dups) == 0 {
				fmt.Println(i)
				break
			}
		}
	}

	return "Done"
}

func PartB(input string) string {
	lines := strings.Split(input, "\n")
	lines = lines[:len(lines)-1] // Remove blank last line

	winSize := 14

	for _, line := range lines {
		for i := winSize; i < len(line); i++ {
			window := strings.Split(line[i-winSize:i], "")
			dups := lo.FindDuplicates(window)
			if len(dups) == 0 {
				fmt.Println(i)
				break
			}
		}
	}

	return "Done"
}
