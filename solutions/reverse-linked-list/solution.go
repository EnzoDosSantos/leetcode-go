package solutions

// PROBLEM URL: https://leetcode.com/problems/reverse-linked-list/

/**
 * Definition for singly-linked list.
 * type ListNode struct {
 *     Val int
 *     Next *ListNode
 * }
 */

type ListNode struct {
	Val  int
	Next *ListNode
}

func ReverseList(cur *ListNode) *ListNode {
	var list *ListNode

	for cur != nil {
		// We need to keep tracking of the next head
		currentTemp := cur.Next

		// Shift the pointer to the previous node
		cur.Next = list

		// Change the pointer value
		list = cur

		// Shif the cur to real next element
		cur = currentTemp
	}

	return list
}
