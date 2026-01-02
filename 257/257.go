package main

/**
 * Definition for a binary tree node.
 * type TreeNode struct {
 *     Val int
 *     Left *TreeNode
 *     Right *TreeNode
 * }
 */

import (
	"fmt"
)

type TreeNode struct {
	Val   int
	Left  *TreeNode
	Right *TreeNode
}

func createTreeByList(nums ...interface{}) *TreeNode {

	if len(nums) == 0 {
		return nil
	}

	val := nums[0]
	if val == nil {
		return nil
	}

	node := &TreeNode{Val: val.(int)}
	// Recursively create the left and right subtrees
	if len(nums) > 1 {
		node.Left = createTreeByList(nums[1:]...)
		if len(nums) > 2 {
			node.Right = createTreeByList(nums[2:]...)
		}
	}

	return node
}

func main() {

	nums := []interface{}{1, 2, 3, nil, 5}

	root := createTreeByList(nums...)

	ans := binaryTreePaths(root)

	fmt.Print(ans)

}

func binaryTreePaths(node *TreeNode) []string {

	if node == nil {
		return nil
	}

	var ans []string

	if node.Left == nil && node.Right == nil {
		ans = append(ans, fmt.Sprintf("%d", node.Val))
		return ans
	}

	left := binaryTreePaths(node.Left)
	right := binaryTreePaths(node.Right)

	for _, v := range left {
		ans = append(ans, fmt.Sprintf("%d->%s", node.Val, v))
	}

	for _, v := range right {
		ans = append(ans, fmt.Sprintf("%d->%s", node.Val, v))
	}

	return ans

}
