package main

/*
给你两个 非空 的链表，表示两个非负的整数。它们每位数字都是按照 逆序 的方式存储的，并且每个节点只能存储 一位 数字。

请你将两个数相加，并以相同形式返回一个表示和的链表。

你可以假设除了数字 0 之外，这两个数都不会以 0 开头。
 */

/**
 * Definition for singly-linked list.
 * type ListNode struct {
 *     Val int
 *     Next *ListNode
 * }
 */
func addTwoNumbers(l1 *ListNode, l2 *ListNode) *ListNode {
	res := &ListNode{}
	l1_cur, l2_cur, res_cur := l1, l2, res
	carry := 0
	for {
		if l1_cur == nil && l2_cur == nil && carry == 0 {
			return res.Next
		}
		sum := carry
		if l1_cur != nil {
			sum += l1_cur.Val
			l1_cur = l1_cur.Next
		}
		if l2_cur != nil {
			sum += l2_cur.Val
			l2_cur = l2_cur.Next
		}
		carry = sum / 10
		res_cur.Next = &ListNode{
			Val: sum % 10,
		}
		res_cur = res_cur.Next
	}
}

// input:
// genList([]int{2,4,3}), genList([]int{5,6,4})
