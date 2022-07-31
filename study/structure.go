package main

import "fmt"

type User struct {
		Name string
		Age int
}

func main() {
	// 明示的宣言
	var user1 User = User{"sample", 20}
	fmt.Println(user1)

	// 暗黙的宣言
	user2 := User{"sample2", 30}
	fmt.Println(user2)

	// field指定して宣言
	user3 := User{Name: "sample3", Age: 40}
	fmt.Println(user3)

	// 各フィールドを指定
	fmt.Println(user3.Name)
	fmt.Println(user3.Age)

	// ポインタ型で宣言
	user4 := &User{"sample4", 50}
	fmt.Println(user4)
}