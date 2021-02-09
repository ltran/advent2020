package day4

import (
	"strconv"
	"strings"

	"github.com/ltran/advent2020/adventutil"
)

type Passport string

func validateIntWithinRange(str string, b int, e int) bool {
	num, err := strconv.Atoi(str)
	if err != nil {
		return false
	}

	return b <= num && num <= e
}
func validateHeight(str string) bool {
	pos := -1
	pos = strings.Index(str, "cm")
	if pos > -1 {
		return validateHeightCm(str[0:pos])
	}

	pos = strings.Index(str, "in")
	if pos > -1 {
		return validateHeightIn(str[0:pos])
	}

	return false
}

func validateHeightCm(str string) bool {
	return validateIntWithinRange(str, 150, 193)
}

func validateHeightIn(str string) bool {
	return validateIntWithinRange(str, 59, 76)
}

func validateHairClr(str string) bool {
	if len(str) != 7 {
		return false
	}

	if str[0] != '#' {
		return false
	}

	for _, l := range str[1:] {
		switch l {
		case '0', '1', '2', '3', '4', '5', '6', '7', '8', '9':
		case 'A', 'B', 'C', 'D', 'E', 'F', 'a', 'b', 'c', 'd', 'e', 'f':
		default:
			return false
		}
	}

	return true
}

func validateEyeClr(str string) bool {
	switch str {
	case "amb", "blu", "brn", "gry", "grn", "hzl", "oth":
		return true
	}

	return false
}

func validatePassportId(str string) bool {
	if len(str) != 9 {
		return false
	}

	for _, l := range str {
		switch l {
		case '0', '1', '2', '3', '4', '5', '6', '7', '8', '9':
		default:
			return false
		}
	}

	return true
}

func (p Passport) String() string {
	return string(p)
}

func validateField(field string, val string) bool {
	switch field {
	case "byr":
		truth := validateIntWithinRange(val, 1920, 2002)
		// fmt.Println(field, val, truth)
		return truth
	case "iyr":
		truth := validateIntWithinRange(val, 2010, 2020)
		// fmt.Println(field, val, truth)
		return truth
	case "eyr":
		truth := validateIntWithinRange(val, 2020, 2030)
		// fmt.Println(field, val, truth)
		return truth
	case "hgt":
		truth := validateHeight(val)
		// fmt.Println(field, val, truth)
		return truth
	case "hcl":
		truth := validateHairClr(val)
		// fmt.Println(field, val, truth)
		return truth
	case "ecl":
		truth := validateEyeClr(val)
		// fmt.Println(field, val, truth)
		return truth
	case "pid":
		truth := validatePassportId(val)
		// fmt.Println(field, val, truth)
		return truth
	}
	return false
}

func (p Passport) Valid() bool {
	attrs := map[string]bool{
		"byr": false, "iyr": false, "eyr": false, "hgt": false, "hcl": false, "ecl": false, "pid": false,
	}

	for _, attr := range strings.Split(string(p), " ") {
		key := strings.Split(attr, ":")[0]
		val := strings.Split(attr, ":")[1]
		if key == "cid" {
			continue
		}
		_, ok := attrs[key]
		if !ok {
			return false
		}

		attrs[key] = validateField(key, val)
	}

	for _, v := range attrs {
		if !v {
			return false
		}
	}

	return true
}

func NewPassport(s string) Passport {
	return Passport(s)
}

func ValidNumPassport(filepath string, debug bool) int {
	strs := strings.Split(adventutil.ReadFile(filepath), "\n\n")
	passports := []Passport{}
	for _, s := range strs {
		passports = append(passports, NewPassport(strings.ReplaceAll(s, "\n", " ")))
	}

	count := 0
	for _, p := range passports {
		if p.Valid() {
			count++
		}
	}

	return count
}
