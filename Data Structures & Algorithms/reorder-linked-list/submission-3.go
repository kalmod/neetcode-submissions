/**
 * Definition for singly-linked list.
 * type ListNode struct {
 *     Val int
 *     Next *ListNode
 * }
 */

func reorderList(head *ListNode) {
   // I can cut list in half with fast & slow walkers 
   // then reverse
   // then update?
   fast := head
   slow := head
   for fast != nil && fast.Next != nil {
	slow = slow.Next
	fast = fast.Next.Next
   }
   mid := slow.Next
   slow.Next = nil

   var prev *ListNode
   revList := mid
   for revList != nil {
	n := revList.Next
	revList.Next = prev
	prev = revList
	revList = n	
   }


   list1 := head
   list2 := prev
	for list2 != nil {
		a := list1
		b := list2
		list1 = list1.Next
		list2 = list2.Next
		a.Next = b
		b.Next = list1
	}

}
