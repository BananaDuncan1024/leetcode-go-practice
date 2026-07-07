package main

func main() {

	root := &TreeNode{Val: 4}
	root.Left = &TreeNode{Val: 2}
	root.Right = &TreeNode{Val: 7}
	root.Left.Left = &TreeNode{Val: 1}
	root.Left.Right = &TreeNode{Val: 3}

	println(searchBST(root, 2).Val) // Output: 2
	println(searchBST(root, 5))     // Output: nil (<nil>)

}

type TreeNode struct {
	Val   int
	Left  *TreeNode
	Right *TreeNode
}

// searchBST: 利用 BST 特性，比對值往左右子樹縮小範圍
// 時間: O(h) h=樹高，空間: O(1)
func searchBST(root *TreeNode, val int) *TreeNode {
	for root != nil && root.Val != val {
		if val < root.Val {
			root = root.Left
		} else {
			root = root.Right
		}
	}
	return root
}
