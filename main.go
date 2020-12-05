package main

import (
	"fmt"
	"strconv"
	"strings"

	"github.com/ltran/advent2020/day1"
	"github.com/ltran/advent2020/day2"
	"github.com/ltran/advent2020/day3"
)

func main() {
	fmt.Println("day1-1:", day1.SumTwo("data/input1.txt", 2020))
	fmt.Println("day1-2:", day1.SumThree("data/input1.txt", 2020))

	fmt.Println("day2-1:", day2.CountValid("data/input2.txt", false))
	fmt.Println("day2-1:", day2.CountValid("data/input2.txt", true))

	fmt.Println("day3-1", day3.CountTrees("data/input3.txt", false))
}

func intstringToIntSlice(str string) []int {
	strEntries := strings.Split(str, "\n")
	entries := make([]int, len(strEntries))
	for i, e := range strEntries {
		entries[i], _ = strconv.Atoi(e)
	}
	return entries
}

func stringToStringSlice(str string) []string {
	return strings.Split(str, "\n")
}
