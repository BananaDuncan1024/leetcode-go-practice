package main

type TreeNode struct {
	Val   int
	Left  *TreeNode
	Right *TreeNode
}

func main() {

	root1 := &TreeNode{Val: 3}
	root1.Left = &TreeNode{Val: 5}
	root1.Right = &TreeNode{Val: 1}
	root1.Left.Left = &TreeNode{Val: 6}
	root1.Left.Right = &TreeNode{Val: 2}
	root1.Left.Right.Left = &TreeNode{Val: 7}
	root1.Left.Right.Right = &TreeNode{Val: 4}
	root1.Right.Left = &TreeNode{Val: 9}
	root1.Right.Right = &TreeNode{Val: 8}

	root2 := &TreeNode{Val: 3}
	root2.Left = &TreeNode{Val: 5}
	root2.Right = &TreeNode{Val: 1}
	root2.Left.Left = &TreeNode{Val: 6}
	root2.Left.Right = &TreeNode{Val: 7}
	root2.Right.Left = &TreeNode{Val: 4}
	root2.Right.Right = &TreeNode{Val: 2}
	root2.Right.Right.Left = &TreeNode{Val: 9}
	root2.Right.Right.Right = &TreeNode{Val: 8}

	result := leafSimilar(root1, root2)

	println(result) // Output: true
}

// leafSimilar: 先用 DFS 收集兩棵樹所有葉子序列，再逐一比對
// 時間: O(n1+n2)，空間: O(L) L=葉子數，必須跑完整棵樹才能比較
func leafSimilar(root1 *TreeNode, root2 *TreeNode) bool {
	subseq1 := []int{}
	subseq2 := []int{}

	dfs(root1, &subseq1)
	dfs(root2, &subseq2)

	if len(subseq1) != len(subseq2) {
		return false
	}
	for i := 0; i < len(subseq1); i++ {
		if subseq1[i] != subseq2[i] {
			return false
		}
	}
	return true

}

func dfs(node *TreeNode, leaves *[]int) {
	if node == nil {
		return
	}
	if node.Left == nil && node.Right == nil {
		*leaves = append(*leaves, node.Val)
	}
	dfs(node.Left, leaves)
	dfs(node.Right, leaves)
}

// leafSimilar2: iterator 版本，兩棵樹同步逐葉比對，第一個不同立即返回
// 時間: O(n1+n2) 最差，但遇到不同可提早終止，不需走完整棵樹
// 空間: O(h1+h2) 只需 stack 深度（樹高），優於原版 O(L) 葉子數
func leafSimilar2(root1 *TreeNode, root2 *TreeNode) bool {
	next1 := leafIter(root1)
	next2 := leafIter(root2)
	for {
		v1, ok1 := next1()
		v2, ok2 := next2()
		if ok1 != ok2 || v1 != v2 {
			return false
		}
		if !ok1 {
			return true
		}
	}
}

func leafIter(root *TreeNode) func() (int, bool) {
	stack := []*TreeNode{root}
	return func() (int, bool) {
		for len(stack) > 0 {
			n := stack[len(stack)-1]
			stack = stack[:len(stack)-1]
			if n.Left == nil && n.Right == nil {
				return n.Val, true
			}
			if n.Right != nil {
				stack = append(stack, n.Right)
			}
			if n.Left != nil {
				stack = append(stack, n.Left)
			}
		}
		return 0, false
	}
}
