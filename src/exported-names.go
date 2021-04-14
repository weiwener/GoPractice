package main

import (
	"fmt"
	"math"
)

//如果一个名字以大写字母开头，那么就是已导出的。在导入一个包时，只能引用其中已导出的名字
func main() {
	fmt.Println(math.Pi)
}
