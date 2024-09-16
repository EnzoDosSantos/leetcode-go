package solutions

// PROBLEM URL: https://leetcode.com/problems/sum-root-to-leaf-numbers/

/**
 * Definition for a binary tree node.
 * type TreeNode struct {
 *     Val int
 *     Left *TreeNode
 *     Right *TreeNode
 * }
 */

type TreeNode struct {
	Val   int
	Left  *TreeNode
	Right *TreeNode
}

func SumNumbers(root *TreeNode) int {
	return dfs(root, 0)
}

func dfs(root *TreeNode, num int) int {
	if root == nil {
		return 0
	}

	num = num*10 + root.Val

	if root.Left == nil && root.Right == nil {
		return num
	}

	return dfs(root.Left, num) + dfs(root.Right, num)
}
