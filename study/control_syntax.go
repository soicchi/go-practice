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

		// swicth分もあるが省略


		// 型アサーション
		var anything interface{} = 3
		result := anything.(int) + 2
		fmt.Println(result)
		/*
				変数 := i.(T)
						interface型に代入された値の型を復元してくれる
						ただし、代入している型に一致していないとpanicを引き起こす
						ex) var sample interface{} = 1
								sample.(float64)  -> もともとはint型なのでfloat型には復元できない
						変数を二つにすると2つ目の変数に復元できれば、true
						復元できなければ、falseになりpanicにならずに処理が進む
		*/

}