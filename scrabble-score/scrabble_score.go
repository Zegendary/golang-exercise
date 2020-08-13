package scrabble

import "unicode"

var scoreMap = map[rune]int{
	65: 1,
	66: 3,
	67: 3,
	68: 2,
	69: 1,
	70: 4,
	71: 2,
	72: 4,
	73: 1,
	74: 8,
	75: 5,
	76: 1,
	77: 3,
	78: 1,
	79: 1,
	80: 3,
	81: 10,
	82: 1,
	83: 1,
	84: 1,
	85: 1,
	86: 4,
	87: 4,
	88: 8,
	89: 4,
	90: 10,
}

// Score scrabble
func Score(input string) int {
	score := 0
	inputRune := []rune(input)
	for _, item := range inputRune {
		score += scoreMap[unicode.ToUpper(item)]
	}
	return score
}
