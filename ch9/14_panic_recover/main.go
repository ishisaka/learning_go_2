// パニックとリカバー
package main

import (
	"fmt"
)

func div60(i int) {
	defer func() {
		// recover関数でPanicをリカバーできる
		if v := recover(); v != nil {
			fmt.Println(v)
		}
	}()
	fmt.Println(60 / i)
}

func main() {
	for _, val := range []int{1, 2, 0, 6} {
		div60(val)
	}
}
