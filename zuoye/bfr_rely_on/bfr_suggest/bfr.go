package bfr1 //注意区分包名，文件名，文件夹名

func Calcbfr(bmi float64, age int, sexval float64) (bfr float64) {
	bfr = (1.2*(bmi*100) + 0.23*float64(age) - 5.4 - 10.8*(sexval)) / 100
	return
}
