package word

import (
	"strings"
)

//UseCount returns number of timeseach word is used in given string
func UseCount(s string) map[string]int {
	xs := strings.Fields(s)
	m := make(map[string]int)
	for _, v := range xs {
		m[v]++
	}
	return m
}

//Count returns number of words in given string
func Count(s string) int {
	xs := strings.Fields(s)
	//xs := strings.Split(s, " ")
	c := len(xs)
	return c
}
