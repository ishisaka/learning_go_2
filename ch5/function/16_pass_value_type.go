// Goの関数は値渡し
package main

import "fmt"

type person struct {
	age  int
	name string
}

func modifyFails(i int, s string, p person) {
	i = i * 2
	s = "Goodbye"
	p.name = "Bob"
}

func main() {
	p := person{}
	i := 2
	s := "Hello"
	modifyFails(i, s, p)
	fmt.Println(i, s, p) // 2 Hello {0 } 値渡しなので、元々の値は変更されない
}
