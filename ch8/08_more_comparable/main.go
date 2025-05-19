// さらに比較可能に
package main

import (
	"fmt"
	"reflect"
)

// Thinger インターフェースは、Thing メソッドを実装する型を定義します。
type Thinger interface {
	Thing()
}

// ThingerInt は int を基にした新しい型を表します。
// Thinger インターフェースを実装することが可能です。
// 主に整数値を扱う用途で使用されます。
type ThingerInt int

// Thing は ThingerInt 型の値を使用して処理を行い、出力を表示します。
func (t ThingerInt) Thing() {
	fmt.Println("ThingInt:", t)
}

// ThingerSlice は int スライスを基にした新しい型を表します。
// Thinger インターフェースを実装することが可能です。
type ThingerSlice []int

// Thing は ThingerSlice 型の値を使用して処理を行い、出力を表示します。
func (t ThingerSlice) Thing() {
	fmt.Println("ThingSlice:", t)
}

// Comparer はジェネリック関数として、2 つの値 t1 と t2 を比較します。
// t1 と t2 が等しい場合、等しいことを示すメッセージを出力します。
// それぞれの値が比較可能かどうかも確認し、その結果を出力します。
func Comparer[T comparable](t1, t2 T) {
	v1 := reflect.ValueOf(t1)
	v2 := reflect.ValueOf(t2)
	fmt.Println("comparable v1:", v1.Comparable())
	fmt.Println("comparable v2:", v2.Comparable())
	if t1 == t2 {
		fmt.Println("equal!")
	}
}

func main() {
	var a int = 10
	var b int = 10
	Comparer(a, b)

	var a2 ThingerInt = 20
	var b2 ThingerInt = 20
	Comparer(a2, b2)

	var a3 ThingerSlice = []int{1, 2, 3}
	var b3 ThingerSlice = []int{1, 2, 3}
	//Comparer(a3, b3)

	var a4 Thinger = a2
	var b4 Thinger = b2
	Comparer(a4, b4)

	a4 = a3
	b4 = b3
	Comparer(a4, b4) // compiles, panics at runtime See https://go.dev/blog/comparable

	b4 = b2
	Comparer(a4, b4)
}
