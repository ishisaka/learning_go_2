package main

import "fmt"

func main() {
	// 変数のシャドウイング
	{
		x := 10
		if x > 5 {
			fmt.Println(x) // 10
			// 以下がシャドウイング変数となる
			x := 5
			fmt.Println(x) // 5
		}
		fmt.Println(x) // 10
	}
	// 複数のシャドウイング
	fmt.Println("---")
	{
		x := 10
		if x > 5 {
			x, y := 5, 20
			fmt.Println(x, y)
		}
		fmt.Println(x)
	}
	// ユニバースブロックのシャドウイング
	fmt.Println("---")
	{
		fmt.Println(true) // trye
		// ユニバースブロックで定義されている true をシャドウイング
		true := 10
		fmt.Println(true) // 10
	}
	// パッケージ名のシャドウイング
	fmt.Println("---")
	{
		x := 10
		fmt.Println(x)
		fmt := "oops"    // パッケージ名がシャドウイングされてしまう
		fmt.Println(fmt) // ここでエラーとなる
	}

}
