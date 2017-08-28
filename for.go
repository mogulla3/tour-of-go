package main

import "fmt"

// Goのループ機構はforのみ
// forはセミコロンで分割された、3つのコンポーネントから構成される
// 1. 初回のイテレーション前に実行される
// 2. イテレーションの前に毎回評価される
// 3. イテレーションの後に毎回実行される

func main() {
	sum := 0
	for i := 0; i < 10; i++ {
		sum += i
	}

	fmt.Println(sum) // => 45
}
