// go语言文件读写
// www.kancloud.cn/imdszxs/golang/1509743
package main

import (
	"os"
)

//读文件
//在go语言中，文件是使用指向os.File类型的指针来表示的，也叫做文件句柄
//之前的标准输入os.Stdin，和标准输出os.Stdout都是*os.File类型/
/*
func main() {
	inputFile, inputError := os.Open("E:/Geek/src/learn.go/buchong/shili_wenjianduxie/cookies.dat")
	if inputError != nil {
		fmt.Printf("打开文件时出错", inputError.Error())
		return // 退出函数
	}
	defer inputFile.Close()
	inputReader := bufio.NewReader(inputFile)
	i := 0
	for {
		inputString, readerError := inputReader.ReadString('\n')
		if readerError == io.EOF {
			return
		}
		i++
		fmt.Printf("第 %v 行：%s", i, inputString)
	}
}

//变量inputFile是*os.File类型的
//该类型是一个结构，表示一个打开文件的描述符（文件句柄）
//使用os包里的Open函数函数打开一个文件，
//该函数的参数是文件名，类型为string
//在上面的程序中，我们只读模式打开cookies.dat文件
//如果文件不存在或程序没有足够权限打开这个文件，Open函数会返回一个错误
//如果文件打开正常，就使用defer.Close()语句确保在程序退出前关闭该文件
//然后，使用bufio.NewReader来获得一个读取器变量
//通过使用bufio包提供的读取器（写入器也类似），可以很方便的操作相对高层的string对象，而避免了去操作比较底层的字节
//接着，我们在一个无限循环中使用ReadString('\n')或ReadBytes('\n')将文件的内容逐行（行结束符'\n'）读取出来
//注意：在之前的例子中，Unix、Linux的运行结束符是\n，
//而Windows的行结束符是\r\n
//在使用ReadString和ReadBytes方法的时候，我们不需要关心操作系统的类型
//直接使用\n就可以了。
//另外，我们也可以使用ReadLine()方法来实现相同的功能

//一旦读取到文件末尾，变量readerError的值将变成非空（事实上，常量io.EOF的值是true）
//我们就会执行return语句从而退出循环

*/

//其他类似函数
/*
//1. 将整个文件的内容读到一个字符串里
//可以使用，io/ioutil包里的ioutil.ReadFile()方法，
//该方法第一个返回值的类型是[]byte，里面存放读取到的内容，
//第二个返回值是错误，如果没有错误发生，返回值是nil
//示例：使用函数WriteFile()将[]byte的值写入文件

func main() {
	inputFile := "E:/Geek/src/learn.go/buchong/shili_wenjianduxie/cookies.dat"
	outputFile := "E:/Geek/src/learn.go/buchong/shili_wenjianduxie/copy.dat"
	buf, err := ioutil.ReadFile(inputFile)
	if err != nil {
		fmt.Fprintf(os.Stderr, "File Error: %s\n", err)
		//panic(err.Error)
	}
	fmt.Printf("%s\n", string(buf))
	err = ioutil.WriteFile(outputFile, buf, 0x644)
	if err != nil {
		panic(err.Error())
	}
}

*/

/*
//2.带缓冲的读取
//在很多情况下，文件的内容是不按行划分的
//或者干脆就是一个二进制文件
//在这种情况下，ReadString()就无法使用了，
//我们可以使用bufio.Reader的Read()
//它只接收一个参数：、

func main() {
	inputFile, inputError := os.Open("E:/Geek/src/learn.go/buchong/shili_wenjianduxie/cookies.dat")
	if inputError != nil {
		fmt.Printf("打开文件时出错", inputError.Error())
		return // 退出函数
	}
	defer inputFile.Close()

	inputReader := bufio.NewReader(inputFile)
	buf := make([]byte, 1024)
	for {
		n, err := inputReader.Read(buf)
// n是读取到的字节数
		fmt.Println(string(buf))
		if err != nil {
			fmt.Println(err)
		}
		if n == 0 {
			break
		}
	}
}
*/

