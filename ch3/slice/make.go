// make スライスのキャパシティを予め設定する
package main

import "fmt"

func main() {
	x := make([]int, 5) // 長さ5のスライス
	x = append(x, 10)
	fmt.Println(x)
	fmt.Println(len(x), cap(x))

	y := make([]int, 5, 10) // 長さ5キャパシティ10のスライス
	y[4] = 11
	y = append(y, 10)
	fmt.Println(y)
	fmt.Println(len(y), cap(y))

}
