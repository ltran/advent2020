package day3

import (
	"fmt"
	"strings"

	"github.com/ltran/advent2020/adventutil"
)

func CountTrees(filepath string, debug bool) int {
	debug = false
	rows := strings.Split(adventutil.ReadFile(filepath), "\n")
	currX, forest := 0, 0
	tileWidth := len(rows[0])
	for _, row := range rows[1:] {
		nextX := (currX + 3) % tileWidth
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
