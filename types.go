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


		// bool型
		var t, f bool = true, false
		fmt.Println(t, f)


		// string型
		var s string = "hello go"
		fmt.Println(s)
		fmt.Printf("%T\n", s)

		// 複数行の文字列を表示
		fmt.Println(`test
		test
		test
		`)

		/*
				string型はbyte配列の集まり
						var s string = "Hello"
						fmt.Println(s[0]) -> 72  バイドで表示される

						先頭の文字列を取得したい場合は、string型に変換する
						fmt.Println(string(s[0])) -> H
		*/


		// 配列型（要素数を変更できない）
		var arr1 [3]int = [3]int{1, 2, 4}
		fmt.Println(arr1)
		// 型を確認
		fmt.Printf("%T\n", arr1)  // -> [3]int 要素数まで含めた型になる

		// 要素数を格納した要素の数に自動的に設定する
		arr2 := [...]string{"a", "b"}
		fmt.Println(arr2)

		// 要素を取り出す
		fmt.Println(arr2[1])  // -> b

		// 要素の更新
		arr2[1] = "c"
		fmt.Println(arr2)

		// 要素数の確認
		fmt.Println(len(arr2))  // -> 2
}