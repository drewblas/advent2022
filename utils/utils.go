package utils

import (
	"fmt"
	"regexp"
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
