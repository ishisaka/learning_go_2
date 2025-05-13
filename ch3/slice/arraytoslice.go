// 配列のスライスへの変換
package main

import "fmt"

func main() {
	xArray := [4]int{5, 6, 7, 8}
	xSlice := xArray[:]
	fmt.Println(xSlice)

	x := [4]int{5, 6, 7, 8}
	y := x[:2]
	z := x[2:]
	fmt.Println(x, y, z)

	// 配列をスライスにしてもコピーされるのではなく参照となる。
	x = [4]int{5, 6, 7, 8}
	y = x[:2]
	z = x[2:]
	x[0] = 10
	fmt.Println("x:", x) // x: [10 6 7 8]
	fmt.Println("y:", y) // y: [10 6]
	fmt.Println("z:", z) // z: [7 8]
}
