// メソッドは関数
package main

import "fmt"

type Adder struct {
	start int
}

func (a Adder) AddTo(val int) int {
	return a.start + val
}

func main() {
	myAdder := Adder{start: 10}
	fmt.Println(myAdder.AddTo(5)) // prints 15
	f1 := myAdder.AddTo
	fmt.Println(f1(10)) // prints 20
	f2 := Adder.AddTo
	fmt.Println(f2(myAdder, 15)) // prints 25
}
