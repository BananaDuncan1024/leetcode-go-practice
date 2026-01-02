package main

func main() {

	nums := []interface{}{1, 2}

	nums2 := []interface{}{1, nil, 2}

	root := createTreeByList(nums...)

	root2 := createTreeByList(nums2...)

	isSameTree(root, root2)

}

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

func isSameTree(p *TreeNode, q *TreeNode) bool {
	return false
}
