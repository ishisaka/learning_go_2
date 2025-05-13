// スライスの追加 append関数
package main

func main() {
	var x []int
	// スライスにデータを追加する際にはappend関数を使用する
	x = append(x, 1)
	x = append(x, 2)
	x = append(x, 3)
	for i := 0; i < len(x); i++ {
		println(x[i])
	}
	println("")
	x = append(x, 4, 5, 6)
	for i := 0; i < len(x); i++ {
		println(x[i])
	}
	println("")
	// スライスに別のスライスを追加する
	y := []int{20, 30, 40}
	x = append(x, y...)
	for i := 0; i < len(x); i++ {
		println(x[i])
	}
	println("")

}
