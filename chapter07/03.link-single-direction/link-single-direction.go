package main

import (
	"fmt"
)

type LinkNode struct { //定义了链的最小单元  data是数据， next相当于锚点
	data int
	next *LinkNode
}

func main() {
	n1 := &LinkNode{
		data: 1,
		next: nil,
	}
	n2 := &LinkNode{
		data: 2,
		next: nil,
	}
	n3 := &LinkNode{
		data: 3,
		next: nil,
	}
	n4 := &LinkNode{
		data: 4,
		next: nil,
	}
	n6 := &LinkNode{
		data: 6,
		next: nil,
	}

	n1.next = n2
	n2.next = n3
	n3.next = n4
	n4.next = n6

	{
		rangeLink(n1)
	}

	{
		fmt.Println("插入 5")
		n5 := &LinkNode{
			data: 5,
			next: nil,
		}
		insertNode(n1, n5)
		insertNode(n1, &LinkNode{
			data: 7,
			next: nil,
		})
		insertNode(n1, &LinkNode{
			data: 5,
			next: nil,
		})
		insertNode(n1, &LinkNode{
			data: 3,
			next: nil,
		})

		insertNode(n1, &LinkNode{
			data: 0,
			next: nil,
		})
		rangeLink(n1)
	}

	{
		fmt.Println("测试插到第一个")
		n0 := &LinkNode{
			data: 0,
			next: nil,
		}
		n10 := &LinkNode{
			data: 10,
			next: nil,
		}
		ne := insertNode(n10, n0)
		rangeLink(ne)
	}

	//{
	//	fmt.Println("测试nil节点")
	//	var n *LinkNode //没有初始化（实例化）的Node是没法用的。
	//	n10 := &LinkNode{
	//		data: 10,
	//		next: nil,
	//	}
	//	fmt.Println(n)
	//	insertNode(n, n10)
	//	rangeLink(n)
	//}
	{
		fmt.Println("删除节点")
		n1 = deleteNode(n1, 3)
		n1 = deleteNode(n1, 5)
		n1 = deleteNode(n1, 1)
		rangeLink(n1)
	}
}

func rangeLink(root *LinkNode) {
	tmpNode := root
	for {
		if tmpNode != nil {
			fmt.Println(tmpNode.data)
		} else {
			break
		}
		tmpNode = tmpNode.next //这就是链上了
	}
}

func insertNode(root *LinkNode, newNode *LinkNode) *LinkNode {
	tmpNode := root
	if newNode.data < root.data {
		newNode.next = root
		return newNode
	}
	for {
		if tmpNode != nil {
			if newNode.data > tmpNode.data {
				if tmpNode.next == nil { //如果已经到结尾了，直接追加
					tmpNode.next = newNode
				} else { //如果不是结尾
					if tmpNode.next.data >= newNode.data {
						//找到合适位置，准备插入数据
						newNode.next = tmpNode.next
						tmpNode.next = newNode
						break
					}
				}
			}
		} else { //如果节点是nil的情况 newNode 成为链的第一个
			break
		}
		tmpNode = tmpNode.next
	}
	return tmpNode
}

func deleteNode(root *LinkNode, data int) *LinkNode {
	tmpNode := root
	if root != nil && root.data == data {
		if root.next == nil { //说明这个链就这一个值
			return nil // 给个nil，就说明把这只有一个值的删了
		}
		rightRoot := root.next
		tmpNode.next = nil
		return rightRoot //删掉第一个
	}
	for {
		if tmpNode.next == nil {
			break
		}
		right := tmpNode.next
		if right.data == data {
			//找到要删除的节点，开始删除
			tmpNode.next = right.next
			right.next = nil
			return root
		}
		tmpNode = tmpNode.next
	}
	return root
}
