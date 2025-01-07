package plus_One

func PlusOne(digits []int) []int {

	n := len(digits)

	// Start from the last digit and go backward
	for i := n - 1; i >= 0; i-- {
		if digits[i] < 9 { // if the digit is less than 9, just increment it by 1
			digits[i]++
			return digits
		}
		// Otherwise, set it to 0 and continue to the next more significant digit
		digits[i] = 0
	}

	// If we finished the loop without returning, it means we had all 9's
	// For example: 999 + 1 = 1000, we need to add a leading 1
	return append([]int{1}, digits...)

}
