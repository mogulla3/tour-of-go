package main

// Go's basic types
// - bool
// - string
// - int int8 int16 int32 int64
// - uint uint8 uint16 uint32 uint64 uintptr
// - bytes (alias for uint8)
// - rune (alias for int32)
// - float32 float64
// - complex64 complex128

// int, uint, uintptrは32bitシステムであれば32bitとなり、64bitシステムであれば64bitになる
// サイズや符号なし整数を使いたい特別な理由がない限り、整数を使う場合はintを使うべき

import (
	"fmt"
	"math/cmplx"
)

// importのようにまとめて宣言もできる
var (
	ToBe   bool       = false
	MaxInt uint64     = 1<<64 - 1
	z      complex128 = cmplx.Sqrt(-5 + 12i)
)

func main() {
	fmt.Printf("Type: %T Value: %v\n", ToBe, ToBe)     // => Type: bool Value: false
	fmt.Printf("Type: %T value: %v\n", MaxInt, MaxInt) // => Type: uint64 value: 18446744073709551615
	fmt.Printf("Type: %T value: %v\n", z, z)           // => Type: complex128 value: (2+3i)
}
