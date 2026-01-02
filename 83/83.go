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

func deleteDuplicates(head *ListNode) *ListNode {

	if head == nil || head.Next == nil {
		return nil
	} else if head.Val == head.Next.Val {
		head = deleteDuplicates(head.Next)

		return head
	} else {
		head.Next = deleteDuplicates(head.Next)
		return head
	}

}
func deleteDuplicates2(head *ListNode) *ListNode {
	if head == nil || head.Next == nil {
		return head
	}
	if head.Next = deleteDuplicates(head.Next); head.Val == head.Next.Val {
		head = head.Next
	}
	return head
}

func main() {

	head := []int{1, 1, 2, 3, 3}

	testdata := createLinkedList(head)

	deleteDuplicates(testdata)

}
