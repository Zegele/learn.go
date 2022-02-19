package main

import (
	"fmt"
	"math/rand"
	"time"
)

func quickSort(arr *[]int, start, end int) {
	// todo 确认终止条件，否则将无限递归下去

	pivotIdx := (start + end) / 2 //选了中间的数，为标识 start=0; end=9 pivotIdx = 4
	pivotV := (*arr)[pivotIdx]    //标识位对应的数是pivotV
	//fmt.Println(pivotV)
	l, r := start, end
	for l <= r { //老师这样写for l <= r {，刚开始觉得应该不用=，经过推理后还是要有的
		for (*arr)[l] < pivotV { //这里不能有 = 号
			// 如案例一： {6 6 6} 就会一直加，超出切片长度；
			//案例2：{9 6} 如果有等号，一次循环后，r = l 都指着6，然后就结束了（一直调用递归），但一直没有排序。死递归
			l++
		}
		for (*arr)[r] > pivotV {
			r--
		}
		// 缺少这三行导致失败的数组： 45 12 42 33 10 44 0 27 27 20
		if l >= r { //在这里交错了 // 课堂上我们没有在这里break导致出现意外的排序失败。
			//->a b c<-  ---  a  ->b<-  c 这个情况是 l = r
			// ->a b c d<-  ---- a ->b c<- d ---- a  b<- ->c d 这种情况是 l > r
			break // 课堂上我们没有在这里break导致出现意外的排序失败
		} // 课堂上我们没有在这里break导致出现意外的排序失败
		(*arr)[l], (*arr)[r] = (*arr)[r], (*arr)[l]
		l++
		r--
		//这里++，--后有可能l==r 也要继续进行运算
	}
	// fmt.Println("l:",l, "r:",r)
	// fmt.Println(*arr)
	if l == r {
		l++
		r--
	} //这里会让这种情况 a  ->b<-  c 变成 a<- b ->c 这样下一轮，就不包括b
	//这里++ ，--没有使用切片，所以不存在溢出可能。只是为了错开，或比较start和end以判断是否调用下一次。

	if r > start {
		quickSort(arr, start, r) //交错形成2个新的slice， 递归
	}
	if l < end {
		quickSort(arr, l, end)
	}
}

func main() {
	arrSize := 10
	arr := []int{}
	rand.Seed(time.Now().UnixNano())
	for i := 0; i < arrSize; i++ {
		arr = append(arr, rand.Intn(50))
	}
	fmt.Println("初始数组：", arr)

	start := time.Now()
	quickSort(&arr, 0, arrSize-1)
	finish := time.Now()
	fmt.Println(finish.Sub(start))
	fmt.Println("最终数组：", arr)
}

// 1 5 3 4 5 1 2 9 6
//		   5
// 1 2 3 4 1 | 5 5 9 6
//         r   l
//	   3       5
// 1 2 1 | 4 3      5 | 5 9 6
//         r(l)     r   l  //交换后，r和l都在4的位置 r == l ，还要经过计算
//                         //应该就一次计算，因为这个数和标识数最后就只有一个结果，要么大，要么小，要么等于），
//                         //结果就会让r，l交错开，形成新的切片，然后继续递归
//     r   l        r   l
//   2   | 4        5 |   9
// 1 1 | 2   4        5   5 | 6 9
//   r   l				  l   r
// 1 1
// 1     2   4        5   5   6
// 1 1   2   4        5   5   6 9
// r l                        r(l)
// 1 1   2   4        5   5   6 9

//其实就是到最后一层解决以下3种情况：
// 1 2    1 1   1 0
// r(l)   r l   0 1
// r(l)   r l   r l

// 人工测验 for (*arr)[l] <= pivotV 和 for (*arr)[l] <= pivotV
// 测验有等号会怎样 ？ 结论： 1. 溢出 ；2. 无限递归
// 1 5 3 4 5 1 2 9 6
//		   5
// 1 5 3 4 5 1 2 | 9 6
//             r | l
// 		 4		   9
// 1 2 3 4 1 | 5 5   9 6
//         r | l	   r(l) 96不排序了 且一直调用递归
//     3       5     9
// 1 2 3 1 4   5 5   9 6
//       r l   r l     r(l)
//   2     4   5 5
// 1 2 1 3 4   5 5
//     r l
//   2
// 1 2 1
//     r l （l溢出了）
