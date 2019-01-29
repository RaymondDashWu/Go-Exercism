package raindrops

import "strconv"

// If the number has 3 as a factor, output 'Pling'.
// If the number has 5 as a factor, output 'Plang'.
// If the number has 7 as a factor, output 'Plong'.
// PSEUDO BRAINSTORM:
//

func Convert(a int) string {
	result := ""
	if a%3 == 0 {
		result += "Pling"
	}
	if a%5 == 0 {
		result += "Plang"
	}
	if a%7 == 0 {
		result += "Plong"
	}
	if result == "" {
		result += strconv.Itoa(a)
	}
	return result
}
