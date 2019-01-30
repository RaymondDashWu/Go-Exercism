package diffsquares

func SquareOfSum(num int) int {
	var total = 0
	for i := 0; i <= num; i++ {
		total += i
	}
	return total * total
}

func SumOfSquares(num int) int {
	var total = 0
	for i := 0; i <= num; i++ {
		total += i * i
	}
	return total
}

func Difference(sums int) int {
	return SquareOfSum(sums) - SumOfSquares(sums)
}
