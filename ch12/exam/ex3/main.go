package main

import (
	"fmt"
	"math"
	"sync"
)

/*
map[int]float64を構築する関数を記述します。この関数では、キーは0（含む）から100,000（含まない）までの数値で、
値はそれらの数値の平方根です（平方根を計算するには、math.Sqrt関数を使用します）。
sync.OnceValueを使用して、この関数によって返されるマップをキャッシュする関数を生成し、
キャッシュされた値を使用して0から100,000までの1,000番目の数値ごとに平方根を検索します。
*/

// buildSquareRootMap は 0 から 99,999 の整数の平方根を計算し、結果をマップとして返します。
// 各キーは整数、その値は対応する平方根です。
// 内部でマップの初期容量を 100,000 に設定し、効率的なメモリ管理を実現します。
func buildSquareRootMap() map[int]float64 {
	out := make(map[int]float64, 100_000)
	for i := 0; i < 100_000; i++ {
		out[i] = math.Sqrt(float64(i))
	}
	return out
}

var squareRootMapCache = sync.OnceValue(buildSquareRootMap)

func main() {
	for i := 0; i < 100_000; i += 1_000 {
		fmt.Println(i, squareRootMapCache()[i])
	}
}
