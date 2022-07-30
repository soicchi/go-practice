package main

import "fmt"

func main() {
		// int型
		var num int = 100

		/*
				数値型にも以下のように種類がある
						int8, int16, int32, int64

				intのように数字を指定しない場合は使用しているマシンスペックでint32やint64になる

				しかし、マシンスペックによって決定したint型と明示的に定義したint型では別の型扱い
				→ つまり型が違うので計算ができない
				ex) var num int = 100  マシンスペックにより32ビットのint型になったと仮定
						var num2 int32 = 200
						num + num2 -> invalid operation: num + num2 (mismatched types int and int32)
		*/

		// Printfで変数の型を確認することができる
		fmt.Printf("%T\n", num)

		/*
				fmt.Printf("%T, %T, ...", 変数, 変数2, ...)
						"%T" string: 各変数のTypeを表示
								末尾に改行を入れる場合は"%T\n"
						変数 any: 型を調べたい変数
		*/

		// 型の変換
		fmt.Println(int32(num))

		/*
				型(変数)
						型: int, floatなど
						変数: 型を変換したい変数
		*/

		// float型
		var num2 float64 = 2.4
		fmt.Println(num2)

		// 暗黙的な定義の場合はfloat64型になる
		num3 := 4.5
		fmt.Println(num3)
}