package main

/*表达式T(v)将值v转换为类型T
var i int = 42
var f float64 = float64(i)
或者i := 42
f := float64(i)
*/
import (
	"fmt"
	"math"
)

func main() {
	var x, y int = 3, 4
	var f float64 = math.Sqrt(float64(x*x + y*y))
	var z uint = uint(f)
	fmt.Println(x, y, z)
}
