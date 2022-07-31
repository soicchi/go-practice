package main

import (
		"fmt"
		"time"
)

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


		// map
		var m = map[int]string{1: "JAPAN", 2: "USA"}
		fmt.Println(m)
		// 暗黙的宣言
		m2 := map[int]string{}
		fmt.Println(m2)
		// 要素追加
		m[3] = "CHINA"
		fmt.Println(m)
		// 要素の削除
		delete(m, 3)
		fmt.Println(m)
		// 第二引数を加えると取得できたかをtrue or falseで返す
		_, ok := m[3]
		if !ok {
				fmt.Println("Error")
		}
		fmt.Println(ok)


		// チャネル(複数のルーチン間でデータの送受信をするためのデータ構造)
		// 送受信可能
		var ch1 chan int

		// 受信専用
		// var ch2 <-chan int

		// 送信専用
		// var ch3 chan<- int

		// チャネル生成
		ch1 = make(chan int)

		// 一気にチャネル生成
		ch2 := make(chan int)
		fmt.Println(cap(ch1))
		fmt.Println(cap(ch2))

		// バッファサイズを設定する(デフォルトでは0)
		ch3 := make(chan int, 5)  // 第二引数に設定、設定しないとデータを入れることができない。
		fmt.Println(cap(ch3))

		// チャネルにデータ送信
		ch3 <- 1
		ch3 <- 2
		fmt.Println(len(ch3))

		// チャネルからデータを受信
		sampleData := <-ch3
		fmt.Println(sampleData)
		fmt.Println(len(ch3))
		// チャネルのデータはキュー構造になっている。つまり最初に入れたデータから取り出される


		// チャネルとゴルーチン
		ch4 := make(chan int)
		ch5 := make(chan int)

		reciever := func(c chan int) {
				for {
						i := <-c
						fmt.Println(i)
				}
		}

		go reciever(ch4)
		go reciever(ch5)

		i := 0
		for i < 100 {
				ch4 <- i
				ch5 <- i
				time.Sleep(50 * time.Millisecond)
				i++
		}



		ch6 := make(chan int, 2)

		/*
				ch6 <- 2

				チャネルをクローズする
				close(ch6)

				クローズされたチャネルにはデータを送信できない
				ch6 <- 3

				ch_num, ok := <-ch6
				fmt.Println
		*/

		reciever2 := func(name string, ch chan int) {
				for {
						i, recieved := <-ch
						if !recieved {
								break
						}
						fmt.Println(name, i)
				}
				fmt.Println(name, "END")
		}

		go reciever2("tanaka", ch6)
		go reciever2("sato", ch6)

		i2 := 0
		for i2 < 100 {
				ch6 <- i2
				i2++
		}
		close(ch6)
		time.Sleep(3 * time.Second)


		// チャネルでfor文を使用する場合
		ch7 := make(chan int, 3)
		ch7 <- 1
		ch7 <- 2
		ch7 <- 3
		close(ch7)  // クローズしないとfor文でチャネルにデータがないのに呼び出しが行われdeadlockになる
		for ch7Val := range ch7 {
				fmt.Println(ch7Val)
		}
}