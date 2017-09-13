package main

import (
	"fmt"
	"math"
)

func pow(x, n, lim float64) float64 {
	// ifで定義された変数vはelse節でも参照できる
	if v := math.Pow(x, n); v < lim {
		return v
	} else {
		fmt.Printf("%g >= %g\n", v, lim)
	}

	// 変数vはここでは参照できない
	return lim
}

func main() {
	fmt.Println(pow(3, 2, 10), pow(3, 3, 20))
	// => 27 > 10
	// => 9 20
}
