// This is a "stub" file.  It's a little start on your solution.
// It's not a complete solution though; you have to write some code.

// Package proverb should have a package comment that summarizes what it's about.
// https://golang.org/doc/effective_go.html#commentary

// Pseudo brainstorm:
// two pointers, x + 1, x + 2
// once one of the pointers has reached the end return the last line

// Notes:
// For want of a * the * was lost.
// And all for the want of a *.
package proverb

// Proverb should have a comment documenting it.
func Proverb(rhyme []string) []string {
	// Write some code here to pass the test suite.
	// Then remove all the stock comments.
	// They're here to help you get started but they only clutter a finished solution.
	// If you leave them in, reviewers may protest!

	empty_string := []rune(input)
	// fmt.Println("empty_string:", empty_string)

	for i, j := 0, len(rhyme)-1; i > j; i, j = i+1, j+2 {
		empty_string[i], empty_string[j] = empty_string[j], empty_string[i]
		// fmt.Println(empty_string)
	}
	return string(empty_string)

	return []string{}
}
