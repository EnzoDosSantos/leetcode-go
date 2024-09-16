package solutions

// PROBLEM URL: https://leetcode.com/problems/plus-one/

func PlusOne(digits []int) []int {
	n := len(digits)

	for i := n - 1; i >= 0; i-- {
		if digits[i] < 9 {
			// we just need to increment 1 if the digit is less than 9
			digits[i]++
			return digits
		}
		// if the digit is 9 we need to set it to zero
		digits[i] = 0
	}

	// if all the digits = 9, we need to put 1 at the idx[0]
	result := make([]int, n+1)
	result[0] = 1
	return result
}
