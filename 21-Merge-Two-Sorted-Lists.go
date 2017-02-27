package leetcode_golang


func mergeTwoLists(l1 *ListNode, l2 *ListNode) *ListNode {
	var head, tail, next *ListNode
	for l1 != nil || l2 != nil {
		if l1 == nil {
			next = l2
			l2 = l2.Next
		}else if l2 == nil {
			next = l1
			l1 = l1.Next
		}else if l1.Val < l2.Val {
			next = l1
			l1 = l1.Next
		} else {
			next = l2
			l2 = l2.Next
		}
		if head == nil {
			head = next
			tail = next
		} else {
			tail.Next = next
			tail = next
		}
	}
	return head
}
