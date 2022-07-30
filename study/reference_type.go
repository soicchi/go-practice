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
}