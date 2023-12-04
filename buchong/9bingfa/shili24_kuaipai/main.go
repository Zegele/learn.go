package main

import "fmt"

//www.kancloud.cn/imdszxs/golang/1509654
//baijiahao.baidu.com/s?id=1711157255641708615&wfr=spider&for=pc

func main() {
	Sli := []int{23, 65, 13, 27, 42, 15, 38, 21, 4, 10}
	qsort(Sli, 0, len(Sli)-1)
	fmt.Println(Sli)
}

/*
快速排序：分治法+递归实现
随意取一个值A， 将比A大的放在A的有比那，比A小的放在A的左边；
然后在左边的值AA中再取一个值b，将AA中比b小的放在b的左边，将比b大的放在b的右边
以此类推
*/

func qsort(sli []int, first, last int) {
	flag := first
	left := first
	right := first
	if first >= last {
		return
	}
	// 将大于sli[flag]的都放在右边，小于的都放在左边
	for first < last {
		// 如果flag从左边开始，那么是必须先右边开始比较，也就是先再右边找比flag小的
		for first < last {
			if sli[last] >= sli[flag] {
				last--
				continue
			}
			// 交换数据
			sli[last], sli[flag] = sli[flag], sli[last]
			flag = last
			break
		}

		for first < last {
			if sli[first] <= sli[flag] {
				first++
				continue
			}
			sli[first], sli[flag] = sli[flag], sli[first]
			flag = first
			break
		}
	}
	qsort(sli, left, flag-1)
	qsort(sli, flag+1, right)
}
