package main

// 名前付きの戻り値を返すことができる

import "fmt"

// "x", "y"という名称の返り値を返す
func split(sum int) (x, y int) {
	x = sum * 2
	y = sum + x

	// "return"文に引数を付与しないと名前付きの戻り値を返す => これを"naked return"という
	// naked return は短い関数でのみ使うべし. 長いと読みづらくなる
	return
}

func main() {
	x, y := split(10)
	fmt.Println(x) // => 20
	fmt.Println(y) // => 20

	fmt.Println(split(10)) // => 20 30
}
