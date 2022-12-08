package day6

import (
	"advent/utils"
	"fmt"
	"strconv"
	"strings"

	"github.com/samber/lo"
)

func PartA(input string) string {
	dirSizes := pullDirSizes(input)

	utils.Debug(dirSizes)

	dirSizes = lo.PickBy(dirSizes, func(k string, v int) bool {
		return v <= 100000
	})

	sum := lo.Sum(lo.Values(dirSizes))

	return fmt.Sprint(sum)
}

func PartB(input string) string {
	dirSizes := pullDirSizes(input)

	needed := 30000000 - (70000000 - dirSizes["/"])
	fmt.Println("Needed: ", needed)

	dirSizes = lo.PickBy(dirSizes, func(k string, v int) bool {
		return v >= needed
	})

	pick := lo.Min(lo.Values(dirSizes))

	return fmt.Sprint(pick)
}

func pullDirSizes(input string) map[string]int {
	lines := strings.Split(input, "\n")
	lines = lines[:len(lines)-1] // Remove blank last line

	pwd := []string{""}
	dirSizes := map[string]int{}

	for _, line := range lines {
		if line == "$ cd /" {
			pwd = []string{""}
		} else if line == "$ cd .." {
			pwd = pwd[:len(pwd)-1]
		} else if line[0:4] == "$ cd" {
			pwd = append(pwd, line[5:])
		} else if line == "$ ls" {
			// Nothing to do here
		} else if line[0:3] == "dir" {
			// Nothing to do here
		} else {
			sizeStr := line[0:strings.Index(line, " ")]
			fileSize := lo.Must(strconv.Atoi(sizeStr))
			// Slurp a file inside this dir
			dir := ""
			for _, chunk := range pwd {
				dir = dir + "/" + chunk
				// Add the file to the total of every dir in our path
				dirSizes[dir] += fileSize
			}
		}
	}
	return dirSizes
}
