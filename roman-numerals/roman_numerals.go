package romannumerals

import (
	"errors"
	"fmt"
	"sort"
)

// Pseudo brainstorm
// Symbol	I	V	X	L	C	D	M
// Value	1	5	10	50	100	500	1,000
// Each symbol can only be used 3 times
// 0 can not be done

// Pseudo Brainstorm:
// Ordered dictionary that is iterated through?
// Check to see if input number is bigger than biggest (M aka 1000)
// if it is then subtract that amount up to 3 times. Otherwise move on to next roman numeral
// add all results to empty roman numeral holder

func reverseSortedKeys(m map[int]string) []int {
	// https://stackoverflow.com/questions/18342784/how-to-iterate-through-a-map-in-golang-in-order
	// https://stackoverflow.com/questions/18343208/how-do-i-reverse-sort-a-slice-of-integer-go
	keys := make([]int, len(m))
	for k, _ := range m {
		keys = append(keys, k)
	}
	sort.Sort(sort.Reverse(sort.IntSlice(keys)))
	for _, k := range keys {
		fmt.Println(k, m[k])
	}
	return keys
}

// Method 1
// func sortedKeys(m map[int]string) []int {
// 	keys := make([]int, len(m))
// 	for k, _ := range m {
// 		keys = append(keys, k)
// 	}
// 	sort.Sort(sort.Reverse(sort.IntSlice(keys)))
// 	for _, k := range keys {
// 		fmt.Sprintf(k, m[k])
// 	}
// 	return keys
// }

// Method 2
// func reverseSortedKeys(m map[int]string) ([]int, []string) {
// 	// https://stackoverflow.com/questions/18342784/how-to-iterate-through-a-map-in-golang-in-order
// 	// https://stackoverflow.com/questions/18343208/how-do-i-reverse-sort-a-slice-of-integer-go
// 	keys := make([]int, len(m))
// 	i := 0
// 	for k := range m {
// 		keys[i] = k
// 		i++
// 	}
// 	sort.Sort(sort.Reverse(sort.IntSlice(keys)))
// 	return keys, m[keys]
// }

func ToRomanNumeral(num int) (string, error) {
	// https://stackoverflow.com/questions/18342784/how-to-iterate-through-a-map-in-golang-in-order
	var romanNumeralDict map[int]string = map[int]string{
		1000: "M",
		900:  "CM",
		500:  "D",
		400:  "CD",
		100:  "C",
		90:   "XC",
		50:   "L",
		40:   "XL",
		10:   "X",
		9:    "IX",
		5:    "V",
		4:    "IV",
		1:    "I",
	}
	var romNumeralHolder = ""
	if num <= 0 {
		return "", errors.New("Not able to be converted to Roman numeral")
	}
	for k, v := range reverseSortedKeys(romanNumeralDict) {
		for num >= k {
			romNumeralHolder += v
			num -= k
		}
	}
	return romNumeralHolder, nil
}
