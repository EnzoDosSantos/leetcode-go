package solutions

import "slices"

// PROBLEM URL: https://leetcode.com/problems/longest-subarray-with-maximum-bitwise-and/

func LongestSubarray(nums []int) int {
	maxVal := slices.Max(nums[:])

	maxLen := 0
	currentLen := 0

	for _, n := range nums {
		if n == maxVal {
			currentLen++
		} else {
			if currentLen > maxLen {
				maxLen = currentLen
			}
			currentLen = 0
		}
	}

	if currentLen > maxLen {
		maxLen = currentLen
	}

	return maxLen
}
