package main

/*常量的声明与变量类似，只不过是使用const关键字
可以是字符、字符串、布尔值或数值
不能用 := 语法声明*/
import "fmt"

const Pi = 3.14

func main() {
	const World = "世界"
	fmt.Println("Hello", World)
	fmt.Println("Happy", Pi, "Day")

	const Truth = true
	fmt.Println("Go rules?", Truth)
}
