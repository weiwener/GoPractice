package main

import "fmt"

//go的for语句后面的三个构成部分外没有小括号，大括号是必须的
func main() {
	sum := 0
	for i := 0; i < 10; i++ {
		sum += i
		fmt.Println(sum)
	}
}
