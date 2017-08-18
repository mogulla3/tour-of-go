package main

import "fmt"

// `var`で変数の一覧を宣言
// 型は最後に記述
// package or functionレベルで記述できる
var c, python, java bool

func main() {
	var i int
	fmt.Println(i, c, python, java) // => 0 false false false
}
