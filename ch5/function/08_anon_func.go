// 匿名関数
package main

import "fmt"

func main() {
	// 匿名関数の定義
	f := func(j int) {
		fmt.Println("printing", j, "from inside of an anonymous function")
	}
	for i := 0; i < 5; i++ {
		f(i)
	}
}
