package main

import "fmt"

// 定数は const キーワードを使って定義する
// := を使って宣言はできない
const Pi = 3.14

func main() {
	const World = "world"
	fmt.Println("Hello", World)     // => Hello world
	fmt.Println("Happy", Pi, "Day") // => Happy 3.14 Day

	const Truth = true
	fmt.Println("Go rules?", Truth) // => Go rules? true
}
