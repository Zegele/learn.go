//函数的不定长参数
package main

import "fmt"

func main() {
	bmis := []float64{1.2, 2.3, 3.3}
	avgBmi := calculateAvg(bmis...) //给不定长传参， bmis是切片，加 ... 是将切片展开，后给函数传入参数
	fmt.Println(avgBmi)
	avgBmi = calculateAvg(1, 2, 3, 4, 5, 6, 7, 8)
	fmt.Println(avgBmi)
	avgBmi = calculateAvgOnSlice(bmis)
	fmt.Println(avgBmi)

	fmt.Println(getScoresOfStudent("Tom"))

}

func calculateAvg(bmis ...float64) (avgBmi float64) { //不定长参数
	var total float64 = 0
	for _, item := range bmis {
		total += item
	}
	avgBmi = total / float64((len(bmis)))
	return
}

func calculateAvgOnSlice(bmis []float64) float64 {
	var total float64 = 0
	for _, item := range bmis {
		total += item
	}
	return total / float64((len(bmis)))
}

func getScoresOfStudent(stdName string) (chinese int, math int, english int, physics int, nature int) {
	return 0, 0, 0, 0, 0
}
