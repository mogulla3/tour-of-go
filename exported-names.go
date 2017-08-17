// 大文字で始まる => exportされる
// Pizza, Pi,

// 小文字で始まる => exportされない
// pizza, pi

// packageをimportしたとき、exportされものしか参照できない
// exportされていないものはパッケージ外部からアクセスできない

package main

import (
	"fmt"
	"math"
)

func main() {
	fmt.Println(math.Pi)
}
