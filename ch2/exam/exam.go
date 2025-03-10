package main

import "fmt"

func main() {
	// 課題1
	var i int = 20
	var f float64 = float64(i)
	fmt.Println(i, f)
	// 課題2
	const value = 30
	i = value
	var j float64 = value
	fmt.Println(i, j)
	// 課題3
	var b byte = 255
	var smallI int32 = 2147483647
	var bigI uint64 = 18446744073709551615

	b = b + 1
	smallI = smallI + 1
	bigI = bigI + 1

	fmt.Println(b)
	fmt.Println(smallI)
	fmt.Println(bigI)
}
