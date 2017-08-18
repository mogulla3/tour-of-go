package main

import "fmt"

// 宣言と同時に初期化できる
var i, j int = 1, 2

func main() {
	// 宣言と同時に初期化される場合、型を省略できる
	var c, python, java = true, false, "no!"
	fmt.Println(i, j, c, python, java) // => 1 2 true false no!
}
