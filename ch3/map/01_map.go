// マップ
package main

import (
	"fmt"
	"maps"
)

func main() {
	// 空のマップ
	{
		var nilMap map[string]int
		fmt.Println(nilMap)
	}
	// リテラルで宣言
	{
		teams := map[string][]string{
			"Orcas":   []string{"Fred", "Ralph", "Bijou"},
			"Lions":   []string{"Sarah", "Peter", "Billie"},
			"Kittens": []string{"Waldo", "Raul", "Ze"},
		}
		fmt.Println(teams)
		// 要素無しで初期化
		totalWins := map[string]int{} // string→intのマップを要素なしで初期化
		fmt.Println(totalWins)
	}
	// make関数でキャパシティを指定しての初期化
	{
		ages := make(map[int][]string, 10)
		fmt.Println(ages)
	}
	// mapの読み書き
	{
		totalWins := map[string]int{}
		totalWins["Orcas"] = 1
		totalWins["Lions"] = 2
		fmt.Println(totalWins["Orcas"])
		// 一度も設定されていないキーの値を参照するとエラーではなく0値が返る
		fmt.Println(totalWins["Kittens"])
		// これは0値をインクリメントしている
		totalWins["Kittens"]++
		fmt.Println(totalWins["Kittens"])
		totalWins["Lions"] = 3
		fmt.Println(totalWins["Lions"])
	}
	// カンマokイディオム
	// マップを参照したときそのキーがなければ0値とfalseを返す
	// このためキーが有効かそうではない事の判定のために
	// カンマokイディオムを使用する
	{
		m := map[string]int{
			"hello": 5,
			"world": 0,
		}

		v, ok := m["hello"]
		fmt.Println(v, ok) // 5 true
		v, ok = m["world"]
		fmt.Println(v, ok) // 0 true
		v, ok = m["goodbye"]
		fmt.Println(v, ok) // 0 false
	}
	// mapからの削除
	{
		m := map[string]int{
			"hello": 5,
			"world": 10,
		}
		fmt.Println(m)
		delete(m, "hello")
		fmt.Println(m)
	}
	// mapを空にする
	{
		m := map[string]int{
			"hello": 5,
			"world": 10,
		}
		fmt.Println(m, len(m))
		clear(m)
		fmt.Println(m, len(m))
	}
	// mapの比較
	{
		m := map[string]int{
			"hello": 5,
			"world": 10,
		}
		n := map[string]int{
			"world": 10,
			"hello": 5,
		}
		fmt.Println(maps.Equal(m, n)) // prints true
	}
	// mapをset（集合）として扱う
	{
		intSet := map[int]bool{}
		vals := []int{5, 10, 2, 5, 8, 7, 3, 9, 1, 2, 10}
		for _, v := range vals {
			intSet[v] = true
		}
		fmt.Println(len(vals), len(intSet)) // 11 8
		fmt.Println(intSet[5])              // true
		fmt.Println(intSet[500])            // false
		// 集合に関しては以下のようなサードのライブラリがある
		// https://github.com/deckarep/golang-set
	}
}
