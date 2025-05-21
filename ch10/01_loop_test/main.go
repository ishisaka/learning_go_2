package main

import "fmt"

func main() {
	x := []int{1, 2, 3, 4, 5}
	for _, v := range x {
		fmt.Printf("%p\n", &v) // go 1.21以前とい以後で表示が異なる
	}
}
