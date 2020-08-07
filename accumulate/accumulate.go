package accumulate

// Accumulate string with func
func Accumulate(strings []string, fn func(string) string) []string {
	result := make([]string, len(strings))
	for i, item := range strings {
		result[i] = fn(item)
	}

	return result
}
