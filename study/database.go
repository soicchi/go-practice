package main

import (
	"database/sql"
	"log"

	// アンダースコアはimportしたものをコードで使用しない場合に付ける
	_ "github.com/lib/pq"
)

func main() {
	db, _ := sql.Open("postgres", "host=db user=root dbname=test_db password=password sslmode=disable")
	defer db.Close()

	_, err := db.Exec("CREATE TABLE persons(name STRING, age INTEGER)")
	if err != nil {
		log.Fatalln(err)
	}
}