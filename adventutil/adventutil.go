package adventutil

import (
	"fmt"
	"io/ioutil"
	"strconv"
	"strings"
)

func ReadFile(filepath string) string {
	buf, err := ioutil.ReadFile(filepath)
	if err != nil {
		panic(fmt.Sprintf("cant read file, %s", err))
	}
	return strings.TrimSpace(string(buf))
}

func ToIntSlice(str string) []int {
	strEntries := strings.Split(str, "\n")
	entries := make([]int, len(strEntries))
	for i, e := range strEntries {
		entries[i], _ = strconv.Atoi(e)
	}
	return entries
}
