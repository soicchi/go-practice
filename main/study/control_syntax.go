package main

import (
		"fmt"
		"time"
)

// init(packageで指定している関数の実行前に実行される)
func init() {
		fmt.Println("init")
}


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

		// swicth文は省略


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


		// ラベル付きforは省略


		// defer(関数の処理が終了した時の処理を記載できる)
		sampleFunc := func() {
				fmt.Println("Finish!!!!")
		}

		defer sampleFunc()


		// 並行処理
		subRoop := func() {
				for {
						fmt.Println("sub roop")
						time.Sleep(100 * time.Millisecond)
				}
		}

		mainRoop := func() {
				for {
					fmt.Println("main roop")
					time.Sleep(200 * time.Millisecond)
				}
		}
		// goを実行する関数の前につけるとmainRoopと並行でsubRoopの処理も走る
		go subRoop()
		mainRoop()
}