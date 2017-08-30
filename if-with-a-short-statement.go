package main

import (
	"fmt"
	"math"
)

func pow(x, n, lim float64) float64 {
	// if文は条件式の判定の前に短い文を記述できる
	// ここで宣言された変数はif文内のスコープしか持たない
	if v := math.Pow(x, n); v < lim {
		return v
	}
	return lim
}

func main() {
	fmt.Println(
		pow(3, 2, 10), // 9
		pow(3, 3, 20), // 20
	)
}
