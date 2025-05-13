// スライスのコピー
package main

import "fmt"

func main() {
	x := []int{1, 2, 3, 4}
	y := make([]int, 4)
	// スライスのコピーは、copy関数を使用する
	num := copy(y, x)
	fmt.Println(y, num)

	// 長さの違うスライスへのコピーはコピー先の長さによって変わる
	x = []int{1, 2, 3, 4}
	y = make([]int, 2)  // 長さ2のスライスを作る
	num = copy(y, x)    // 先頭から2要素だけコピー
	fmt.Println(y, num) // [1 2] 2
}
