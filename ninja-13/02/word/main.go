// Package word provides functions for working with strings
package word

import "strings"

// UseCount gets a string with words and returns the frequency of each word
func UseCount(s string) map[string]int {
	xs := strings.Fields(s)
	m := make(map[string]int)
	for _, v := range xs {
		m[v]++
	}
	return m
}

// Count counts the number of words in a string and returns the result as an int
func Count(s string) int {
	xs := strings.Fields(s)
	return len(xs)
}
