package main

import "fmt"

// ":="や"var ="など、型を明示しない宣言をした場合
// 変数の型は右辺の値から推測される
func main() {
	v := 42
	fmt.Printf("v is of type %T\n", v) // => v is of type int
}
