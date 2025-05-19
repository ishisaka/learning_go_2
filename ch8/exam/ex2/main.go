package main

import (
	"fmt"
	"strconv"
)

// Define a generic interface called `Printable` that matches a type that implements
// fmt.Stringer and has an underlying type of int or float64. Define types that meet
// this interface. Write a function that takes in a Printable and prints its value
// to the screen using fmt.Println.

// Printable は fmt.Stringer インターフェースを埋め込み、整数や浮動小数点型を許容する制約を定義します。
type Printable interface {
	fmt.Stringer
	~int | ~float64
}

type PrintInt int

// String はPrintInt型の値を文字列として返します。
func (pi PrintInt) String() string {
	return strconv.Itoa(int(pi))
}

// PrintFloat はPrintFloat型の値を文字列として返します。
type PrintFloat float64

// String はPrintFloat型の値を文字列として返します。
func (pf PrintFloat) String() string {
	return fmt.Sprintf("%f", pf)
}

// PrintIt はPrintable制約を満たす値を受け取り、それを標準出力に出力します。
func PrintIt[T Printable](t T) {
	fmt.Println(t)
}

func main() {
	var i PrintInt = 20
	PrintIt(i)

	var f PrintFloat = 10.23
	PrintIt(f)
}
