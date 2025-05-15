package main

import "fmt"

func main() {
	a := 20
	// クロージャー
	f := func() {
		fmt.Println(a) // 20
		// 変数のシャドウイング
		a := 30
		fmt.Println(a) // 30
	}
	f()
	fmt.Println(a) // 20
}
