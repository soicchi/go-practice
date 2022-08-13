package main

import (
	"log"
	"os"
)

func main() {
	// ログの出力先を標準出力(os.Stdout)に設定
	log.SetOutput(os.Stdout)

	// log.Print("log\n")  // -> 2022/08/13 10:33:14 log
	// log.Println("log2")  // -> 2022/08/13 10:33:14 log2
	// log.Printf("log%d\n", 3)  // -> 2022/08/13 10:33:14 log3

	// // ログを出力しExitする
	// log.Fatal("log\n")  // -> 2022/08/13 10:35:37 log exit status 1
	// log.Fatalln("log2")  // -> 上記のFatalで処理が終了するため出力なし
	// log.Fatalf("log\n", 3)  // -> 同様

	// // panicを発生させ、ログを出力し処理を処理を終了する
	// log.Panic("log\n")
	// log.Panicln("log2")
	// log.Panicf("log\n", 3)

	/*
		ログの出力先をファイルに変更
		f, err := os.Create("log_test.txt")
		if err != nil {
			log.Fatalln(err)
		}

		出力先をlog_test.txtに設定
		log.SetOutput(f)
		log.Println("logを出力")
	*/

	/*
		ログのフォーマットを変更

		// デフォルトのフォーマット指定
		log.SetFlags(log.LstdFlags)
		log.Println("log")

		// マイクロ秒を追加
		log.SetFlags(log.Ldate | log.Ltime | log.Lmicroseconds)
		log.Println("log2")

		// 時刻とファイル行を出力
		log.SetFlags(log.Ltime | log.Lshortfile)
		log.Println("log3")

		// ログにプレフィックスを追加
		log.SetFlags(log.LstdFlags)
		log.SetPrefix("[LOG]")
		log.Println("log4")
	*/
}