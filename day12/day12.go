package day12

import (
	"advent/utils"
	"fmt"
	"strings"

	"github.com/samber/lo"
)

type Pos struct {
	x int
	y int
}

type Map struct {
	heights [][]string
	route   [][]string
	cascade []Pos
	start   Pos
	end     Pos
	cnt     int
}

func PartA(input string) string {
	m := Map{}
	m.Load(input)

	m.PrintPos()
	m.PrintHeights()
	m.PrintRoute()

	success := false

	// Do while success is false
	for !success {
		success = m.Cascade()
		fmt.Println("Cascade: ", m.cnt)
		m.PrintRoute()
	}

	return fmt.Sprintln("Unimplemented. Lines: ", 0)
}

func PartB(input string) string {
	lines := utils.SplitLines(input)

	return fmt.Sprintln("Unimplemented. Lines: ", len(lines))
}

func (m *Map) Load(input string) {
	lines := utils.SplitLines(input)

	m.heights = [][]string{}
	m.route = [][]string{}
	x := -1

	for y, line := range lines {
		lineHeights := strings.Split(line, "")

		x = lo.IndexOf(lineHeights, "S")
		if x >= 0 {
			m.start = Pos{x, y}
		}

		x = lo.IndexOf(lineHeights, "E")
		if x >= 0 {
			m.end = Pos{x, y}
			m.cascade = append(m.cascade, Pos{x, y})
		}

		m.heights = append(m.heights, lineHeights)
		lineRoute := strings.Split(strings.Repeat(".", len(line)), "")
		m.route = append(m.route, lineRoute)
	}

}

func (m *Map) Cascade() bool {
	newCascade := []Pos{}
	m.cnt = m.cnt + 1

	for _, pos := range m.cascade {
		adjacents := m.Adjacent(pos)
		oneStepLower := m.heights[pos.y][pos.x][0] - 1

		for _, adjacent := range adjacents {
			if adjacent == m.start {
				fmt.Println("Found start at", adjacent, "in", m.cnt, "steps")
				return true
			}

			if m.heights[adjacent.y][adjacent.x][0] == oneStepLower &&
				m.route[adjacent.y][adjacent.x] == "." {
				m.route[adjacent.y][adjacent.x] = "X"
				newCascade = append(newCascade, adjacent)
			}
		}
	}

	m.cascade = newCascade
	return false
}

func (m *Map) Adjacent(pos Pos) []Pos {
	width := len(m.heights)
	height := len(m.heights[0])

	adjacents := []Pos{}

	if pos.x > 0 {
		adjacents = append(adjacents, Pos{pos.x - 1, pos.y})
	}

	if pos.x < width-1 {
		adjacents = append(adjacents, Pos{pos.x + 1, pos.y})
	}

	if pos.y > 0 {
		adjacents = append(adjacents, Pos{pos.x, pos.y - 1})
	}

	if pos.y < height-1 {
		adjacents = append(adjacents, Pos{pos.x, pos.y + 1})
	}

	return adjacents
}

func (m *Map) PrintPos() {
	fmt.Println("Start:", m.start)
	fmt.Println("End:", m.end)
}

func (m *Map) PrintHeights() {
	fmt.Println("Heights:")
	for _, line := range m.heights {
		fmt.Println(line)
	}
}

func (m *Map) PrintRoute() {
	fmt.Println("Route:")
	for _, line := range m.route {
		fmt.Println(line)
	}
}
