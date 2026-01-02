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

func mergeTwoLists(list1 *ListNode, list2 *ListNode) *ListNode {

	dummyNode := &ListNode{}
	cusros := dummyNode
	current1 := list1
	current2 := list2

	for current1 != nil && current2 != nil {

		if current1.Val <= current2.Val {

			cusros.Next = current1
			current1 = current1.Next
		} else {
			cusros.Next = current2
			current2 = current2.Next
		}

		cusros = cusros.Next
	}

	if current1 == nil {
		cusros.Next = current2
	} else {
		cusros.Next = current1
	}

	return dummyNode.Next
}

func mergeTwoLists2(list1 *ListNode, list2 *ListNode) *ListNode {
	if list1 == nil {
		return list2
	}
	if list2 == nil {
		return list1
	}
	if list1.Val <= list2.Val {
		list1.Next = mergeTwoLists(list1.Next, list2)
		return list1
	}
	list2.Next = mergeTwoLists(list2.Next, list1)
	return list2
}

func main() {

	list1 := []int{1, 2, 4}
	list2 := []int{1, 3, 4}

	linkList1 := createLinkedList(list1)
	linkList2 := createLinkedList(list2)

	mergeTwoLists(linkList1, linkList2)

}
