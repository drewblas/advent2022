package day8

import (
	"advent/utils"
	"fmt"
	"strconv"
	"strings"

	"github.com/samber/lo"
)

func PartA(input string) string {
	trees := loadTrees(input)

	visible := countVisibleTrees(trees)

	utils.Debug(visible)

	return fmt.Sprint(visible)
}

func PartB(input string) string {
	trees := loadTrees(input)

	// max := scenicScore(trees, 2, 1)

	max := maxScenicScore(trees)

	return fmt.Sprint(max)
}

func loadTrees(input string) [][]int {
	lines := strings.Split(input, "\n")
	lines = lines[:len(lines)-1] // Remove blank last line

	trees := [][]int{}

	for _, line := range lines {
		heights := lo.Map(strings.Split(line, ""), func(s string, i int) int {
			return lo.Must(strconv.Atoi(s))
		})
		trees = append(trees, heights)
	}

	return trees
}

func countVisibleTrees(trees [][]int) int {
	visible := 0

	for i, row := range trees {
		for j, _ := range row {
			if isVisible(trees, i, j) {
				visible += 1
			}
		}
	}

	return visible
}

func maxScenicScore(trees [][]int) int {
	scores := [][]int{}

	for i, row := range trees {
		scores = append(scores, []int{})
		for j, _ := range row {
			scores[i] = append(scores[i], scenicScore(trees, i, j))
		}
	}

	return lo.Max(lo.Flatten(scores))
}

// This is the brute force I'm most embarassed by so far...
func isVisible(trees [][]int, i, j int) bool {
	tree := trees[i][j]

	if i == 0 || j == 0 || i == len(trees)-1 || j == len(trees[i])-1 {
		return true
	}

	row := trees[i]

	if lo.Max(row[0:j]) < tree {
		return true
	}

	if lo.Max(row[j+1:]) < tree {
		return true
	}

	col := getColumn(trees, j)

	if lo.Max(col[0:i]) < tree {
		return true
	}

	if lo.Max(col[i+1:]) < tree {
		return true
	}

	return false
}

func getColumn(trees [][]int, j int) []int {
	return lo.Map(trees, func(row []int, i int) int {
		return row[j]
	})
}

func duplicateIntSlice(slice []int) []int {
	duplicate := make([]int, len(slice))
	copy(duplicate, slice)
	return duplicate
}

func scenicScore(trees [][]int, i, j int) int {
	if i == 0 || j == 0 || i == len(trees)-1 || j == len(trees[i])-1 {
		return 0
	}

	tree := trees[i][j]
	row := trees[i]
	col := getColumn(trees, j)
	rays := [][]int{
		lo.Reverse(duplicateIntSlice(row[0:j])),
		row[j+1:],
		lo.Reverse(duplicateIntSlice(col[0:i])),
		col[i+1:],
	}

	scores := lo.Map(rays, func(ray []int, i int) int {
		for k, h := range ray {
			if h >= tree {
				return k + 1
			}
		}

		return len(ray)
	})

	if len(trees) < 20 {
		fmt.Println(tree, "-", i, ",", j)
		utils.Debug(scores)
	}

	return lo.Reduce(scores, func(agg, item, i int) int {
		return agg * item
	}, 1)
}
