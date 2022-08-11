package main

import "fmt"

func main() {
	var num int = 100

	double := func(num int) {
		num = num * 2
	}

	double(num)
	fmt.Println(num)

	// ポインタを使用して値渡しを参照渡しにする
	var num2 *int = &num
	fmt.Println(*num2)
	*num2 = 300
	fmt.Println(num) // -> 300 numも300に変わっている。

}
