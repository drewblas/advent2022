package day11

import (
	"advent/utils"
	"fmt"
	"regexp"
	"sort"
	"strconv"

	"github.com/samber/lo"
)

type Monkey struct {
	Name            int
	items           []int
	op              string
	opNum           int
	divisibleBy     int
	trueTarget      int
	falseTarget     int
	inspectionCount int
}

type Toss struct {
	Item int
	To   int
}

func PartA(input string) string {
	monkeys := parseInput(input)
	fmt.Println(monkeys)

	rounds := 20

	for i := 0; i < rounds; i++ {
		fmt.Println("Round: ", i+1)

		for j := 0; j < len(monkeys); j++ {
			tosses := monkeys[j].Turn()
			lo.ForEach(tosses, func(t Toss, i int) {
				monkeys[t.To].Catch(t.Item)
			})
		}

		fmt.Println(monkeys)
	}

	inspections := lo.Map(monkeys, func(m Monkey, i int) int {
		return m.inspectionCount
	})
	sort.Ints(inspections)

	utils.Debug(inspections)

	monkeyBusiness := inspections[len(inspections)-1] * inspections[len(inspections)-2]

	return fmt.Sprintf("Monkey Business: %d", monkeyBusiness)
}

func PartB(input string) string {
	lines := utils.SplitLines(input)

	return fmt.Sprintln("Unimplemented. Lines: ", len(lines))
}

func (m *Monkey) Catch(item int) {
	m.items = append(m.items, item)
}

func (m *Monkey) Turn() []Toss {
	tosses := []Toss{}

	for _, item := range m.items {
		m.inspectionCount = m.inspectionCount + 1

		// Inspect
		if m.op == "+" {
			item = item + m.opNum
		} else if m.op == "*" {
			item = item * m.opNum
		} else if m.op == "^" {
			item = item * item
		}

		//Relief
		item = item / 3

		// Test
		if item%m.divisibleBy == 0 {
			tosses = append(tosses, Toss{item, m.trueTarget})
		} else {
			tosses = append(tosses, Toss{item, m.falseTarget})
		}
	}

	m.items = []int{}

	return tosses
}

func (m Monkey) String() string {
	return fmt.Sprintf("%#v", m)
}

func parseInput(input string) []Monkey {
	lines := utils.SplitLines(input)

	monkeys := make([]Monkey, 0)
	activeMonkey := -1

	reMonkey := regexp.MustCompile(`Monkey (?P<Num>\d+):`)
	reItems := regexp.MustCompile("  Starting items: (?P<ItemList>.*)")
	reOperation := regexp.MustCompile(`  Operation: new = old (?P<Operator>.) (?P<Operand>.+)`)
	reTest := regexp.MustCompile(`  Test: divisible by (?P<Num>\d+)`)
	reTrue := regexp.MustCompile(`    If true: throw to monkey (?P<Num>\d+)`)
	reFalse := regexp.MustCompile(`    If false: throw to monkey (?P<Num>\d+)`)

	for _, line := range lines {
		_, err := utils.MatchLineToMap(line, reMonkey)

		if err == nil {
			activeMonkey = activeMonkey + 1
			monkeys = append(monkeys, Monkey{Name: activeMonkey})
			continue
		}

		match, err := utils.MatchLineToMap(line, reItems)
		if err == nil {
			items := utils.SplitAndConvertToInts(match["ItemList"], ", ")
			monkeys[activeMonkey].items = items
			continue
		}

		match, err = utils.MatchLineToMap(line, reOperation)
		if err == nil {
			if match["Operand"] == "old" {
				monkeys[activeMonkey].op = "^"
				monkeys[activeMonkey].opNum = 2
			} else {
				monkeys[activeMonkey].op = match["Operator"]
				monkeys[activeMonkey].opNum = lo.Must(strconv.Atoi(match["Operand"]))
			}
			continue
		}

		match, err = utils.MatchLineToMap(line, reTest)
		if err == nil {
			num := lo.Must(strconv.Atoi(match["Num"]))
			monkeys[activeMonkey].divisibleBy = num
			continue
		}

		match, err = utils.MatchLineToMap(line, reTrue)
		if err == nil {
			num := lo.Must(strconv.Atoi(match["Num"]))
			monkeys[activeMonkey].trueTarget = num
			continue
		}

		match, err = utils.MatchLineToMap(line, reFalse)
		if err == nil {
			num := lo.Must(strconv.Atoi(match["Num"]))
			monkeys[activeMonkey].falseTarget = num
			continue
		}
	}

	return monkeys
}
