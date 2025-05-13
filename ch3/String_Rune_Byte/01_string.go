package main

import "fmt"

func main() {
	var s string = "Hello there"
	var b byte = s[6]
	fmt.Println(b)
	var s2 string = s[4:7]
	var s3 string = s[:5]
	var s4 string = s[6:]
	fmt.Println(s2)
	fmt.Println(s3)
	fmt.Println(s4)
	{
		var s string = "Hello 🌞"
		var s2 string = s[4:7]
		var s3 string = s[:5]
		var s4 string = s[6:]
		fmt.Println(s2)
		fmt.Println(s3)
		fmt.Println(s4)
		fmt.Println(len(s))
	}
	// byteやruneをstringに変換
	{
		var a rune = 'x'
		var s string = string(a)
		var b byte = 'y'
		var s2 string = string(b)
		fmt.Println(s, s2)
	}
	// 文字列をbyteもしくはruneのスライスに変換
	{
		var s string = "Hello, 🌞"
		var bs []byte = []byte(s)
		var rs []rune = []rune(s)
		fmt.Println(bs)
		fmt.Println(rs)
	}
}
