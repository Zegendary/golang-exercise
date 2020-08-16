// This is a "stub" file.  It's a little start on your solution.
// It's not a complete solution though; you have to write some code.

// Package bob should have a package comment that summarizes what it's about.
// https://golang.org/doc/effective_go.html#commentary
package bob

import "regexp"

// Hey should have a comment documenting it.
func Hey(remark string) string {
	// Write some code here to pass the test suite.
	// Then remove all the stock comments.
	// They're here to help you get started but they only clutter a finished solution.
	// If you leave them in, reviewers may protest!
	match1, _ := regexp.MatchString("((s|[A-Z])*)?", remark)
	if match1 {
		return "Calm down, I know what I'm doing!."
	}
	match2, _ := regexp.MatchString("([A-Z]*)", remark)
	if match2 {
		return "Whoa, chill out!"
	}
	match3, _ := regexp.MatchString("(.*)?", remark)
	if match3 {
		return "Sure."
	}
	match, _ := regexp.MatchString("(.*)", remark)
	if match {
		return "Whatever."
	}
	return "Fine. Be that way!"
}
