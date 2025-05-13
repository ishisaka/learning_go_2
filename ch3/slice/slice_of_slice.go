// スライスのスライス
package main

import "fmt"

func main() {
	x := []string{"a", "b", "c", "d"}
	// スライスのスライスは、スライスの切り出しを表す
	// コピーではないので、切り出したスライスの変更は
	// 元のスライスにも反映される
	y := x[:2]
	z := x[1:]
	d := x[1:3]
	e := x[:]
	fmt.Println("x:", x)
	fmt.Println("y:", y)
	fmt.Println("z:", z)
	fmt.Println("d:", d)
	fmt.Println("e:", e)

	fmt.Println("")
	x = []string{"a", "b", "c", "d"}
	y = x[:2]
	z = x[1:]
	x[1] = "y"
	y[0] = "x"
	z[1] = "z"
	fmt.Println("x:", x)
	fmt.Println("y:", y)
	fmt.Println("z:", z)
}
