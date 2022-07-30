package main

import "fmt"

func main() {
		// 加算
		fmt.Println(1 + 2)
		// 減算
		fmt.Println(5 - 1)
		// 乗算
		fmt.Println(5 * 3)
		// 除算
		fmt.Println(60 / 3)
		// 余剰
		fmt.Println(34 % 5)
		// 文字列の結合
		fmt.Println("ABC" + "DFG")
		// 短縮系
		num := 100
		num += 20  // num = num + 20
		num++  // num + 1
		num--  // num - 1

		// 比較演算子
		fmt.Println(1 == 1)  // true
		fmt.Println(1 == 2)  // false
		fmt.Println(1 <= 2)  // true
		fmt.Println(1 >= 2)  // false
		fmt.Println(1 > 2)  // false
		fmt.Println(1 < 2)  // true
}