package main

import "fmt"

/**
 * Definition for singly-linked list.
 * type ListNode struct {
 *     Val int
 *     Next *ListNode
 * }
 */

type ListNode struct {
	Val  int
	Next *ListNode
}

func mergeTwoLists(list1 *ListNode, list2 *ListNode) *ListNode {
	res := &ListNode{}
	head := res

	for list1 != nil || list2 != nil {
		if list1 == nil {
			//Return case
			res.Next = list2
			return head.Next
		} else if list2 == nil {
			res.Next = list1
			return head.Next
		} else {
			if list1.Val <= list2.Val {
				res.Next = &ListNode{Val: list1.Val, Next: nil}
				list1 = list1.Next
				res = res.Next
			} else {
				res.Next = &ListNode{Val: list2.Val, Next: nil}
				list2 = list2.Next
				res = res.Next
			}
		}

	}
	return head.Next

}

