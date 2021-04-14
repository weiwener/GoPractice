package main

import "fmt"

//for是go中的 while,去掉分号就是了

func main() {
	sum := 1
	for sum < 20 {
		sum += sum
		fmt.Println(sum)
	}
}
