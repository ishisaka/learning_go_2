// 関数は値
package main

import "fmt"

func f1(a string) int {
	return len(a)
}

func f2(a string) int {
	total := 0
	for _, v := range a {
		total += int(v)
	}
	return total
}

func main() {
	// 関数は値なので、関数変数を宣言できる
	var myFuncVariable func(string) int
	// 引数と戻り値の型が一致すれば関数を関数の変数に代入できる
	myFuncVariable = f1
	result := myFuncVariable("Hello")
	fmt.Println(result)

	myFuncVariable = f2
	result = myFuncVariable("Hello")
	fmt.Println(result)
}
