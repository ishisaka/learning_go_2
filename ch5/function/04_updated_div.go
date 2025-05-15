// 複数の戻り値
package main

import (
	"errors"
	"fmt"
	"os"
)

// 複数の戻り値を返す関数
func divAndRemainder(num, denom int) (int, int, error) {
	if denom == 0 {
		return 0, 0, errors.New("cannot divide by zero")
	}
	return num / denom, num % denom, nil
}

func main() {
	// 以下のように複数の戻り値を受け取る
	result, remainder, err := divAndRemainder(5, 2)
	if err != nil {
		fmt.Println(err)
		os.Exit(1)
	}
	fmt.Println(result, remainder)
	// 戻り値の無視 _ で戻り値を受ける
	_, _, err = divAndRemainder(5, 0)
	if err != nil {
		fmt.Println(err)
		os.Exit(1)
	}
}
