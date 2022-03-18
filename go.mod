module learn.go

go 1.17

require (
	github.com/armstrongli/go-bmi v0.0.1
	github.com/gin-contrib/pprof v1.3.0
	github.com/gin-gonic/gin v1.7.7
	github.com/go-sql-driver/mysql v1.6.0
	github.com/spf13/cobra v1.3.0
	google.golang.org/protobuf v1.27.1
	gorm.io/driver/mysql v1.3.2
	gorm.io/gorm v1.23.1
	learn.go.tools v0.0.0-00010101000000-000000000000
)

require (
	github.com/gin-contrib/sse v0.1.0 // indirect
	github.com/go-playground/locales v0.13.0 // indirect
	github.com/go-playground/universal-translator v0.17.0 // indirect
	github.com/go-playground/validator/v10 v10.4.1 // indirect
	github.com/golang/protobuf v1.5.2 // indirect
	github.com/inconshreveable/mousetrap v1.0.0 // indirect
	github.com/jinzhu/inflection v1.0.0 // indirect
	github.com/jinzhu/now v1.1.4 // indirect
	github.com/json-iterator/go v1.1.12 // indirect
	github.com/leodido/go-urn v1.2.0 // indirect
	github.com/mattn/go-isatty v0.0.14 // indirect
	github.com/modern-go/concurrent v0.0.0-20180306012644-bacd9c7ef1dd // indirect
	github.com/modern-go/reflect2 v1.0.2 // indirect
	github.com/spf13/pflag v1.0.5 // indirect
	github.com/ugorji/go/codec v1.1.7 // indirect
	golang.org/x/crypto v0.0.0-20210817164053-32db794688a5 // indirect
	golang.org/x/sys v0.0.0-20211205182925-97ca703d548d // indirect
	gopkg.in/yaml.v2 v2.4.0 // indirect
)

replace learn.go.tools => ../learn.go.tools //自己写在这里的。但是vendor后会自动生成上面那个v0.0.0  程序可以运行了。 这个replace和上面require自动生成的，都不可以不用。

// 依赖本地库的用replace？
// github上的用require？
replace learn.go/zuoye/bfr_rely_on/bmi => ./github.com/!zegele/learn.go@v0.0.0-20211220135454-4870e4b74f1a/zuoye/bfr_rely_on/bmi

replace github.com/armstrongli/go-bmi => ./staging/src/github.com/armstrongli/go-bmi
