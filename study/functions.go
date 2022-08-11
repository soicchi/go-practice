package main

import "fmt"

//	func Plus(x int, y int) int {
//			return x + y
//	}
//
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
func Double(price int) (result int) { // 返り値の型宣言で変数を宣言
	result = price * 2
	return // 宣言した変数を返す場合、返り値を省略
}

// 返り値を省略2
func Noreturn() {
	fmt.Println("No Return")
	return
}

// クロージャー
func addNum() func() int {
	x := 0
	return func() int {
		x++
		return x
	}
}

/*
	func main() {
		x := 0
		addNum := func() int {
				x++
				return x
		}
		fmt.Println(addNum()) -> 1
		fmt.Println(addNum()) -> 2
		fmt.Println(addNum()) -> 3
		fmt.Println(addNum()) -> 4
	}
	上記の定義だとaddNumに代入した関数外なのでxの値を書き換えられる可能性がある

	では変数xを関数内に移動させたらどうか？
	func main() {
		addNum := func() int {
				x := 0
				x++
				return x
		}
		fmt.Println(addNum()) -> 1
		fmt.Println(addNum()) -> 1
		fmt.Println(addNum()) -> 1
		fmt.Println(addNum()) -> 1
	}
	上記では呼出ごとにxが0で初期化されてしまうためうまく動作しない

	ということで関数を定義してその中で変数xを外側に持つ関数を定義すれば良い（上記参照）
*/

func main() {
	num := Plus(1, 2)
	fmt.Println(num)

	num2, num3 := Div(40, 3)
	fmt.Println(num2, num3)

	num4 := Double(1000)
	fmt.Println(num4)

	Noreturn()

	// 無名関数を変数へ代入
	f := func(x, y int) int {
		return x + y
	}
	fmt.Println(f(1, 2))

	// 返り値を持つ関数は一度変数に代入して使用
	f2 := addNum()
	fmt.Println(f2())
	fmt.Println(f2())
	fmt.Println(f2())
	fmt.Println(f2())
	fmt.Println(f2())
}
