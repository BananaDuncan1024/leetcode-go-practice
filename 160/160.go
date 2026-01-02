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

// Time : O(N+M) Space: O(1)  a kind of fast slow pointer mech solution
func getIntersectionNode(headA, headB *ListNode) *ListNode {

	if headA == nil || headB == nil {
		return nil
	}
	currentA, currentB := headA, headB
	for currentA != currentB {
		if currentA == nil {
			currentA = headB
		} else {
			currentA = currentA.Next
		}
		if currentB == nil {
			currentB = headA
		} else {
			currentB = currentB.Next
		}
	}
	return currentA
}

// Time : O(N+M) but Space: O(N)
func getIntersectionNode2(headA, headB *ListNode) *ListNode {
	seen := make(map[*ListNode]bool)
	for n := headA; n != nil; n = n.Next {
		seen[n] = true
	}

	for n := headB; n != nil; n = n.Next {
		if seen[n] {
			return n
		}
	}
	return nil
}

func main() {

	list1 := []int{4, 1, 8, 4, 5}
	list2 := []int{5, 6, 1, 8, 4, 5}

	getIntersectionNode(createLinkedList(list1), createLinkedList(list2))

}
