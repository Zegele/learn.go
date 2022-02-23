package main

func midSearch(root *Node, targetNum int) bool {
	if root == nil {
		return false
	}
	totalCompare++
	if midSearch(root.left, targetNum) {
		return true
	}
	if root.data == targetNum {
		return true
	}
	if midSearch(root.right, targetNum) {
		return true
	}
	return false
}
