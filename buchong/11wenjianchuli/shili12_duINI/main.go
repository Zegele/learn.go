// GO语言从INI配置文件中读取需要的值
// juejin.cn/post/7029272388776755213
// www.cdsy.xyz/computer/programme/golang/20210307/cd161507567110718.html
package main

import (
	"bufio"
	"fmt"
	"os"
	"strings"
)

// .ini文件是Initialization File的缩写，即初始化文件，
// 是windows的系统配置文件所采用的存储格式，统管windows的各项配置，
//一般用户就是用windows提供的各项图形化管理界面就可实现相同的配置了。
//但在某些情况下，还是要直接编辑ini才方便，一般只有很熟悉windows才能去直接编辑
//INI文件由多行文本组成，整个配置由[]拆分为多个“段”（section）
//每个段中又以=分割为“键”和“值”
//INI文件以';' 至于行首视为注释
//注释后将不会被处理和识别
//
//我们的目的是获取woner.ini文件中的CMCDLLNAME=mapi32.dll这行的mapi32.dll

// 根据文件名，段名，键名获取ini的值
func getValue(filename, expectSection, expectKey string) string {
	//1.读取文件
	//Go语言的OS包中提供了文件打开函数os.Open()，文件读取完成后需要及时关闭
	//否则文件会发送占用，系统无法释放缓存资源
	//打开文件
	file, err := os.Open(filename) //提供INI文件名，或及路径
	// 文件找不到，返回空
	if err != nil {
		return "-------------123"
	}
	defer file.Close()

	//2.读取行文本
	//INI文件的格式是由多行文本组成，因此需要构造一个循环
	//不断地读取INI文件的所有行
	//go语言总是将文件以二进制格式打开，通过不同的读取方式对二进制文件进行操作
	//go语言对二进制读取有专门的代码
	//bufio包即可方便地以比较常见的方式读取二进制文件。
	// 在函数结束时，关闭文件、
	//读取INI文件时，需要注意各种异常情况。本文中的空白符，就是经常容易忽略的部分
	//空白符在调试时完全不可见，需要打印出字符的ASCII码才能辨别
	//抛开各种异常情况拿到了每行的行文本linestr后，就可以方便地读取INI文件的段和键值了。

	// 使用读取器读取文件
	reader := bufio.NewReader(file) // 传入文件并构造一个读取器
	//b := make([]byte, 1024)  // var b []byte //没有容量
	//n, err := reader.Read(b) // []byte得有容量 ,读出来就把reader中的内容读空了
	//fmt.Println(n, b)
	//b, err = reader.ReadBytes('\n')
	//fmt.Println(b)

	// 当前读取的段的名字
	var sectionName string // 提前声明段的名字字符串，方便后面的段和键值读取
	for {                  //构建一个读取循环，不断地读取文件
		// 读取文件的一行
		linestr, err := reader.ReadString('\n')
		//reader.ReadString从文件中读取字符串，直到碰到\n，也就是行结束
		//整个函数返回读取到的行字符串（包括\n）和可能的读取错误err，例如文件读取完毕。
		if err != nil {
			break
		}
		// 切掉行的左右两边的空白字符
		linestr = strings.TrimSpace(linestr) //TrimSpace 切掉行左右的空白
		//每一行的文本可能会在标识符两笔那混杂有空格、回车符等不可见的空白字符，使用strings.TrimSpace()可以去掉这些空白字符
		// 忽略空行
		if linestr == "" { // 可能存在空行的情况，继续读取下一行，忽略空行
			continue
		}
		// 忽略注释
		if linestr[0] == ';' { // string的[0],返回的是byte类型的
			//当行首为 ; 号，表示这一行是注释行，忽略一整行的读取
			continue
		}

		//3.读取段
		//行字符串linestr已经取出了空白字符串
		//段的起止符又以[开头，以]结尾，因此可以直接判断行首和行尾的字符串
		//匹配段的起止符匹配时读取的是段
		//此时，段只是一个标识，而无任何内容，因此需要将段的名字取出保存在sectionName（已在前面的代码中定义了）中
		//待读取段后面的键值对时使用。

		// 行首和尾巴分别是方括号的，说明是段标记的起止符
		if linestr[0] == '[' && linestr[len(linestr)-1] == ']' {
			// 将段名取出
			sectionName = linestr[1 : len(linestr)-1]
			//fmt.Println("65467", sectionName)

			// 这个段是希望读取的
		} else if sectionName == expectSection {
			//4.读取键值
			//当前行不是段（但是是希望读取的段），是键值对

			//切开等号分割的键值对
			pair := strings.Split(linestr, "=")
			// 使用strings.Split() 对linerstr,进行切割，使用=分割
			//分割后返回字符串切片，
			fmt.Println(pair)
			//保证切开只有一个等号分割的键值对情况
			if len(pair) == 2 { // 只考虑分割出2个元素的情况。其他情况忽略？如没有=,或多个=的情况
				//去掉键的多余空白字符
				key := strings.TrimSpace(pair[0])
				// 是期望的键
				if key == expectKey { // 判断是否为期望的键
					// 返回去掉空白字符的值
					return strings.TrimSpace(pair[1]) //返回干净的值
				}
			}
		}
	}
	return "????"
}
func main() {
	//v := getValue("E:/Geek/src/learn.go/buchong/11wenjianchuli/shili12_duINI/woner.ini", "Mail", "CMCDLLNAME32")
	file := "E:/Geek/src/learn.go/buchong/11wenjianchuli/shili12_duINI/woner.ini"
	fmt.Println(getValue(file, "Mail", "CMC"))
	fmt.Println(getValue(file, "Mail", "CMCDLLNAME32"))
	fmt.Println(getValue(file, "Mail", "MAPIXVER"))
}

//getValue()函数
//本例并不是将整个INI文件读取保存后在获取需要的字段数据并返回，
//这里使用getValue()函数,每次从指定文件中找到需要的段（Section）及键（Key）对于的值
//
//getValue()函数的声明如下：
//func getValue(filename, expectSection, expectKey string) string
//参数说明：filename:INI文件的文件名； expectSection：期望读取的段； expectKey：期望读取段中的键
// getValue()函数的逻辑由4部分组成：即读取文件、读取行文本、读取段和读取键值组成
