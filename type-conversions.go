package main

import (
	"fmt"
	"math"
)

// T(v) => vをT型に変換する
// Cと違い、異なる型の項目間での割り当てには明示的な変換が必要になる
func main() {
	var x, y int = 3, 4
	var f float64 = math.Sqrt(float64(x*x + y*y))
	var z uint = uint(f)
	fmt.Println(x, y, z) // => 3 4 5
}
