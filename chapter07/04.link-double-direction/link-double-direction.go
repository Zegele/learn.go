package main

import "fmt"

// type CompareFunc func(left, right interface{} bool

type LinkNode struct {
	data     int //interface{} data用interface 一般用一个接口或方法，进行比较
	next     *LinkNode
	previous *LinkNode
}

func buildDLink() *LinkNode {
	n1 := &LinkNode{data: 1}
	n2 := &LinkNode{data: 5}
	n3 := &LinkNode{data: 10}

	n1.next = n2
	n2.previous = n1

	n2.next = n3
	n3.previous = n2

	return n1
}

func insertNode(root *LinkNode, newNode *LinkNode) *LinkNode {
	tmpNode := root

	//整个链表都是空的情况，新增
	if root == nil {
		return newNode
	}

	// 在链表的头，添加节点
	if root.data >= newNode.data {
		newNode.next = tmpNode
		tmpNode.previous = newNode

		return newNode
	}

	for {
		if tmpNode.next == nil {
			//已经到尾部了，追加节点即可
			tmpNode.next = newNode
			newNode.previous = tmpNode
			return root //因为tmpNode 和 root 都是指针类型的，指的是同一个地址
		} else {
			if tmpNode.next.data >= newNode.data {
				newNode.previous = tmpNode  //新节点的前端接到tmpNode
				newNode.next = tmpNode.next //新节点的后端接到tmpNode.Next

				newNode.next.previous = newNode //原节点的后端（tmpNode.Next）此时已经成为新节点的后端（newNode.next）,它的前端要链接newNode
				tmpNode.next = newNode          //原节点的后端，链接newNode000

				return root
			}
		}
		tmpNode = tmpNode.next //对比下一个
	}
}
func deleteNode(root *LinkNode, v int) *LinkNode {
	if root == nil {
		return nil
	}

	if root.data == v {
		//要删除的数据在第一个节点
		leftHand := root
		root = root.next

		leftHand.next = nil
		root.previous = nil

		// todo 需要解决只有一个节点的情况

		return root
	}

	tmpNode := root
	for {
		if tmpNode.next == nil { //带着循环思维理解，经过循环，直到tmpNode.next ==nil 也没有找到data = v的。说明链表中没有要删除的该数字。
			//走到链表的尾部，任然没有找到要删除的数据，直接返回原root
			return root
		} else {
			if tmpNode.next.data == v {
				// 找到节点，开始删除，删除完成后返回原root
				rightHand := tmpNode.next
				tmpNode.next = rightHand.next
				rightHand.next.previous = tmpNode

				//清理掉右手上的link，保证GC正常回收
				rightHand.next = nil
				rightHand.previous = nil

				return root
			}
		}
		tmpNode = tmpNode.next
	}
}

func deleteNodeZiJi(root *LinkNode, v int) *LinkNode {

	if root == nil {
		return nil
	}

	if root.data == v { // 要删除的数据在第一个节点（有两种情况，1.链表有多个节点；2.链表只有一个节点）
		if root.next != nil {
			leftHand := root
			root = root.next

			leftHand.next = nil
			root.previous = nil //这个root已经是原来的root.next了
		} else { //链表只有这一个节点
			root = nil
		}
		return root
	}

	tmpNode := root
	for {
		if tmpNode.next.data == v {
			if tmpNode.next.next == nil { // 删掉最后一个节点

				tmpNode.next.previous = nil
				tmpNode.next = nil //如果顺序颠倒就会错。如果先tmpNode.next = nil ,然后tmpNode.next.previous 就找不到了，更不用说赋值了。
				return root
			} else { //删除中间的节点
				tmpNode.next.next.previous = tmpNode
				tmpNode.next = tmpNode.next.next
				return root
			}
		}
		tmpNode = tmpNode.next
		if tmpNode.next == nil { //找遍了所有的节点，都没有匹配的数，就返回原链表
			fmt.Println("没有找到要删除的数字对应的节点")
			return root
		}
	}

}

func rangeLink(root *LinkNode) {
	fmt.Println("从头到尾")
	tmpNode := root

	for {
		fmt.Println(tmpNode.data)
		if tmpNode.next == nil {
			break
		}
		tmpNode = tmpNode.next
	}

	//fmt.Println("从尾到头")
	//for {
	//	fmt.Println(tmpNode.data)
	//	if tmpNode.previous == nil {
	//		break
	//	}
	//	tmpNode = tmpNode.previous
	//}
}

func main() {
	root := buildDLink()

	root = insertNode(root, &LinkNode{data: 3})
	root = insertNode(root, &LinkNode{data: 7})
	root = insertNode(root, &LinkNode{data: 11})
	root = insertNode(root, &LinkNode{data: 0})
	root = insertNode(root, &LinkNode{data: -1})

	rangeLink(root)

	fmt.Println("删除节点")
	fmt.Println("删除第一个节点：")
	root = deleteNodeZiJi(root, -1) // 删除第一个
	rangeLink(root)

	fmt.Println("删除中间节点：")
	root = deleteNodeZiJi(root, 3) // 删除中间的
	rangeLink(root)

	fmt.Println("删除最后一个节点")
	root = deleteNodeZiJi(root, 11) // 删除最后一个
	rangeLink(root)

	fmt.Println("删除不存在的节点")
	root = deleteNodeZiJi(root, 4) //删除不存在的
	rangeLink(root)

	//root = deleteNode(root, -1)
	//root = deleteNode(root, 3)
	//root = deleteNode(root, 11)
	//rangeLink(root)
}
