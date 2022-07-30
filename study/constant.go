package main

import "fmt"

// 定数の先頭を大文字にするとグローバルに呼び出し可能になる
const Pi = 3.14

// 複数定義
const (
		a = iota  // iotaは後続する定数に0から始まる連続する数値を入れてくれる
		b
		c
)

func main() {
		fmt.Println(Pi)
		fmt.Println(a, b, c)  // -> 0 1 2
}