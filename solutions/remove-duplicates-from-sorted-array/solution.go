package solutions

// PROBLEM URL: https://leetcode.com/problems/remove-duplicates-from-sorted-array/

func RemoveDuplicates(nums []int) int {
	left := 0

	for right, _ := range nums {
		if nums[left] != nums[right] {
			nums[left+1] = nums[right]
			left++
		}
	}
	return len(nums[:left+1])
}
