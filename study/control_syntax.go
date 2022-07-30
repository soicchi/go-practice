package main

import "fmt"

func main() {
		// 条件分岐
		a := 1
		if a == 1 {
			fmt.Println("True")
		} else {
			fmt.Print("False")
		}


		// for文
		sampleList := []string{"sample1", "sample2", "sample3"}
		for _, v := range sampleList {
				fmt.Println(v)
		}

		i := 0
		for i < 5 {
				fmt.Println(i)
				i++
		}
}