/**
 * Definition for singly-linked list.
 * type ListNode struct {
 *     Val int
 *     Next *ListNode
 * }
 */

func addTwoNumbers(l1 *ListNode, l2 *ListNode) *ListNode {
	dummyNode := &ListNode{}

	curr := dummyNode
	var r *ListNode = &ListNode{0, nil}
	for l1 != nil && l2 != nil {
		// Calculate
		sum := l1.Val + l2.Val + r.Val
		r = &ListNode{sum/10, nil} // sum / 10 should always be 0 if sum <= 9
		next := &ListNode{sum%10, nil}

		// Move pointers
		curr.Next = next
		curr = curr.Next
		l1 = l1.Next
		l2 = l2.Next
	}
	// Handle numbers of not the same length.
	// can't just append rest of list because remainder could propogate all the way to the end
	for l1 != nil {
		sum := l1.Val + r.Val
		r = &ListNode{sum/10, nil}
		next := &ListNode{sum%10, nil}
		curr.Next = next
		curr = curr.Next
		l1 = l1.Next
	}
	for l2 != nil {
		sum := l2.Val + r.Val
		r = &ListNode{sum/10, nil}
		next := &ListNode{sum%10, nil}
		curr.Next = next
		curr = curr.Next
		l2 = l2.Next
	}

	// Once finished if r != 0, next is r
	if r.Val != 0 {
		curr.Next = r
	}

	return dummyNode.Next	
}
