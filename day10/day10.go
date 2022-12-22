package day10

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

type Machine struct {
	cycle          int
	x              int
	signalStrength int
	crt            map[Pos]bool
}

func PartA(input string) string {
	lines := strings.Split(input, "\n")
	lines = lines[:len(lines)-1] // Remove blank last line

	m := Machine{0, 1, 0, make(map[Pos]bool)}

	for _, line := range lines {
		parts := strings.Split(line, " ")
		instruction := parts[0]
		var arg int
		if len(parts) > 1 {
			arg = lo.Must(strconv.Atoi(parts[1]))
		}

		m.Execute(instruction, arg)
	}

	return fmt.Sprint(m.signalStrength)
}

func PartB(input string) string {
	lines := strings.Split(input, "\n")
	lines = lines[:len(lines)-1] // Remove blank last line

	m := Machine{1, 1, 0, make(map[Pos]bool)}

	for _, line := range lines {
		parts := strings.Split(line, " ")
		instruction := parts[0]
		var arg int
		if len(parts) > 1 {
			arg = lo.Must(strconv.Atoi(parts[1]))
		}

		m.Execute(instruction, arg)
	}

	m.PrintCRT()

	return ""
}

func (m *Machine) PrintCRT() {
	for y := 0; y < 6; y++ {
		for x := 0; x < 40; x++ {
			if m.crt[Pos{x, y}] {
				fmt.Print("#")
			} else {
				fmt.Print(".")
			}
		}
		fmt.Println()
	}
}

func (m *Machine) step() {
	beamX := (m.cycle - 1) % 40
	beamY := (m.cycle - 1) / 40
	if beamX >= m.x-1 && beamX <= m.x+1 {
		m.crt[Pos{beamX, beamY}] = true
	}

	m.cycle++

	// If cycle == 20 or % 40, then add signalStrength
	if (m.cycle-20)%40 == 0 {
		m.signalStrength += m.x * m.cycle
		// fmt.Println("Cycle", m.cycle, "x", m.x, "signalStrength", m.signalStrength)
	}
}

func (m *Machine) Execute(instruction string, arg int) {
	switch instruction {
	case "noop":
		m.step()
	case "addx":
		m.step()
		m.step()
		m.x += arg
	}
}
