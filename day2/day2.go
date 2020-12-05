package day2

import (
	"strconv"
	"strings"

	"github.com/ltran/advent2020/adventutil"
)

type Row string

func CountValid(filepath string, newValidation bool) int {
	list := strings.Split(adventutil.ReadFile(filepath), "\n")

	counter := 0
	for _, item := range list {
		row := Row(item)
		if newValidation {
			if row.IsValid2() {
				counter++
			}
		} else {
			if row.IsOldValid() {
				counter++
			}
		}
	}

	return counter
}

func (r Row) Min() int {
	i, _ := strconv.Atoi(strings.Split(string(r), "-")[0])
	return i
}

func (r Row) Max() int {
	i, _ := strconv.Atoi(strings.Split(strings.Split(string(r), "-")[1], " ")[0])
	return i
}

func (r Row) C() rune {
	return []rune(strings.Split(strings.Split(string(r), " ")[1], ":")[0])[0]
}

func (r Row) Pwd() string {
	sl := strings.Split(string(r), " ")
	return sl[len(sl)-1]
}

func (r Row) IsValid2() bool {
	if r.Min() >= r.Max() {
		return false
	}

	pwd := r.Pwd()
	minC := pwd[r.Min()-1]
	maxC := pwd[r.Max()-1]

	if minC == maxC || (rune(minC) != r.C() && rune(maxC) != r.C()) {
		return false
	}

	return true
}

func (r Row) IsOldValid() bool {
	password, c, min, max := r.Pwd(), r.C(), r.Min(), r.Max()
	memo := [26]int{}
	letter := c - 97
	if min > max || letter < 0 || letter > 26 {
		return false
	}

	letters := []rune(password)
	for _, letter := range letters {
		memo[int(letter)-97] = memo[int(letter)-97] + 1
	}

	return min <= memo[letter] && memo[letter] <= max
}
