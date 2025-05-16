// メソッド
package main

import "fmt"

type Person struct {
	FirstName string
	LastName  string
	Age       int
}

// Stringメソッドを定義
// func とメソッド名の間に「レシーバー」を記述し、これがその型のメソッドであることを示す
func (p Person) String() string {
	return fmt.Sprintf("%s %s, age %d", p.FirstName, p.LastName, p.Age)
}

func main() {
	p := Person{"Fred", "Williamson", 25}
	output := p.String()
	fmt.Println(output)
	// Stringという名前のメソッドは以下のように記述しても型を文字列に変換するときに
	//自動で呼び出される特別なメソッド
	fmt.Println(p)
}
