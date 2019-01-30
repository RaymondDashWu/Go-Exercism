package luhn

// PSEUDO Brainstorm:
// double every other number starting with index 0
// pass every other number starting with index 1
func Valid(num int) int {
	for i := 0; i < len(num); i += 2 {
		total := 0
		total = num[i] * 2
		if total > 9 {
			total = num[i] - 9
		}
		return total
	}
}
