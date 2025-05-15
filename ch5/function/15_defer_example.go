// 複数のdefer
package main

import "fmt"

func main() {
	deferExample()
}

// deferされた関数は後入れ先出しで実行される
func deferExample() int {
	a := 10
	defer func(val int) {
		fmt.Println("first:", val)
	}(a)
	a = 20
	defer func(val int) {
		fmt.Println("second:", val)
	}(a)
	a = 30
	fmt.Println("exiting:", a)
	return a
}
