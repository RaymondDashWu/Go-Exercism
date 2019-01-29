package isogram

import "unicode"

// PSEUDO BRAINSTORM:
// iterate through letters in word
// append to array
// iterate through array to see if letter exists
// unicode.IsLetter
// emptySlice = {B: 1, A: 1, T: 1}
// BAT
func IsIsogram(a string) bool {
	var emptySlice = []string{}
	// m := make(map[string]int)
	for i := 0; i < len(a); i++ {
		// rune will return int?
		if unicode.IsLetter(rune(a[i])) {
			emptySlice = append(emptySlice, string(a[i]))
			// check if letter in emptySlice

			for ind, v := range emptySlice {
				// emptySlice = append(emptySlice, string(a[i]))
				if v == emptySlice[ind] {
					return false
				}
			}
			// if letter in emptySlice:
			// return False
		} else {
			continue
		}
	}
	return true
}

// for _, b := range list {
// 	if b == a {
// 		return true
// 	}
