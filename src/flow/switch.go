package main

import (
	"fmt"
	"runtime"
	"time"
)

/*go只运行选定的case,而非之后所有的case，实际上，go自动提供在每个case后面所需的break语句；
除非以fallthrough语句结束，否则分支自动终止；
注意：case无需为常量，且取值不必为整数*/

func main() {
	fmt.Println("Go runs on")
	switch os := runtime.GOOS; os {
	case "darwin":
		fmt.Println("OS X.")
	case "linux":
		fmt.Println("Linux.")
	default:
		fmt.Printf("%s \n", os)
	}

	//swtich的case语句从上到下顺次执行，直到匹配成功时停止
	fmt.Println("When's Saturday?")
	today := time.Now().Weekday()
	switch time.Saturday {
	case today + 0:
		fmt.Println("Today.")
	case today + 2:
		fmt.Println("In tow dayas.")
	default:
		fmt.Println("Too far away.")
	}

	//没有条件的switch同switch true一样，这种形式比if-then-else写的更加清晰
	t := time.Now()
	switch {
	case t.Hour() < 12:
		fmt.Println("Good morning!")
	case t.Hour() < 17:
		fmt.Println("Good afternoon.")
	default:
		fmt.Println("Good evening.")
	}
}
