// Go语言自定义数据文件
// c.biancheng.net/view/4543.html
package main

import (
	"fmt"
	"io"
	"time"
)

//对一个程序非常普遍的需求包括维护内部数据结构，以数据交换提供导入导出功能，也支持使用外部工具来处理数据。
//这章只关注文件处理，因此只关心如何从程序内部数据结构中读取数据，
//并将其写入标准和自定义格式的文件中

// 该程序接受两个文件名作为命令行参数，一个用于读，另一个用于写（它们必须是不同的文件）
//程序从第一个文件中读取数据（以其后缀表示的任何格式）
//并将数据写入第二个文件（也是以其后缀所表示的任何格式）
//
//由invoicedata程序创建的文件可跨平台使用，也就是说，无论是什么格式，
//Windows上创建的文件都可在Mac OS X 以及Linux上读取，反之亦然
//Gzip格式压缩的文件（如invoice.gob.gz）可以无缝读写
//
//这些数据由一个[]invoice组成，也就是说，是一个保存了指向Invoice值的指针的切片。
//每一个发票数据都保存在一个invoice类型的值中，同时每一个发票数据都以[]*Item的形式保存着0个或多个项

type Invoice struct {
	Id         int
	Customerld int
	Raised     time.Time
	Due        time.Time
	Paid       bool
	Note       string
	Items      []*Item
}

type Item struct {
	Id       string
	Price    float64
	Quantity int
	Note     string
}

//这两个结构体用于保存数据。下表给出了一些非正式的对比，展示了每种格式下相同的50000份随机发票数据所需的时间
// 以及以格式所存储文件的大小
//
//计时按秒计，并向上舍入到最近的十分之一秒。
//我们应该把计时结果认为是无绝对单位的，因为不同硬件以及不同负载情况下该值都不尽相同
//大小一栏以千字节（KB）算，该值在所有机器上应该都是相同的
//
//对于该数据集，虽然未压缩文件的大小千差万别，但压缩文件的大小都惊人的相似。
//而代码的函数不包括 所有格式通用的代码（如：哪些用于压缩和解压以及定义结构体的代码）
//表：各种格式的速度以及大小对比
//
//这些读写时间和文件大小在我们的合理预期范围内，除了纯文本格式的读写异常快之外。
//这得益于fmt包优秀的打印和扫描函数，以及我们设计的易于解析的自定义文本格式。
//
//对于JSON和XML格式，我们只简单地存储了日期部分而非存储默认的time.Time值（一个ISO-8601日期/时间字符串）
//通过牺牲一些速度和增加一些额外代码稍微减小了文件的大小

//例如:如果让JSON代码来处理time.Time值，它能够运行得更快，并且其代码行数与Go语言二进制编码差不多
//
//对于二进制数据，Go语言的二进制格式是最便于使用的，它非常快且极端紧凑，所需的代码非常少
//并且相对容易适应数据的变化
//然而，如果我们使用的自定义类型不原生支持gob编码，我们必须让该类型满足gob.Encoder和gob.Decoder接口，
//这样会导致gob格式的读写相当得慢，并且文件大小也会膨胀
//
//对于可读的数据，XML可能是最好使用的格式，特别是作为一种数据交换格式时非常有用。
//与处理JSON格式相比，处理XML格式需要更多行代码。这是因为Go没有一个xml.Marshaler接口，
//也因为我们这里使用了并行的数据类型（XMLInvoice和XMLItem）来帮助映射XML数据和发票数据（invoice和Item）
//
//使用XML作为外部存储格式的应用程序可能不需要并行的数据类型或者也不需要invoicedata程序这样的转换，
//因此就有可能比invoicedata例子中所给出的更快，并且所需的代码也更少
//
//除了读写速度和问及爱你大小以及代码行数之外，还有另一个问题值得考虑：格式的稳健性。
//例如，如果我们为Invoice结构体和Item结构体添加了一个字段，那么就必须再改变文件的格式
//我们的代码实行读写新格式并继续支持读九个师的难以程度如何？
//如果我们为文件格式定义版本，这样的变化就很容易被适应，除了让JSON格式同时适应读写新旧格式稍微复杂一点之外。

// 除了Invoice和Item结构体之外，所有文件格式都共享以下常量：
const (
	fileType    = "INVOICE"    // 用于纯文本格式
	magicNumber = 0x125D       // 用于二进制格式
	fileVersion = 100          // 用于所有的格式
	dataFormat  = "2006-01-02" // 必须总是使用该时间
)

// magicNumber 用于唯一标记发票文件
// fileVersion 用于标记发票文件的版本
// 该标记便于之后修改程序来适应数据格式的改变
// dataFormat稍后介绍，它表示我们希望数据如何按照可读的格式进行格式化
// 同时，我们也创建了一个接口。
type InvoiceMarshaler interface {
	MarshalInvoices(writer io.Writer, invoices []*Invoice) error
}

type InvoiceUnmarshaler interface {
	UnmarshalInvoices(reader io.Reader) ([]*Invoice, error)
}

// 这样做的目的是统一的方式针对特定格式使用reader和writer
//例如，下列函数是invoicedata程序用来从也给打开的文件中读取发票数据的。

func readinvoices(reader io.Reader, suffix string) ([]*Invoice, error) {
	var unmarshaler InvoicesUnmarshaler
	switch suffix {
	case ".gobn":
		unmarshaler = GobMarshaler{}
	case "H.inv":
		unmarshaler = InvMarshaler{}
	case "f. jsn", H.jsonn: // 什么玩意儿
		unmarshaler = JSONMarshaler{}
	case ".txt":
		unmarshaler = TxtMarshaler{}
	case ".xml":
		unmarshaler = XMLMarshaler{}
	}
	if unmarshaler != nil {
		return unmarshaler.UnmarshalInvoices(reader)
	}
	return nil, fmt.Errorf("unrecognized input suffix: %s", suffix)
}

// 其中， reader是任何能够io.Reader接口的值，
//例如：一个打开的文件（其类型为*os.File） 一个gzip解码器（其类型为*gzip.Reader）
//或者一个string.Readero字符串suffix是问及爱你的后缀名（从.gz文件中解压之后）
//
//后面，会看到GobMarshaler和InvMarshaler等自定义的类型，他们提供了MarshmInvoices()和UnmershalInvoice()方法
//因此满足InvoicesMarshaler和InvoicesUnmarshaler接口。

func main() {

}
