package day1

import (
	"fmt"
	"testing"
)

func TestDoesSumTwo(t *testing.T) {
	testcases := []struct {
		entries  []int
		sum      int
		expected [2]int
	}{
		{
			entries: []int{
				1721,
				979,
				366,
				299,
				675,
				1456,
			},
			sum:      2020,
			expected: [2]int{1721, 299}},
	}

	for _, tc := range testcases {
		ints, _ := DoesSumTwo(tc.entries, tc.sum)
		if ints != tc.expected {
			t.Log("expected", tc.expected)
			t.Log("got", ints)
			t.Fail()
		}
	}
}

func TestDoesSumThree(t *testing.T) {
	testcases := []struct {
		entries  []int
		sum      int
		expected [3]int
	}{
		{
			entries: []int{
				1721,
				979,
				366,
				299,
				675,
				1456,
			},
			sum:      2020,
			expected: [3]int{979, 366, 675}},
	}

	for _, tc := range testcases {
		ints, _ := DoesSumThree(tc.entries, tc.sum)
		if ints != tc.expected {
			fmt.Println(ints)
			fmt.Println(tc.expected)
			t.Log("expected", tc.expected)
			t.Log("got", ints)
			t.Fail()
		}
	}
}
