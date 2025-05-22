// Goroutine
package main

import "fmt"

func main() {
	x := []int{1, 2, 3, 4, 5, 6, 7, 8, 9, 10, 11, 12, 13, 14, 15, 16, 17, 18, 19, 20}
	result := processConcurrently(x)
	fmt.Println(result)
}

// process は、整数を受け取り、その値を2倍にして返す関数です。
func process(val int) int {
	// do something with val
	return val * 2
}

const numGoroutines = 5

// processConcurrently は、入力された整数スライスを並列処理し、処理後の結果スライスを返す関数です。
// 入力チャネルと出力チャネルを使用して、並列処理を効率的に実行します。
// 並列処理のユニット数は固定数のゴルーチンで実行され、numGoroutines に依存します。
// 処理の全体が完了するまで、出力チャネルを順次読み取って結果を収集します。
// 入力スライスを指定すると、ゴルーチンによって process() が並列実行されます。
// チャネルは適切にクローズされることで、リソースリークを防ぎます。
func processConcurrently(inVals []int) []int {
	// create the channels
	in := make(chan int, numGoroutines)
	out := make(chan int, numGoroutines)
	// launch numGoroutines
	for i := 0; i < numGoroutines; i++ {
		go func() {
			for val := range in {
				result := process(val)
				out <- result
			}
		}()
	}
	// load the data into the channel in another goroutine
	go func() {
		for _, v := range inVals {
			in <- v
		}
		close(in)
	}()
	// read the data
	outVals := make([]int, 0, len(inVals))
	for i := 0; i < len(inVals); i++ {
		outVals = append(outVals, <-out)
	}
	return outVals
}
