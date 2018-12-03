// Package word Functions for strings
package word

import "strings"

// no need to write an example for this one.
// Writing a test for this one is a bonus challenge; harder

// UseCount Returns the number of times each word is used in a string
func UseCount(s string) map[string]int {
	xs := strings.Fields(s)
	m := make(map[string]int)
	for _, v := range xs {
		m[v]++
	}
	return m
}

// Count returns the number of words in a string.
func Count(s string) int {
	xs := strings.Split(s, " ")
	count := len(xs)
	return count
	// return len(strings.Fields(s))
}
