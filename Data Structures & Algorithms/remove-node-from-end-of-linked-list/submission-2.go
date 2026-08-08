/**
 * Definition for singly-linked list.
 * type ListNode struct {
 *     Val int
 *     Next *ListNode
 * }
 */

func removeNthFromEnd(head *ListNode, n int) *ListNode {
	// n from the end, 1 count
	// I need to find the length
	// I would like to reference previous nodes.....hash?
	dummyHead := &ListNode{Next: head}
	left := dummyHead
	right := head
	for n > 0 {
		right = right.Next
		n--
	}
	for right != nil {
		left = left.Next
		right = right.Next
	}
	left.Next = left.Next.Next
	
	return dummyHead.Next
}
