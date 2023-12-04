// bytes中Reader和Buffer两个结构的使用
// strings.NewReader呢？
package main

import (
	"bytes"
	"fmt"
)

func useReader() {
	data := "abcdefghijk"
	//通过[]byte创建Reader
	re := bytes.NewReader([]byte(data)) // 只是为data创建了Reader，还没有开始读取数据
	//返回未读取部分的长度
	fmt.Println("re len : ", re.Len()) //re len :  11

	//返回底层数据总长度
	fmt.Println("re size : ", re.Size()) //re size :  11

	buf := make([]byte, 2)
	for {
		//读取数据
		n, err := re.Read(buf) // 把data中的数据，读到buf中。
		if err != nil {
			break
		}
		fmt.Println(string(buf[:n]))
		//返回未读取部分的长度
		fmt.Println("for 循环中：re len : ", re.Len()) //re len :  11
	}

	//设置偏移量，因为上面的操作已经修改了读取位置等信息
	//fmt.Println("设置offset")
	//re.Seek(10, 0) //第一个参数是后移几个字节，whence是三个数：0，1，2。 0代表从头开始，1表示当前位置，2表示光标在最后
	// 如果offset是3， whence0， 就是从开头数，然后后移3个字节，再开始
	// 如果offset是3， whence1， 就是从光标当前位置数，然后后移3个字节，再开始。
	// 如果offset是3， whence2， 就是从最后的位置数，然后后移3个字节，再开始。
	//fmt.Println("whence")
	re.Seek(0, 0)
	for {
		//一个字节一个字节的读
		b, err := re.ReadByte()
		if err != nil {
			break
		}
		fmt.Println(string(b))
	}

	re.Seek(0, 0)
	fmt.Println("off是什么？")
	off := int64(3)
	for {
		//指定偏移量读取
		n, err := re.ReadAt(buf, off) // 如果off设置为3，就是后移3个字节，再开始读取内容。
		// 且off的长度是按总长算的。所以后面off通过for循环累加，就继续读取后移之后的内容。
		if err != nil {
			break
		}
		off += int64(n)
		fmt.Println(off, string(buf[:n]))
	}
}

func useBuffer() {
	data := "123456789"
	//通过[]byte创建一个Buffer
	bf := bytes.NewBuffer([]byte(data))

	//Len()返回未读取的数据长度
	fmt.Println("bf len : ", bf.Len())

	//Cap()缓存容量
	fmt.Println("bf cap : ", bf.Cap()) //cap是8个8个增加的

	//Bytes()返回未读取的数据切片
	bys := bf.Bytes()
	for _, v := range bys {
		//fmt.Println(v)
		fmt.Print(string(v) + " ")
	}
	fmt.Println()

	//Next() （返回）读取未读取部分前n个字节数据的切片
	for i := 0; i < 10; i++ {
		tmp := bf.Next(2) // 相当于已经读取了，因为循环后光标后移了。
		fmt.Print("??", string(tmp)+" ")
	}
	fmt.Println()
	//再次Next，返回[]byte，说明没有未读取的 ，所以就是读取完了。
	fmt.Println(bf.Next(1))

	//重设缓冲，丢弃全部内容
	bf.Reset() // 没有任何东西了。

	//通过string创建Buffer
	bf2 := bytes.NewBufferString(data)
	//读取第一个 delim 及其之前的内容，返回遇到的错误
	line, _ := bf2.ReadBytes('0') //读取到第一个是3的数据为止，并返回 字节类型 。如果没有找到该元素，就已经读完了，那就返回全部数据。
	fmt.Println("===", string(line))

	//效果同上，返回string
	line2, _ := bf2.ReadString('7') // 效果同上，但返回的是字符串类型
	fmt.Println("++++", line2)      //有没有办法让光标回到最开头再开始读？类似 re.Seek(0, 0)的方法

	//创建一个空Buffer
	bf3 := bytes.Buffer{}
	//自动增加缓存容量，保证有n字节剩余空间
	bf3.Grow(16)
	//写入rune编码，返回写入的字节数和错误。
	n, _ := bf3.WriteRune(rune('中')) // 返回的n是写入了n个字节数，内容已经写到bf3中了。问题：那怎么把buffer中的内容放出来？
	n1, _ := bf3.WriteRune(int32('中'))

	fmt.Println("bf3 write ", n)
	fmt.Println("bf3 write ", n1)

	n, _ = bf3.WriteString("国人")
	fmt.Println("bf3 write ", n)
	//返回未读取的字符串
	fmt.Println(bf3.String())
	//将数据长度截断到n字节
	bf3.Truncate(6)
	fmt.Println(bf3.String())

}

func main() {
	//防止main中代码过多，我新建两个函数单独写
	//useReader()
	useBuffer()
}
