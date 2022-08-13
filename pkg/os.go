package main

import (
	"fmt"
	"log"
	"os"
)

func main() {
	// Exitで強制的に終了する
	// os.Exit(1)
	// fmt.Println("Start")

	/*
		deferで定義されている関数等は実行されない。
		【実行順序】
		os.Exit(1)が実行され処理は終了 → deferで定義された処理
		→ deferを実行する前に処理が終了されている。
	*/

	// ファイルを開く
	// _, err := os.Open("sample.txt")
	// if err != nil {
	// 	log.Fatalln(err)  // Fatallnは引数のログを出力し、Exitで処理を終了させる
	// }

	// コマンドライン引数を受け取る
	// fmt.Println(os.Args[0])
	// fmt.Println(os.Args[1])
	// fmt.Println(os.Args[2])
	// fmt.Println(os.Args[3])

	// ファイル操作
	// f, err := os.Open("test.txt")
	// if err != nil {
	// 	log.Fatalln(err)
	// }
	// defer f.Close()

	// ファイル作成
	// os.Createで指定したファイル名が既に存在する場合は、既存ファイルを削除し新規作成する
	// f, _ := os.Create("test2.txt")
	// f.WriteAt([]byte("Golang"), 4)
	// f.Seek(0, os.SEEK_END)
	// f.WriteString("テスト！！！！！")

	// ファイルの読み込み
	// f, err := os.Open("test.txt")
	// if err != nil {
	// 	log.Fatalln(err)
	// }
	// defer f.Close()

	// bs := make([]byte, 128)
	// n, _ := f.Read(bs)
	// fmt.Println(n)
	// fmt.Println(string(bs))

	// OpenFile
	/*
		フラグはor条件のように複数指定できる
		ex) f, err := os.OpenFile("test.txt", os.O_RDWR|os.O_CREATE, 0666)
	*/
	f, err := os.OpenFile("test.txt", os.O_RDONLY, 0666)
	if err != nil {
		log.Fatalln(err)
	}
	defer f.Close()

	bs := make([]byte, 128)
	n, err := f.Read(bs)
	fmt.Println(n)
	fmt.Println(string(bs))
}