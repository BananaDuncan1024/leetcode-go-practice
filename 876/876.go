package main

type ListNode struct {
	Val  int
	Next *ListNode
}

func middleNode(head *ListNode) *ListNode {

	store := make([]*ListNode, 0, 10)

	for head != nil {

		store = append(store, head)
		head = head.Next

	}

	return store[len(store)/2]
}

func createLinkedList(nums []int) *ListNode {
	// Create a dummy node as the head of the linked list
	dummy := &ListNode{}
	current := dummy

	// Iterate over the slice and create new nodes
	for _, num := range nums {
		node := &ListNode{Val: num}
		current.Next = node
		current = current.Next
	}

	return dummy.Next
}

func main() {

	head := []int{1, 2, 3, 4, 5}

	test := createLinkedList(head)

	middleNode(test)

}
