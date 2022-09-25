package main

type TreeNode struct {
	Val int
	Left *TreeNode
	Right *TreeNode
}

func genTree(array []int) *TreeNode {
	if len(array) == 0 {
		return nil
	}

	treeList := make([]*TreeNode, len(array))
	for i, n := range array {
		if n == -1 {
			continue
		}
		treeList[i] = &TreeNode{Val: n}
	}
	for i := range treeList {
		if treeList[i] == nil {
			continue
		}
		if i*2+1 >= len(treeList) {
			break
		}
		treeList[i].Left = treeList[i*2+1]
		if i*2+2 >= len(treeList) {
			break
		}
		treeList[i].Right = treeList[i*2+2]
	}
	return treeList[0]
}