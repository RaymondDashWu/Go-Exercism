package luhn

// PSEUDO Brainstorm:
// double every other number starting with index 0
// pass every other number starting with index 1
func Valid(num string) bool {
	var total = 0
	for i := 0; i < len(num); i++ {
		// var total = 0
		if i%2 == 0 {
			int(num[i]) *= 2
			if num[i] > 9 {
				num[i] -= 9
				total += num[i]
			}
		} else {
			total += num[i]
		}
	}
	if total%10 == 0 {
		return true
	}
	return false
}
