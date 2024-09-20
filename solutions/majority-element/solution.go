package solutions

// PROBLEM URL: https://leetcode.com/problems/majority-element/

func MajorityElement(nums []int) int {
	response := 0
	counter := 0

	for _, num := range nums {
		if counter == 0 {
			response = num
		}

		if num == response {
			counter += 1
		} else {
			counter -= 1
		}
	}

	return response
}
