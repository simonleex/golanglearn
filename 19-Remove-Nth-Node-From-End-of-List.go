package leetcode_golang

/*
Given a linked list, remove the nth node from the end of list and return its head.

For example,

Given linked list: 1->2->3->4->5, and n = 2.

After removing the second node from the end, the linked list becomes 1->2->3->5.
Note:
Given n will always be valid.
Try to do this in one pass.

Subscribe to see which companies asked this question.
*/

/**
 * Definition for singly-linked list.
 * type ListNode struct {
 *     Val int
 *     Next *ListNode
 * }
 */
type ListNode struct {
	Val int
	Next *ListNode
}
func removeNthFromEnd(head *ListNode, n int) *ListNode {
	if n == 0 {
		return head
	}
	iter := head
	span := 0
	flag := head
	for {
		if iter.Next == nil {
			if flag == head {
				return flag.Next
			}


		}
		if span < n {
			span ++
		} else {
			flag = flag.Next
		}
		if iter.Next == nil {

			flag.Next = flag.Next.Next
			break
		}
	}
	return head
}
