package main

type ListNode struct {
	Val int
	Next *ListNode
}

func (l *ListNode) toSlice() []int {
	cur := l
	res := []int{}
	for {
		if cur == nil {
			return res
		}
		res = append(res, cur.Val)
		cur = cur.Next
	}
}

func genList(array []int) *ListNode {
	res := &ListNode{}
	cur := res
	for _, n := range array {
		cur.Next = &ListNode{
			Val: n,
		}
		cur = cur.Next
	}
	return res.Next
}
