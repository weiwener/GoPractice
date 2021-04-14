package main

import (
	"fmt"
	"math"
)

//if语句可以在条件表达式前执行一个简单的语句，该语句的变量作用域仅在id之内
//在if的简短语句中声明的变量同样可以在任何对应的else块中使用
func pow(x, n, lim float64) float64 {
	if v := math.Pow(x, n); v < lim {
		return v
	} else {
		fmt.Printf("%g >= %g\n", v, lim)
	}
	return lim
}

func main() {
	fmt.Println(pow(3, 2, 10), pow(3, 3, 20))
}
