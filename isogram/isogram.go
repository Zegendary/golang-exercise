package isogram

import "unicode"

// IsIsogram determine if a word or phrase is an isogram.
func IsIsogram(input string) bool {
	r := map[rune]bool{}

	for _, item := range input {
		_, ok := r[unicode.ToLower(item)]
		if ok {
			return false
		} else if unicode.IsLetter(item) {
			r[unicode.ToLower(item)] = true
		}
	}
	return true
}
