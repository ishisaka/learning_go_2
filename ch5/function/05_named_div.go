// 名前付き戻り値
package main

import (
	"errors"
	"fmt"
)

// 名前付き戻り値を返す関数
func divAndRemainder(num, denom int) (result int, remainder int, err error) {
	// 戻り値に名前を付けるということは関数内で使用する変数を宣言すること
	if denom == 0 {
		err = errors.New("cannot divide by zero")
		return result, remainder, err
	}
	result, remainder = num/denom, num%denom
	return result, remainder, err
}

// 名前付き戻り値を使用する場合にはブランクリターンを使わない
func divAndRemainder2(numerator int, denominator int) (result int, remainder int, err error) {
	if denominator == 0 {
		err = errors.New("0で割ることはできません")
		// ブランクリターンを使うとその時点での戻り値の値が全て返ってしまう
		return
	}
	result, remainder = numerator/denominator, numerator%denominator
	return
}

func main() {
	// 名前付き戻り値はあくまでもその関数ローカル。呼び出し側で強制されるわけではない
	x, y, z := divAndRemainder(5, 2)
	fmt.Println(x, y, z)
}
