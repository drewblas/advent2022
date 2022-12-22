package day9

import (
	"fmt"
	"strconv"
	"strings"

	"github.com/samber/lo"
)

type Pos struct {
	x int
	y int
}

func PartA(input string) string {
	lines := strings.Split(input, "\n")
	lines = lines[:len(lines)-1] // Remove blank last line

	h, t := Pos{0, 0}, Pos{0, 0}
	visited := make(map[Pos]bool)

	for cnt, line := range lines {
		dir, dist := line[0], lo.Must(strconv.Atoi(line[2:]))

		for i := 0; i < dist; i++ {
			h = move(h, dir)
			t = follow(h, t)

			visited[t] = true

			if false {
				fmt.Println("Line", cnt, ": ", line, " - T", t, " : H", h)
				printMatrix(visited, h, t)
			}
		}

	}

	return fmt.Sprint(len(visited))
}

func move(h Pos, dir byte) Pos {
	switch dir {
	case 'U':
		h.y += 1
	case 'D':
		h.y -= 1
	case 'R':
		h.x += 1
	case 'L':
		h.x -= 1
	}

	return h
}

func follow(h, t Pos) Pos {
	// Tail follows head
	if h.x-t.x > 1 {
		t.x += 1
		if h.y > t.y {
			t.y += 1
		} else if h.y < t.y {
			t.y -= 1
		}
	} else if h.x-t.x < -1 {
		t.x -= 1
		if h.y > t.y {
			t.y += 1
		} else if h.y < t.y {
			t.y -= 1
		}
	} else if h.y-t.y > 1 {
		t.y += 1
		if h.x > t.x {
			t.x += 1
		} else if h.x < t.x {
			t.x -= 1
		}
	} else if h.y-t.y < -1 {
		t.y -= 1
		if h.x > t.x {
			t.x += 1
		} else if h.x < t.x {
			t.x -= 1
		}
	}

	return t
}

func printMatrix(visited map[Pos]bool, h, t Pos) {
	minX, maxX, minY, maxY := 0, 0, 0, 0
	for pos := range visited {
		if pos.x < minX {
			minX = pos.x
		}
		if pos.x > maxX {
			maxX = pos.x
		}
		if pos.y < minY {
			minY = pos.y
		}
		if pos.y > maxY {
			maxY = pos.y
		}
	}

	for y := maxY; y >= minY; y-- {
		for x := minX; x <= maxX; x++ {
			p := Pos{x, y}
			if p == h {
				fmt.Print("H")
			} else if p == t {
				fmt.Print("T")
			} else if visited[Pos{x, y}] {
				fmt.Print("#")
			} else {
				fmt.Print(".")
			}
		}
		fmt.Println()
	}
}

func PartB(input string) string {
	lines := strings.Split(input, "\n")
	lines = lines[:len(lines)-1] // Remove blank last line

	snake := lo.RepeatBy(10, func(i int) Pos { return Pos{0, 0} })
	visited := make(map[Pos]bool)

	for _, line := range lines {
		dir, dist := line[0], lo.Must(strconv.Atoi(line[2:]))

		for i := 0; i < dist; i++ {
			snake[0] = move(snake[0], dir)
			for i := 1; i < len(snake); i++ {
				snake[i] = follow(snake[i-1], snake[i])
			}

			visited[snake[9]] = true
		}

		// printMatrix(visited, snake[0], snake[9])
	}

	return fmt.Sprint(len(visited))
}
