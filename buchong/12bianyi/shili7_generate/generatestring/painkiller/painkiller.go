//go:generate stringer -type=Pill
package painkiller

type Pill int

const (
	Placebo Pill = iota
	Aspirin
	Ibuprofen
	Paracetamol
	Acetaminophen = Paracetamol
)

// 然后，在painkiller.go文件所在的目录下运行go generate 命令
//执行成功后没有任何提示信息，
//但会在当前目录生成一个pill_string.go文件，
//文件实现了我们需要的String()方法
//
//问题：这个stringer工具能干嘛？？？
