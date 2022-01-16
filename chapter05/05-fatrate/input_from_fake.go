package main

type fakeInput struct {
}

func (*fakeInput) GetInput() Person {
	return Person{
		name:   "小强",
		sex:    "男",
		tall:   1.8,
		weight: 65,
		age:    30,
	}
}
