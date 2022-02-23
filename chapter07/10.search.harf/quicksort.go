package main

//main包最好只有一个main文件
/*
否则怪事多

func quickSort(arr *[]int64, start, end int) {
	//todo 确认终止条件，否则将无限递归下去

	pivotIdx := (start + end) / 2
	pivotV := (*arr)[pivotIdx]
	l, r := start, end
	for l <= r {
		for (*arr)[l] < pivotV {
			l++
		}
		for (*arr)[r] > pivotV {
			r--
		}
		if l >= r {
			break
		}

		(*arr)[l], (*arr)[r] = (*arr)[r], (*arr)[l]
		l++
		r--
	}

	if l == r {
		l++
		r--
	}
	if r > start {
		quickSort(arr, start, r)
	}
	if l < end {
		quickSort(arr, l, end)
	}
}
*/
