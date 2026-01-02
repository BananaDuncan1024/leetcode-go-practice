package main

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

func main() {

	head := []int{1, 2, 3, 4, 5}
	test := createLinkedList(head)
	reverseList(test)

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

func reverseList(head *ListNode) *ListNode {

	revHead := head

	if head != nil && head.Next != nil {
		revHead = reverseList(head.Next)
		head.Next.Next = head
		head.Next = nil
	}

	return revHead
}
