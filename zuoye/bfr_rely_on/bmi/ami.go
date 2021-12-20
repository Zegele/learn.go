package bmi //文件名和包名不一样

func Calcbmi(weight, tall float64) (bmi float64) {
	bmi = weight / (tall * tall) / 100
	return
}
