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
	m := make(map[int]*ListNode)
	dummyHead := &ListNode{}
	dummyHead.Next = head
	curr := dummyHead
	i := 0
	for curr != nil {
		m[i] = curr
		curr = curr.Next
		i++
	}
	targetNode :=  i - n
	m[targetNode-1].Next = m[targetNode].Next
	
	return dummyHead.Next
}
