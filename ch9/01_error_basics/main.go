// エラー処理の基本
package main

import (
	"errors"
	"fmt"
	"os"
)

// calcRemainderAndMod は指定された分子と分母を使って商と余りを計算し、それらを返します。
// 分母が 0 の場合はエラーを返します。
func calcRemainderAndMod(numerator, denominator int) (int, int, error) {
	if denominator == 0 {
		return 0, 0, errors.New("denominator is 0")
	}
	return numerator / denominator, numerator % denominator, nil
}

func main() {
	numerator := 20
	denominator := 3
	// 基本的なエラー処理
	remainder, mod, err := calcRemainderAndMod(numerator, denominator)
	// errがnilか調べてnil出なかったら必要な処理を行う
	// エラーがないことを示すのにnilを返すのは、nilがインタフェース型のゼロ値だから
	if err != nil {
		fmt.Println(err)
		os.Exit(1)
	}
	fmt.Println(remainder, mod)
}
