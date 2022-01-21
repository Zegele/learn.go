package main

import "fmt"

func main() {
	frSvc := &fatRateService{ //fatRateService是结构体 可以实现一整套
		s:     GetFatRateSuggestion(), // suggest 是结构体。是GetFatRateSuggestion()的返回值，返回值是个结构体，该结构体里是个多维的切片。
		input: &inputFromStd{},        // 该结构体，用input接口实现
		//看到结构体时，看它的上级，搞清楚，这里是结构体，还是用接口实现了结构体。
		output: &StdOut{}, // 该结构体，用output接口实现
		// 这里的output 相比其他结构体，多了层接口的含义。理解上按结构体理解是没有问题的。
		//input: &fakeInput{},
		//output:&fakeOutput{},
	} //todo

	for {
		p := frSvc.input.GetInput()
		frSvc.GiveSuggestionToPerson(&p)
		fmt.Println((*frSvc).output) //&{{小强 男 1.8 65 30 0.2006172839506173 0.14474074074074073} 标准}
		//fmt.Println(frSvc.output)//效果同上
	}
}
