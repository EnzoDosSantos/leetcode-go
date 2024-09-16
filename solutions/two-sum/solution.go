package solutions

// PROBLEM URL: https://leetcode.com/problems/two-sum/

func TwoSum(nums []int, target int) []int {
	hashMap := make(map[int]int)

	var complement int

	for i, num := range nums {
		complement = target - num

		if value, ok := hashMap[complement]; ok {
			return []int{value, i}
		}

		hashMap[num] = i
	}

	return []int{}
}
