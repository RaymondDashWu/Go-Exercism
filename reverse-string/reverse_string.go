package reverse

func String(input string) string {
	// var blank_string = ""

	// for i := 0; i < len(input); i++ {
	// 	blank_string += string(input[i])
	// }

	// for i, j := 0, len(empty_string)-1; i < j; i, j = i+1, j-1 {
	// 	empty_string[i], empty_string[j] = empty_string[j], empty_string[i]
	// 	fmt.Println(empty_string)
	// }
	// return strings.Join(empty_string, ", ")

	// fmt.Println("input:", input)
	empty_string := []rune(input)
	// fmt.Println("empty_string:", empty_string)

	for i, j := 0, len(empty_string)-1; i < j; i, j = i+1, j-1 {
		empty_string[i], empty_string[j] = empty_string[j], empty_string[i]
		// fmt.Println(empty_string)
	}
	return string(empty_string)

	// func String(s string) string {
	// 	runes := []rune(s)
	// 	for i, j := 0, len(runes)-1; i < j; i, j = i+1, j-1 {
	// 		runes[i], runes[j] = runes[j], runes[i]
	// 	}
	// 	return string(runes)

	// strings.Join(arr []string, seperator string) string
	// func main() {

	// 	s := []int{5, 2, 6, 3, 1, 4}

	// 	for i, j := 0, len(s)-1; i < j; i, j = i+1, j-1 {
	// 		s[i], s[j] = s[j], s[i]
	// 	}

	// 	fmt.Println(s)

	// rune_string := [] rune(input)
	// for range(len(input)) {
	// 	newVal
	// 	newVal := input[len(input)-1:]

	// }
	// return newVal

	// 	var negative_counter = -1
	// 	var empty_string = ""
	// 	for i := 0; i = len(input); i++ {
	// 		empty_string += string(input[negative_counter])
	// 		negative_counter--
	// 	}
	// 	return empty_string
	// }

	// func String(s string) string {
	// 	runes := []rune(s)
	// 	for i, j := 0, len(runes)-1; i < j; i, j = i+1, j-1 {
	// 		runes[i], runes[j] = runes[j], runes[i]
	// 	}
	// 	return string(runes)
}
