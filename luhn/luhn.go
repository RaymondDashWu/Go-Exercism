package luhn

import (
	"fmt"
	"strconv"
	"unicode"
)

// PSEUDO Brainstorm:
// double every other number starting with index 0
// pass every other number starting with index 1
// return all of these numbers added together
// if total % 10 = 0, return true; else false
func Valid(num string) bool {
	fmt.Println("Tested against num:", num)
	var total = 0
	for i := 0; i < len(num); i++ {
		if unicode.IsNumber(rune(num[i])) {
			digit, _ := strconv.Atoi(string(num[i]))
			fmt.Println("digit:", digit)
			// Below line does not work because starts at index 0
			if i%2 == 0 {
				// Int(num[i]) *= 2
				digit, _ := strconv.Atoi(string(num[i]))
				digit *= 2
				total += digit * 2
				fmt.Println("total in first if:", total)
				if digit > 9 {
					digit -= 9
					total += digit
				}
			} else {
				total += digit
				fmt.Println("total in else:", total)
			}
		}
	}
	if total%10 == 0 && len(num) > 1 {
		return true
	}
	return false
}
