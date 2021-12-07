package main

import "fmt"

var scoremap map[string]int

func main() {
	//	var testmap = map[string]int{"a": 0, "b": 99, "c": 100, "d": 99, "e": 99}
	putOutMap := putInMessage()

	averscore := averagemap(putOutMap)
	fmt.Println("平均分是：", averscore)
	rankscore(putOutMap)

}

func putInMessage() (m map[string]int) {
	var peoplevalue int
	fmt.Print("请输入要录入的人数：")
	fmt.Scanln(&peoplevalue)

	scoremap = make(map[string]int)
	var name string
	var score int
	fmt.Println("请输入“姓名 分数”：")
	for i := 0; i < peoplevalue; i++ {
		fmt.Scanf("%s%d\n", &name, &score)
		scoremap[name] = score
	}
	fmt.Println("分数表是：", scoremap)
	return scoremap
}

func averagemap(m map[string]int) (aver float64) {
	i := 0
	j := 1
	for _, v := range m {
		i += v
		aver = float64(i) / float64(j)
		j++
	}

	return aver
}

func rankscore(m map[string]int) {
	var keyslice []string
	var valslice []int

	for key, val := range m {
		keyslice = append(keyslice, key)
		valslice = append(valslice, val)
	}

	for i := 0; i < len(valslice)-1; i++ {
		for j := 0; j < len(valslice)-1; j++ {
			if valslice[j] < valslice[j+1] {
				valslice[j], valslice[j+1] = valslice[j+1], valslice[j]
				keyslice[j], keyslice[j+1] = keyslice[j+1], keyslice[j]
			}
		}
	}

	for k := 0; k < len(valslice); k++ {
		fmt.Printf("%s的成绩是 ：%d\n", keyslice[k], valslice[k])
	}
}
