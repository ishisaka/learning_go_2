package main

import (
	"fmt"
	"sync"
)

/*
チャネルを使用して通信する3つのゴルーチンを起動する関数を作成します。
最初の2つのゴルーチンはそれぞれ10個の数値をチャネルに書き込みます。
3番目のゴルーチンはチャネルからすべての数値を読み取り、それらを出力します。
すべての値が出力されたら、関数は終了する必要があります。
いずれのゴルーチンもリークしないことを確認してください。
必要に応じて追加のゴルーチンを作成できます。
*/

func ProcessData() {
	ch := make(chan int)
	// use 2 waitgroups!
	// the 1st waitgroup controls when to close the channel
	var wg sync.WaitGroup
	wg.Add(2)
	go func() {
		defer wg.Done()
		for i := 0; i < 10; i++ {
			ch <- i
		}
	}()
	go func() {
		defer wg.Done()
		for i := 0; i < 10; i++ {
			ch <- i*100 + 1
		}
	}()
	// launch this helper goroutine to close the channel when the two writing goroutines are done
	go func() {
		wg.Wait()
		close(ch)
	}()
	// the second waitgroup signals when the reading goroutine is done
	var wg2 sync.WaitGroup
	wg2.Add(1)
	go func() {
		defer wg2.Done()
		for v := range ch {
			fmt.Println(v)
		}
	}()
	wg2.Wait()
}

func main() {
	ProcessData()
}
