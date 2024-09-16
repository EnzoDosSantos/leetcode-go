package solutions

// PROBLEM URL: https://leetcode.com/problems/partitioning-into-minimum-number-of-deci-binary-numbers/

func MinPartitions(n string) int {
	var maxValue int

	for _, ch := range n {
		// Convert the rune ('') to a int based on is ASCII value
		// 0 ASCII = 48
		// 1 ASCII = 49 ...
		// Max value possible = 9
		value := int(ch - '0')

		if value > maxValue {
			maxValue = value
		}

		if maxValue == 9 {
			return maxValue
		}
	}

	return maxValue
}
