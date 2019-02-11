package summultiples

import "fmt"

// PSEUDO BRAINSTORM:
// Given a number find multiples up to that number that are divisible by given numbers then add them together
// Ex: Given 20, find UNIQUE numbers up to but not including 20 that are divisible by 3 & 5
// Solution: 3, 5, 6, 9, 10, 12, 15, and 18. Added together equals 78

// start out with a total variable set to 0. Will get added to as modulo operator finds numbers
// for all numbers in the max number - 1 divide by array of ints in divisors
// how to best go through array of divisors? Another for loop adds would make it n^2 operation

func SumMultiples(limit int, divisors []int) int {
	var sum = 0
	// temporary slice
	temp_map := make(map[int]int)
	for i := 0; i < limit; i++ {
		for _, v := range divisors {
			if i%v == 0 {
				if _, ok := temp_map[divisors[i]]; ok {
					temp_map[divisors[i]]++
				} else {
					temp_map[divisors[i]] = 1
				}

				// STUCK. Should check if temp_slice contains number currently being used in divisors aka divisors[i]
				if temp_map[divisors[i]] < 1 {
					// append to slice
					// STUCK. Is this how you append to a make slice?
					temp_map[i]++
					// temp_map = append(temp_map, i)
					sum += i
					fmt.Println("number:", i)
					fmt.Println("divisible by:", v)
					fmt.Println("sum:", sum)
					fmt.Println("")
				}
			}
		}
	}
	return sum
}
