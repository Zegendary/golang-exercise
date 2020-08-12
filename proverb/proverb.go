// This is a "stub" file.  It's a little start on your solution.
// It's not a complete solution though; you have to write some code.

// Package proverb should have a package comment that summarizes what it's about.
// https://golang.org/doc/effective_go.html#commentary
package proverb

import "fmt"

const (
	stanza = "For want of a %s the %s was lost."
	last   = "And all for the want of a %s."
)

// Proverb should have a comment documenting it.
func Proverb(rhyme []string) []string {
	// Write some code here to pass the test suite.
	// Then remove all the stock comments.
	// They're here to help you get started but they only clutter a finished solution.
	// If you leave them in, reviewers may protest!
	var result []string
	for index, item := range rhyme {
		if index == len(rhyme)-1 {
			result = append(result, fmt.Sprintf(last, rhyme[0]))
		} else {
			result = append(result, fmt.Sprintf(stanza, item, rhyme[index+1]))
		}
	}
	return result
}
