// for-range文
package main

import "fmt"

func main() {
	// for-rangeループ
	{
		evenVals := []int{2, 4, 6, 8, 10, 12}
		for i, v := range evenVals {
			fmt.Println(i, v)
		}
	}
	// キーだけを取り出す
	{
		uniqueNames := map[string]bool{"Fred": true, "Raul": true, "Wilma": true}
		for k := range uniqueNames {
			fmt.Println(k)
		}
	}
	// 値だけを取り出す
	{
		uniqueNames := map[string]bool{"Fred": true, "Raul": true, "Wilma": true}
		for _, v := range uniqueNames {
			fmt.Println(v)
		}
	}
	// マップの反復処理
	// 毎回同じ順序でキーと値が返ってくるわけではない
	fmt.Println("---")
	{
		m := map[string]int{
			"a": 1,
			"c": 3,
			"b": 2,
		}

		for i := 0; i < 3; i++ {
			fmt.Println("Loop", i)
			for k, v := range m {
				fmt.Println(k, v)
			}
		}
	}
	// 文字列の反復処理
	fmt.Println("---")
	{
		samples := []string{"hello", "apple_π!"}
		for _, sample := range samples {
			// 文字列に対してrangeを使うとbyteではなくruneが返ってくる
			for i, r := range sample {
				fmt.Println(i, r, string(r))
			}
			fmt.Println()
		}
	}
	// for-rangeの値はコピーである
	fmt.Println("---")
	{
		evenVals := []int{2, 4, 6, 8, 10, 12}
		for _, v := range evenVals {
			v *= 2
		}
		fmt.Println(evenVals) // [2 4 6 8 10 12]
	}
}
