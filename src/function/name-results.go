package main

//go的返回值可被命名，会被视作在函数顶部的变量；
//没有参数的return语句返回已命名的返回值，也就是直接返回,为避免影响代码的可读性尽量用在短函数中
import "fmt"

func split(sum int) (x, y int) {
	x = sum * 4 / 9
	y = sum - x
	return
}

func main() {
	fmt.Println(split(8))
}
