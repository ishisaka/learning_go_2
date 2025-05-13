package main

import (
	"log"
)

func main() {
	// 配列の宣言
	var x [3]int
	x[0] = 1
	x[1] = 2
	x[2] = 3
	for i := 0; i < 3; i++ {
		println(x[i])
	}
	println("")
	// 配列の宣言と初期化
	var y [3]int = [3]int{10, 20, 30}
	for i := 0; i < 3; i++ {
		println(y[i])
	}
	println("")
	// 途中に空きがある配列の宣言と初期化
	var z = [12]int{1, 5: 4, 6, 10: 100, 15}
	for i := 0; i < 12; i++ {
		println(z[i])
	}
	println("")
	// 配列の初期化にリテラルを使う
	x = [...]int{1, 2, 3}
	for i := 0; i < 3; i++ {
		println(x[i])
	}
	// 配列の比較
	x = [...]int{1, 2, 3}
	y = [3]int{1, 2, 3}
	log.Println(x == y)
	println("")
	// 多次元配列はジャグ配列で表現する
	var xx [2][3]int
	xx[0][0] = 1
	xx[0][1] = 2
	xx[0][2] = 3
	xx[1][0] = 4
	xx[1][1] = 5
	xx[1][2] = 6
	for i := 0; i < 2; i++ {
		for j := 0; j < 3; j++ {
			println(xx[i][j])
		}
	}
	// lenとcapは配列の長さと容量を返す
	log.Println(len(xx))
	log.Println(cap(xx))
}
