// if文のブロック
package main

import (
	"fmt"
	"math/rand"
)

func main() {
	// ifとelse
	{
		n := rand.Intn(10)
		if n == 0 {
			fmt.Println("That's too low")
		} else if n > 5 {
			fmt.Println("That's too big:", n)
		} else {
			fmt.Println("That's a good number:", n)
		}
	}
	// if文への変数のスコープ
	fmt.Println("---")
	{
		// 変数nはif文のブロック内でのみ有効な変数
		if n := rand.Intn(10); n == 0 {
			fmt.Println("That's too low")
		} else if n > 5 {
			fmt.Println("That's too big:", n)
		} else {
			fmt.Println("That's a good number:", n)
		}
	}
	// スコープ外
	fmt.Println("---")
	{
		if n := rand.Intn(10); n == 0 {
			fmt.Println("That's too low")
		} else if n > 5 {
			fmt.Println("That's too big:", n)
		} else {
			fmt.Println("That's a good number:", n)
		}
		fmt.Println(n) // 変数nスコープ外なのでコンパイルエラー
	}

}
