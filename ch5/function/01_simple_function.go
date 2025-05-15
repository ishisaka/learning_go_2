// シンプルな関数
package main

import "fmt"

// div は整数の割り算を行い、商を返します。分母が 0 の場合は 0 を返します。
func div(num int, denom int) int {
	if denom == 0 {
		return 0
	}
	return num / denom
}

func main() {
	result := div(5, 2)
	fmt.Println(result)
}
