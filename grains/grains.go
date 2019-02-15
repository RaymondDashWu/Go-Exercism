package grains

import (
	"errors"
	"fmt"
	"math"
	"strconv"
)

// PSEUDO Brainstorm
// double numbers per iteration
// Edge cases

// expects int & bool

// convert a float64 to uint64
func conv(n float64) uint64 {
	s := fmt.Sprintf("%.0f", n)
	result, _ := strconv.ParseUint(s, 10, 64)
	return result
}

// Traverses through squares
func Square(input int) (expectedVal uint64, expectError error) {

	if input <= 0 || input > 64 {
		var err = errors.New("Error")
		return 0, err
	}
	return conv(math.Pow(2, float64(input-1))), nil
	// return uint64(math.Pow(2, float64(input))), err
}

// returns total grains at Square 64
func Total() uint64 {
	return conv(math.Pow(2, float64(64)))
}

// var total = 0
// for i := 0; i <= input && input > 0; i++ {
// 	total += i * 2
// }
// if err != nil {
// 	return 0, err
// }
// return uint64(total), err
