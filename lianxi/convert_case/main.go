package main

import "fmt"

func main() {
	var slice1 = []string{"HELLO", "你好", "hello", "HellO", "！", "!"}
	fmt.Println(slice1)

	for i, vslice := range slice1 {
		tempByte := []byte(vslice)
		//fmt.Println(tempByte)
		for j, vbyte := range tempByte {
			if vbyte >= 65 && vbyte <= 90 { //65-90 : A-Z
				tempByte[j] = vbyte + 32
			} else if vbyte >= 97 && vbyte <= 122 { //97-122 :a-z
				tempByte[j] = vbyte - 32
			}
		}
		slice1[i] = string(tempByte)
	}
	fmt.Println(slice1)
}
