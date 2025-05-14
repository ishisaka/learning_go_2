// for文のブロック
package main

import "fmt"

func main() {
	//完全なfor文
	{
		for i := 0; i < 10; i++ {
			fmt.Println(i)
		}
	}
	// 初期化文の省略
	fmt.Println("---")
	{
		i := 0
		for ; i < 10; i++ {
			fmt.Println(i)
		}
	}
	// インクリメントの省略
	fmt.Println("---")
	{
		for i := 0; i < 10; {
			fmt.Println(i)
			if i%2 == 0 {
				i++
			} else {
				i += 2
			}
		}
	}
	// 条件のみのfor
	fmt.Println("---")
	{
		i := 1
		for i < 100 {
			fmt.Println(i)
			i = i * 2
		}
	}
	// continue
	fmt.Println("---")
	{
		for i := 1; i <= 100; i++ {
			if i%3 == 0 && i%5 == 0 {
				fmt.Println("FizzBuzz")
				continue
			}
			if i%3 == 0 {
				fmt.Println("Fizz")
				continue
			}
			if i%5 == 0 {
				fmt.Println("Buzz")
				continue
			}
			fmt.Println(i)
		}
	}
	// break
	fmt.Println("---")
	{
		CONDITION := true
		for {
			// things to do in the loop
			if !CONDITION {
				fmt.Println("break")
				break
			}
			CONDITION = false
		}
	}
}
