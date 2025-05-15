package main

import "fmt"

// 匿名関数の定義をパッケージレベルで行う
var (
	add = func(i int, j int) int { return i + j }
	sub = func(i int, j int) int { return i - j }
	mul = func(i int, j int) int { return i * j }
	div = func(i int, j int) int { return i / j }
)

func main() {
	x := add(2, 3)
	fmt.Println(x)
	changeAdd()
	y := add(2, 3)
	fmt.Println(y)
}

func changeAdd() {
	add = func(i int, j int) int { return i + j + j }
}
