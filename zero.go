package main

import "fmt"

// 初期値が与えられなかった変数はゼロ値となる
// ゼロ値
// - 数値型の場合 => 0
// - 真偽値型の場合 => false
// - 文字列型の場合 => ""
func main() {
	var i int
	var f float64
	var b bool
	var s string

	fmt.Printf("%v %v %v %q\n", i, f, b, s) // => 0 0 false ""
}
