package main

import (
	"fmt"
	"strconv"
	"strings"

	"github.com/ltran/advent2020/day1"
	"github.com/ltran/advent2020/day2"
)

func main() {
	fmt.Println("day1-1:", day1.SumTwo(intstringToIntSlice(day1.INPUT), 2020))
	fmt.Println("day1-2:", day1.SumThree(intstringToIntSlice(day1.INPUT), 2020))

	fmt.Println("day2-1:", day2.CountValid(stringToStringSlice(day2.INPUT), false))
	fmt.Println("day2-1:", day2.CountValid(stringToStringSlice(day2.INPUT), true))
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
