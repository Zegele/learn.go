package main

// 参考：www.kancloud.cn/imdszxs/golang/1543822
import (
	"flag"
	"fmt"
)

var Input_pstrName = flag.String("name", "gerry", "input ur name")
var Input_piAge = flag.Int("age", 20, "input ur age")
var Input_flagvar int

func Init() {
	flag.IntVar(&Input_flagvar, "flagname", 1234, "hielp message for flagname")
}

func main() {
	Init()
	flag.Parse()

	// After parsing, the arguments after the flag are available as the slice flag.Args() or individually as falg.Arg(i).
	// The arguments are indexed from 0 through flag.NArg()-1
	// Args returns the non-flag command-line arguments
	// NArg is the number of arguments remaining after flags have been processed.
	fmt.Printf("args=%s, arg(i)=%s num=%d\n", flag.Args(), flag.Arg(1), flag.NArg())
	for i := 0; i != flag.NArg(); i++ {
		fmt.Printf("arg[%d]=%s\n", i, flag.Arg(i))
	}

	fmt.Println("name=", *Input_pstrName)
	fmt.Println("age=", *Input_piAge)
	fmt.Println("flagname=", Input_flagvar)
}
