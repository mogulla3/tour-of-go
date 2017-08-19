package main

import "fmt"

// 関数内では var のかわりに := という省略記法が使える
// := は型宣言を省略できる
// 関数の外では必ずキーワード（var, funcなど）から始まる必要があるため、:= は使えない
func main() {
	var i, j int = 1, 2
	k := 3

	c, python, java := true, false, "no!"
	fmt.Println(i, j, k, c, python, java) // => 1 2 3 true false no!
}
