package utils

import "strings"

// Search method
func Search(a []string, x string) int {
	for i, n := range a {
		if strings.EqualFold(x, n) {
			return i
		}
	}
	return len(a)
}
