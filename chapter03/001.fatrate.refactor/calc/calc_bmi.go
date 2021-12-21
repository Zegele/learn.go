package calculator

import (
	gobmi "github.com/armstrongli/go-bmi" //这个就能自动下载并mod好。我的怎么不行？就是因为没有tag么？？？
	//如果在mod中把这个github源文件replace成staging文件后，这里的名字看着是github，其实引用的是staging里的内容。
	//（有点像障眼法，好处是，程序不会因为github的源文件变动而无法运行。
	//坏处是：这样要找引用地点的时候会比较容易搞错，需要结合mod看，或注意看代码提示）
)

func CalcBMI(tall float64, weight float64) (bmi float64) {
	bmi, _ = gobmi.BMI(weight, tall)
	return
}
