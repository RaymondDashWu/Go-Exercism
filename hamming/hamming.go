package hamming

import (
	"errors"
	"fmt"
)

// PSEUDO BRAINSTORM:
// compare two strings with a single for loop
// for letter_number in string1, check against string2[letter_number]
// counter, initialized at 0, + 1 for every difference

// Distance does bla blalbaalbj
func Distance(a, b string) (int, error) {

	counter := 0
	// stop := false
	if len(a) == 0 {
		fmt.Println(a)
		return 0, nil
	} else if len(a) != len(b) {
		return 0, errors.New("strings not of same length")
	} else {
		for i := range a {
			if a[i] != b[i] {
				counter++
			}
		}
	}
	return counter, nil
}
