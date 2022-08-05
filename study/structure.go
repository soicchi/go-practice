package main

import "fmt"

type User struct {
		Name string
		Age int
}

func updateName(user User, name string, age int) {
		user.Name = name
		user.Age = age
}

func updateName2(user *User, name string, age int) {
		user.Name = name
		user.Age = age
}

// (u User)の部分はレシーバ
func (u User) sayName() {
		fmt.Println(u.Name)
}

// レシーバはポインタ型で定義するのが基本
func (u *User) setName(name string) {
		u.Name = name
}

// コンストラクタの機能を持った関数（golangではこのようにコンストラクタを作成する）
func newUser(name string, age int) *User {
		return &User{Name: name, Age: age}
}


// スライス型の構造体
type Users []*User

/*  下記のようにもスライス型の構造体を定義できるが上記の方が一般的
		type Users struct {
			Users []*Users
		}
*/

func main() {
	// 明示的宣言
	var user1 User = User{"sample", 20}
	fmt.Println(user1)

	// 暗黙的宣言
	user2 := User{"sample2", 30}  // fieldを宣言しない場合は、field順にひきすうを与えないとエラーになる
	fmt.Println(user2)

	// field指定して宣言
	user3 := User{Name: "sample3", Age: 40}
	fmt.Println(user3)

	// 各フィールドを指定
	fmt.Println(user3.Name)
	fmt.Println(user3.Age)

	// ポインタ型で宣言
	user4 := &User{"sample4", 50}  // user4 := new(User{"sample4", 50})でも同じ
	fmt.Println(user4)

	updateName(user1, "hoge", 30)
	updateName2(user4, "huga", 40)

	fmt.Println(user1)
	fmt.Println(user4)

	user1.sayName()


	user2.setName("aaa")
	fmt.Println(user2)


	// コンストラクタの挙動を行う関数を作成
	user5 := newUser("hogehoge", 55)
	fmt.Println(user5.Name)


	users := Users{}
	users = append(users, &user1)
	users = append(users, &user2)
	users = append(users, &user3)
	users = append(users, user4)

	fmt.Println("Start")
	for _, v := range users {
		fmt.Println(*v)
	}


	m := map[int]User{}
	fmt.Println(m)
}