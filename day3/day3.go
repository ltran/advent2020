package day3

import (
	"fmt"
	"strings"

	"github.com/ltran/advent2020/adventutil"
)

type Slope struct {
	X, Y int
}

func CountForest(filepath string, slopes []Slope, debug bool) int {
	rows := strings.Split(adventutil.ReadFile(filepath), "\n")

	counters := []int{}
	for _, slope := range slopes {
		counters = append(counters, CountTrees(rows, slope.X, slope.Y, debug))
	}

	sum := 1
	for _, c := range counters {
		sum *= c
	}

	return sum
}

func CountTrees(rows []string, xInc int, yInc int, debug bool) int {
	currX, forest := 0, 0
	tileWidth := len(rows[0])

	for y := yInc; y < len(rows); y += yInc {
		row := rows[y]
		nextX := (currX + xInc) % tileWidth
		if debug {
			fmt.Println(string(row), nextX)
			fmt.Printf("%*s^\n", nextX, " ")
		}

		if row[nextX] == '#' {
			forest++
		}

		currX = nextX
	}

	return forest
}
