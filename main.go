package main

import (
	"fmt"

	"github.com/ltran/advent2020/day1"
	"github.com/ltran/advent2020/day2"
	"github.com/ltran/advent2020/day3"
	"github.com/ltran/advent2020/day4"
)

func main() {
	fmt.Println("day1-1:", day1.SumTwo("data/input1.txt", 2020))
	fmt.Println("day1-2:", day1.SumThree("data/input1.txt", 2020))

	fmt.Println("day2-1:", day2.CountValid("data/sample2.txt", false))
	fmt.Println("day2-1:", day2.CountValid("data/input2.txt", false))

	fmt.Println("day3-1", day3.CountForest("data/sample3.txt", []day3.Slope{{X: 3, Y: 1}}, false))
	fmt.Println("day3-2", day3.CountForest("data/input3.txt", []day3.Slope{
		{X: 1, Y: 1},
		{X: 3, Y: 1},
		{X: 5, Y: 1},
		{X: 7, Y: 1},
		{X: 1, Y: 2},
	}, false))

	fmt.Println("day4-1", day4.ValidNumPassport("data/sample4.txt", true))
}
