package etl

import "strings"

// Transform func
func Transform(given map[int][]string) map[string]int {
	result := map[string]int{}

	for k, vs := range given {
		for _, item := range vs {
			result[strings.ToLower(item)] = k
		}
	}

	return result
}
