package main

import "fmt"

//for的初始化语句和后置语句是可选的，但是终止条件应该是得有的吧

func main() {
	sum := 1
	for sum < 20 {
		sum += sum
		fmt.Println(sum)
	}
}
