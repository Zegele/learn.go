package main

// 一般接口，会专门放在interface的go文件中。
type PutElephantIntoRefrigerator interface {
	OpenTheDoorOfRefrigerator(Refrigerator) error
	PutElephantIntoRefrigerator(Elephant, Refrigerator) error
	CloseTheDoorOfRefrigerator(Refrigerator) error
}

type Refrigerator struct {
	Size string
}

type Elephant struct {
	Name string
}
