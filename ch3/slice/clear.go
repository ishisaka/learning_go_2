// スライスを空にする
package main

import "fmt"

func main() {
	s := []string{"first", "second", "third"}
	fmt.Println(s, len(s), cap(s))
	// スライスを空にするにはclear関数を使用する
	clear(s)
	fmt.Println(s, len(s), cap(s))
}
