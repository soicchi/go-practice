package main

import "fmt"

// 関数外では暗黙的な定義はできない
// num5 := 500

// 明示的な定義は可能
var num5 int = 500
// 関数外で関数は呼び出せない
// fmt.Println(num5)

func main() {
		var num int = 100
		fmt.Println(num)

		var str string = "hello go"
		fmt.Println(str)

		// 同じ型を複数定義する
		var good, bad bool = true, false
		fmt.Println(good, bad)

		// 異なる型をまとめて定義する
		var (
				num2 = 100
				str2 = "sample"
		)
		fmt.Println(num2, str2)

		// 何も値を入れない場合
		var (
				num3 int
				str3 string
		)
		fmt.Println(num3, str3)

		// 再定義可能
		num3 = 300
		str3 = "Go"
		fmt.Println(num3, str3)

		// 暗黙的な定義（型を定義しない）
		num4 := 400  // 代入された値から暗黙的にint型になる
		fmt.Println(num4)

		// string型を入れるとエラーになる
		// num4 = "string"
		// fmt.Println(num4)

		// 再度暗黙的な定義はできない
		// num4 := 500
		// fmt.Println(num4)
}