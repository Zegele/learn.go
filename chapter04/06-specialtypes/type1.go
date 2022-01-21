package main

import (
	"crypto/rand"
	"fmt"
	"math/big"
)

type Math = int //起了别名
type English = int
type Chinese = int

func main() {
	var mathScore int = 100
	var mathScore2 Math = 100

	mathScore2 = mathScore
	fmt.Println(mathScore2)
	getScoresOfStudent("")

	vg := &voteGame{students: []*student{}} //学习指针的这种用法。
	for i := 0; i < 5; i++ {
		vg.students = append(vg.students, &student{name: fmt.Sprintf("%d", i)}) //Sprintf是什么？？？格式化输出后，返回一个字符串类型？
	}
	leader := vg.goRun()
	fmt.Println(leader)
	leader.Distribute()

	var stdXQ = student{name: "小强"}
	var ldXQ Leader = Leader(stdXQ) //ldXQ 是Leader类型（内其实就是student结构体），stdXQ是定义的student结构体。Leader(*stdXQ1)，将stdCXQ转化为Leader类型。
	fmt.Println(stdXQ)
	fmt.Println(ldXQ)
	ldXQ.Distribute()

	var stdXQ1 = &student{name: "小强"}
	var ldXQ1 Leader = Leader(*stdXQ1) //Leader(*stdXQ1)，将*stdCXQ转化为Leader类型。
	fmt.Println(stdXQ1)
	fmt.Println(ldXQ1)
	ldXQ1.Distribute()

	//bytesTest1 := []byte{}
	//var str1 string = string(bytesTest1)
	std1 := student{ //type student struct{}
		name:      "",
		agree:     0,
		disagress: 0,
	}
	l := Leader{ //type Leader student//Leader是student类型。而student是一个结构体类型。所以Leader的实际类型是结构体。
		name:      "",
		agree:     0,
		disagress: 0,
	}
	fmt.Println(std1, l)
}

func getScoresOfStudent(name string) (Math, Chinese, English) {
	//todo
	return 0, 0, 0
}

type voteGame struct {
	students []*student
}

func (g *voteGame) goRun() *Leader {
	for _, item := range g.students {
		RandomCrypto, _ := rand.Int(rand.Reader, big.NewInt(int64(len(g.students)-1))) //rand.Int是随机数？？？
		randInt := RandomCrypto.Int64()
		fmt.Println("to:", randInt)
		item.voteA(g.students[randInt])
	}

	maxScore := -1
	maxScoreIndex := -1
	for i, item := range g.students {
		if maxScore < item.agree {
			maxScore = item.agree
			maxScoreIndex = i
		}
	}
	if maxScoreIndex >= 0 { // 判断是否大于0，因为如果没有学生，那么index就是默认值-1.
		return (*Leader)(g.students[maxScoreIndex]) //转成Lerder指针类型
		//为什么这样写。貌似能理解，但是不熟悉。自己使用的时候会想到这样用么？会么？
	}
	return nil
}

// 使用嵌套对象定义（继承）方式来定义班长
// type Leader struct{
// student
// }

// 使用类型重定义
type Leader student

func (l *Leader) Distribute() {
	fmt.Println("发作业了！！！")
}

type FoooTestFuncRedefine []string //类型可以是任意的

func (f FoooTestFuncRedefine) test111() {

}

type student struct {
	name      string
	agree     int
	disagress int
}

func (std *student) voteA(target *student) {
	target.agree++
}

func (std *student) VoteD(target *student) {
	target.disagress++
}
