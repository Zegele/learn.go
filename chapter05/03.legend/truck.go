package main

import "fmt"

type truckLegend struct {
}

func (*truckLegend) OpenTheDoorOfRefrigerator(Refrigerator) error {
	fmt.Println("用 truck 做 OpenTheDoorOfRefrigerator")
	return nil
}

func (*truckLegend) PutElephantIntoRefrigerator(Elephant, Refrigerator) error {
	fmt.Println("用 truck 做 PutElephantIntoRefrigerator")
	return nil
}

func (*truckLegend) CloseTheDoorOfRefrigerator(Refrigerator) error {
	fmt.Println("用 truck 做 CloseTheDoorOfRefrigerator")
	return nil
}