/*
//3.按列读取文件中的数据
//如果数据是按列排列并用空格分隔的
//可以使用fmt包提供的以FScan开头的一系列函数来读取他们
//示例：将3列的数据分别读入变量v1，v2，和v3内
//然后分别把他们添加到切片的尾部
//、
func main() {
	file, err := os.Open("E:/Geek/src/learn.go/buchong/shili_wenjianduxie/cookies.dat")
	// 该文件必须是3列，且用空格隔开(tab也可)
	if err != nil {
		panic(err)
	}
	defer file.Close()

	var col1, col2, col3 []string
	for {
		var v1, v2, v3 string
		_, err = fmt.Fscanln(file, &v1, &v2, &v3)
		//scans until newline
		if err != nil {
			break
		}
		col1 = append(col1, v1)
		col2 = append(col2, v2)
		col3 = append(col3, v3)
	}
	fmt.Println(col1)
	fmt.Println(col2)
	fmt.Println(col3)
}

// 注意：path包里含一个子包叫pilepath，这个子包提供了跨平台的函数
//用于处理文件名和路径，例如Base()函数用于获得路径中的最后一个元素（不包含后面的分隔符）：
//import "path/filepath"
//filename := filepath.Base(path)

*/

/*
//4.compress包：读取压缩文件
// compress包提供了读取压缩文件的功能
//支持的压缩二五年间格式为：bzip2,flate, gzip, lzw 和 zlib。
// 示例： 使用go语言读取一个gzip文件
func main() {
	fName := "MyFile.gz"
	var r *bufio.Reader
	fi, err := os.Open(fName)
	if err != nil {
		fmt.Fprintf(os.Stderr, "%v, Can't open %s: error: %s\n", os.Args[0], fName, err)
		os.Exit(1)
	}
	fz, err := gzip.NewReader(fi) //"compress/gzip"包
	if err != nil {
		r = bufio.NewReader(fi)
	} else {
		r = bufio.NewReader(fz)
	}

	for {
		line, err := r.ReadString('\n')
		if err != nil {
			fmt.Println("Done reading file")
			os.Exit(0)
		}
		fmt.Println(line)
	}
}

*/

// 5. 写文件
/*
func main() {
	// var outputWriter *bufio.Writer
	// var outputFIle *os.File
	// var outputError os.Error
	// var outputString string
	outputFile, outputError := os.OpenFile("output.dat", os.O_WRONLY|os.O_CREATE, 0666)
	if outputError != nil {
		fmt.Printf("An error occurred with file opening or creation\n")
		return
	}
	defer outputFile.Close()
	outputWriter := bufio.NewWriter(outputFile)
	outputString := "golang\n"
	for i := 0; i < 10; i++ {
		outputWriter.WriteString(outputString)
	}
	outputWriter.Flush()
}


*/
//除了文件句柄，我们还需要bufio的写入器
//我们以只读模式打开文件output.dat
//如果文件不存在则自动创建 os.O_WRONLY|os.O_CREATE
//OpenFile函数有三个参数：
//文件名，
//一个或多个标志（使用逻辑运算符“|”连接）
//使用的文件权限
//os.O_RDONLY ：只读
//os.O_WRONLY ：只写
//os.O_CREATE ：创建：如果指定文件不存在，就创建该文件
//os.O_TRUNC ： 截断：如果指定文件已存在，就将该文件的长度截为0
//
//在读文件的时候，文件的权限是被忽略的，所以在使用OpenFile时传入的第三个参数可以用0
//而在写文件时，不管是Unix还是Windows，都需要使用0666
// 详细看小强老师的第8章
//然后，创建一个写入器（缓冲区）对象
//接着，使用for循环，将字符串写入缓冲区，写10次
//缓冲区的内容紧接着被完全写入文件：outputWriter.Flush()
//如果写入的东西很简单，我们可以使用
//fmt.Fprintf(outputFile, "some test data.\n")直接将内容写入文件
//fmt包里的F开头的print函数可以直接写入任何io.Writer ,包括文件
//
//示例：不使用fmt.Fprintf函数，使用其他函数如何写文件
func main() {
	os.Stdout.WriteString("hello, world\n") // 在运行输出界面出现hello, world
	f, _ := os.OpenFile("test", os.O_CREATE|os.O_WRONLY, 0)
	defer f.Close()
	f.WriteString("hello, golang in a file\n") //文件中写入这句

}

//使用os.Stdout.WriteString("xxxx"),可以输出到屏幕
//以只写模式创建或打开文件，并且忽略可能发生的错误
//不适用缓冲区，直接将内容写入文件：f.WriteString()
