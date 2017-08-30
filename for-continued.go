package main

import "fmt"

func main() {
	sum := 1

	// 最初と最後の条件は指定しなくてもよい
	for sum < 1000 {
		sum += sum
	}
	fmt.Println(sum)
}
