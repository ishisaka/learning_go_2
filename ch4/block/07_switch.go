// switch文
package main

import "fmt"

func main() {
	// switch文
	{
		words := []string{"a", "cow", "smile", "gopher",
			"octopus", "anthropologist"}
		for _, word := range words {
			switch size := len(word); size {
			case 1, 2, 3, 4:
				fmt.Println(word, "is a short word!")
			case 5:
				wordLen := len(word)
				fmt.Println(word, "is exactly the right length:", wordLen)
			case 6, 7, 8, 9:
			default:
				fmt.Println(word, "is a long word!")
			}
		}
	}
	// switch文の結果からループを抜けたい
	fmt.Println("---")
	{
	loop:
		for i := 0; i < 10; i++ {
			switch i {
			case 0, 2, 4, 6:
				fmt.Println(i, "is even")
			case 3:
				fmt.Println(i, "is divisible by 3 but not 2")
			case 7:
				fmt.Println("exit the loop!")
				// break これだとswitchから抜けるだけで意図したループからの脱出にならない
				break loop // ループを脱出する為にラベルを使用する
			default:
				fmt.Println(i, "is boring")
			}
		}
	}
	// 空白のswitch
	fmt.Println("---")
	{
		words := []string{"hi", "salutations", "hello"}
		for _, word := range words {
			switch wordLen := len(word); { // 空白のswitch
			case wordLen < 5:
				fmt.Println(word, "is a short word!")
			case wordLen > 10:
				fmt.Println(word, "is a long word!")
			default:
				fmt.Println(word, "is exactly the right length.")
			}
		}
	}
}
