// 単純なエラーの際の文字列の使用
package main

import (
	"errors"
	"fmt"
)

// doubleEvenErrorsNew は整数を受け取り、偶数の場合は2倍した値を返し、奇数の場合はエラーを返します。
// 引数 i が偶数でない場合、「only even numbers are processed」というエラーを返します。
// 引数 i が偶数の場合、i を 2 倍して返します。
// エラーハンドリング用途の関数です。
func doubleEvenErrorsNew(i int) (int, error) {
	if i%2 != 0 {
		return 0, errors.New("only even numbers are processed")
	}
	return i * 2, nil
}

// doubleEvenFmtErrorf は整数を受け取り、偶数の場合は2倍した値を返し、奇数の場合はエラーを返します。
// fmt.Errorfを使用してエラーに書式指定子が含まれる詳細なメッセージを提供します。
// 引数 i が偶数でない場合、「%d isn't an even number」というエラーを返します。
// 引数 i が偶数の場合、i を 2 倍して返します。
// エラーハンドリング用途の関数です。
func doubleEvenFmtErrorf(i int) (int, error) {
	if i%2 != 0 {
		// fmt.Errorfを使うと書式指定子を使用できる
		return 0, fmt.Errorf("%d isn't an even number", i)
	}
	return i * 2, nil
}

func main() {
	result, err := doubleEvenErrorsNew(1)
	if err != nil {
		fmt.Println(err) // prints "only even numbers are processed"
	}
	fmt.Println(result)

	result, err = doubleEvenFmtErrorf(1)
	if err != nil {
		fmt.Println(err) // prints "1 isn't an even number"
	}
	fmt.Println(result)
}
