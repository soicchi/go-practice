package main

import "fmt"

func main() {
		// スライス(配列と異なるのは要素数を宣言しない)
		var sl []int = []int{1, 2, 3, 4, 5}
		fmt.Println(sl)

		sl2 := []string{"Apple", "Google", "Amazon", "Meta"}
		fmt.Println(sl2)

		// 最後の要素を取る
		fmt.Println(sl[len(sl) - 1])  // -> 5

		fmt.Println(sl[2])
		fmt.Println(sl[1:])


		// append
		sl3 := []int{1, 2, 3}
		sl3 = append(sl3, 4)
		fmt.Println(sl3)

		// append(複数入れることもできる)
		sl3 = append(sl3, 5, 6, 7)
		fmt.Println(sl3)

		// make(指定した要素数分の初期値をスライスに格納)
		sl4 := make([]int, 5)
		fmt.Println(sl4)

		// len
		sl5 := []int{1, 2}
		fmt.Println(len(sl5))

		// cap(スライスの容量を確認する)
		sl6 := []int{1, 6, 9}
		fmt.Println(cap(sl6))


		// copy
		sl7 := []int{9, 8, 7, 6}
		sl8 := make([]int, 4)
		fmt.Println(sl8)
		copy(sl8, sl7)
		fmt.Println(sl8)


		// 可変長引数
		sampleFunc := func(s ...int) {  // int型の引数をいくらでも受ける
				for _, v := range s {
						fmt.Println(v)
				}
		}
		sampleFunc(1, 2, 3, 4, 5, 6, 7)
		sl9 := []int{1, 2, 3}
		sampleFunc(sl9...)  // スライスを展開して引数として渡す
}