package main

import "fmt"

// 引数の型が同じ場合、省略記法が使える
func add(x, y int) int {
	return x + y
}

func main() {
	fmt.Println(add(1, 2))
}
