package main

type ListNode struct {
	Val  int
	Next *ListNode
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

// using slow fast poiner or using flag to mark the data
func hasCycle(head *ListNode) bool {

	slow := head
	fast := head

	for fast != nil && fast.Next != nil {

		slow = slow.Next
		fast = fast.Next.Next

		if slow == fast {
			return true
		}

	}

	return false
}

func main() {

	head := []int{3, 2, 0, -4}
	// pos := 1

	createLinkedList(head)

	hasCycle(createLinkedList(head))

}
