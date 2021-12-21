module learn.go

go 1.17

require (
	github.com/armstrongli/go-bmi v0.0.1
	github.com/spf13/cobra v1.3.0
	learn.go.tools v0.0.0-00010101000000-000000000000
)

require (
	github.com/inconshreveable/mousetrap v1.0.0 // indirect
	github.com/spf13/pflag v1.0.5 // indirect
)

replace learn.go.tools => ../learn.go.tools //自己写在这里的。但是vendor后会自动生成上面那个v0.0.0  程序可以运行了。 这个replace和上面require自动生成的，都不可以不用。

// 依赖本地库的用replace？
// github上的用require？
replace learn.go/zuoye/bfr_rely_on/bmi => ./github.com/!zegele/learn.go@v0.0.0-20211220135454-4870e4b74f1a/zuoye/bfr_rely_on/bmi

replace github.com/armstrongli/go-bmi => ./staging/src/github.com/armstrongli/go-bmi
