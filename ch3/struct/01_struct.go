package main

import "fmt"

type person struct {
	name string // 名前
	age  int    // 年齢
	pet  string // ペット
}

func main() {
	// 基本
	{
		// 宣言
		var fred person
		bob := person{} // 構造体リテラル。全フィールドがゼロ値で初期化される
		fmt.Println(fred, bob)
		// リテラルで初期化
		julia := person{
			"ジュリア", // name
			40,     // age
			"cat",  // pet
		}
		fmt.Println(julia)
		// 名前付きフィールドで初期化
		beth := person{
			age:  30,
			name: "ベス",
		}
		fmt.Println(beth)
		// 構造体のフィールドへのアクセス
		bob.name = "ボブ"
		fmt.Println(bob.name) // ボブ
	}
	// 無名構造体
	{
		// 通常の構造体
		var person struct { // 変数personを無名の構造体として
			name string // ゼロ値で初期化される
			age  int
			pet  string
		}
		person.name = "ボブ"
		person.age = 50
		person.pet = "dog"
		fmt.Println(person)

		// 無名構造体
		pet := struct {
			name string
			kind string
		}{
			name: "ポチ",
			kind: "dog",
		}
		fmt.Println(pet)
	}
	// 構造体の比較と変換
	{
		type firstPerson struct {
			name string
			age  int
		}

		f := firstPerson{
			name: "Bob",
			age:  50,
		}

		var g struct {
			name string
			age  int
		}
		// 構造体は宣言が異なっていてもフィールド名と型、その並び順が一致していれば
		// 代入も比較も行える
		g = f
		fmt.Println(f == g) // true
	}
}
