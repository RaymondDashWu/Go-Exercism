package isogram

import (
	"strings"
	"unicode"
)

// PSEUDO BRAINSTORM:
// iterate through letters in word
// append to array
// iterate through array to see if letter exists
// unicode.IsLetter
// emptySlice = {B: 1, A: 1, T: 1}
// BAT
func IsIsogram(a string) bool {
	// // hash table method
	empty_dict := make(map[string]int)
	for i := 0; i < len(a); i++ {
		if unicode.IsLetter(rune(a[i])) {
			empty_dict[strings.ToUpper(string(a[i]))]++
			if empty_dict[strings.ToUpper(string(a[i]))] > 1 {
				return false
			}
		}
	}
	return true

	// var emptySlice = []string{}
	// for i := 0; i < len(a); i++ {
	// 	for _, v := range emptySlice {
	// 		if v == string(a[i]) {
	// 			return false
	// 		}
	// 	}
	// }

	// hamsterman
	// {h, a, m}
	// a = strings.ToLower(a)
	// var emptySlice = []string{}
	// // m := make(map[string]int)
	// for i := 0; i < len(a); i++ {
	// 	// rune will return int?
	// 	if unicode.IsLetter(rune(a[i])) {
	// 		// emptySlice = append(emptySlice, string(a[i]))
	// 		// fmt.Println("Appended:", emptySlice)
	// 		// check if letter in emptySlice

	// 		for _, v := range emptySlice {
	// 			// emptySlice = append(emptySlice, string(a[i]))
	// 			if v == string(a[i]) {
	// 				return false
	// 			}
	// 		}
	// 		emptySlice = append(emptySlice, string(a[i]))
	// 		// if letter in emptySlice:
	// 		// return False
	// 	} else {
	// 		continue
	// 	}
	// }
	// return true
}
