package main

import "fmt"

// func Plus(x int, y int) int {
// 		return x + y
// }
// 引数の型はまとめて定義できる
func Plus(x, y int) int {
		return x + y
}

// 返り値が複数の場合
func Div(x, y int) (int, int) {
		q := x / y
		r := x % y
		return q, r
}

// 返り値を省略1
func Double(price int) (result int) {  // 返り値の型宣言で変数を宣言
		result = price * 2
		return  // 宣言した変数を返す場合、返り値を省略
}

// 返り値を省略2
func Noreturn() {
		fmt.Println("No Return")
		return
}

func main() {
		num := Plus(1, 2)
		fmt.Println(num)

		num2, num3 := Div(40, 3)
		fmt.Println(num2, num3)

		num4 := Double(1000)
		fmt.Println(num4)

		Noreturn()


		// 無名関数
		f := func(x, y int) (int) {
				return x + y
		}
		fmt.Println(f(1, 2))
}