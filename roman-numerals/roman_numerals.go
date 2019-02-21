package romannumerals

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

func ToRomanNumeral(num int) (string, bool) {
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
	if num == 0 {
		return "", false
	}
	for k, v := range romanNumeralDict {
		if num >= k {
			romNumeralHolder += v
			num -= k
		}
	}
	return romNumeralHolder, false
}
