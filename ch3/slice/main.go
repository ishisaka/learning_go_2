// このコードはコンパイルが通りません
package main

import (
	"fmt"
	"log"
	"slices"
)

func main() {
	// スライスの初期化
	var x = []int{1, 2, 3}
	println(x[0])
	println(x[1])
	println(x[2])

	println("")
	// 途中空きのあるスライスの初期化
	x = []int{1, 5: 4, 6, 10: 100, 13}
	for i := 0; i < len(x); i++ {
		println(x[i])
	}
	println("")
	// スライスへの値の代入
	x[0] = 10
	for i := 0; i < len(x); i++ {
		println(x[i])
	}
	// 多次元のスライスの宣言
	var xx [][]int

	// 空のスライスはnil
	var yy []int
	log.Println(yy == nil)

	x = []int{1, 2, 3, 4, 5}
	y := []int{1, 2, 3, 4, 5}
	z := []int{1, 2, 3, 4, 5, 6}
	s := []string{"a", "b", "c"}
	fmt.Println(slices.Equal(x, y)) // prints true
	fmt.Println(slices.Equal(x, z)) // prints false
	fmt.Println(slices.Equal(x, s)) // does not compile
}
