package luhn

import (
	"strings"
)

// PSEUDO Brainstorm:
// double every other number starting with index 0
// pass every other number starting with index 1
// return all of these numbers added together
// if total % 10 = 0, return true; else false

// func Valid(num string) bool {
// 	num = strings.Replace(num, " ", "", -1)
// 	if unicode.IsNumber(rune(num)) {
// 		fmt.Println("test")
// 	} else {
// 		return false
// 	}
// }

func Valid(num string) bool {
	// fmt.Println("Tested against num:", num)
	var total = 0
	// strip spaces
	num = strings.Replace(num, " ", "", -1)

	if len(num) <= 1 {
		return false
	}

	for i, j := 0, len(num)-1; i < j; i, j = i+1, j-1 {
		if j % 2 == 0 {
			
		}
	}
	return string(empty_string)

}

// 	// if even number, counter starts at 0. If odd, counter starts at 1. This possibly accounts for starting from right?

// 	// Doesn't work for some reason?
// 	// if unicode.IsNumber(rune(sterile_string)) {
// 	if unicode.IsNumber(rune(num)) {
// 		for i := 0; i < len(num); i++ {

// 			//checks if
// 			if unicode.IsNumber(rune(num[i])) {
// 				digit, _ := strconv.Atoi(string(num[i]))
// 				fmt.Println("digit:", digit)
// 				// Below line does not work because starts at index 0
// 				// TODO: Does not get counter number properly to %2 and possibly does not add them correctly
// 				if counter%2 == 0 {
// 					// Int(num[i]) *= 2
// 					// digit, _ := strconv.Atoi(string(num[i]))
// 					fmt.Println("Counter:", counter)
// 					fmt.Println("total in first if:", total)
// 					fmt.Println("Digit index divisible by 2. Double:", digit)
// 					digit *= 2
// 					if digit > 9 {
// 						digit -= 9
// 						total += digit
// 						counter++
// 					} else {
// 						total += digit
// 						counter++
// 					}

// 				} else {
// 					total += digit
// 					counter++
// 					fmt.Println("total in else:", total)
// 					fmt.Println("Counter:", counter)
// 				}
// 			}
// 		}
// 	}
// 	if total%10 == 0 && len(num) > 1 {
// 		return true
// 	}
// 	return false
// }
