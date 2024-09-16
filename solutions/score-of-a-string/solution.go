package solutions

import "math"

// PROBLEM URL: https://leetcode.com/problems/score-of-a-string/

func ScoreOfString(s string) int {
	bytes := []byte(s)
	count := 0

	for i := 0; i < len(bytes)-1; i++ {
		toSum := int(bytes[i])
		toRest := int(bytes[i+1])

		count += int(math.Abs(float64(toSum - toRest)))
	}

	return count
}
