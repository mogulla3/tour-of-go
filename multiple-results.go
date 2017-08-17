package main

import "fmt"

// 関数は複数の返り値を返すことが可能
func swap(x, y string) (string, string) {
	return y, x
}

func main() {
	a, b := swap("hello", "world")
	fmt.Println(a, b) // => "world hello"
}
