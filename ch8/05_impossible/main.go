// 型タームを使った演算子の指定
package main

import "fmt"

type ImpossiblePrintableInt interface {
	int
	String() string
}

type ImpossibleStruct[T ImpossiblePrintableInt] struct {
	val T
}

type MyInt int

func (mi MyInt) String() string {
	return fmt.Sprint(mi)
}

func main() {
	s := ImpossibleStruct[int]{10} // コンパイルエラー
	fmt.Println(s)
	s2 := ImpossibleStruct[MyInt]{10} // コンパイルエラー
	fmt.Println(s2)
}
