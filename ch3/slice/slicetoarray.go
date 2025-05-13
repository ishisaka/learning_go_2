// スライスの配列への変換
package main

import "fmt"

func main() {
	// スライスを配列に変換するとコピーとなる
	xSlice := []int{1, 2, 3, 4}
	xArray := [4]int(xSlice)
	smallArray := [2]int(xSlice)
	xSlice[0] = 10
	fmt.Println(xSlice)     // [10 2 3 4]
	fmt.Println(xArray)     // [1 2 3 4]
	fmt.Println(smallArray) // [1 2]

	panicArray := [5]int(xSlice) // 配列の長さがスライスより小さいとパニックになる
	fmt.Println(panicArray)
}
