package hamming

import (
	"errors"
)

// Distance function
func Distance(a, b string) (int, error) {
	count := 0
	aRunes := []rune(a)
	bRunes := []rune(b)

	if len(aRunes) != len(bRunes) {
		return 0, errors.New("disallow")
	}

	for i := 0; i <= len(aRunes)-1; i++ {
		if aRunes[i] != bRunes[i] {
			count++
		}
	}

	return count, nil
}
