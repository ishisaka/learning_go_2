// Contextを使ってGoルーチンを安全に中断させる
package main

import (
	"context"
	"fmt"
)

func countTo(ctx context.Context, max int) <-chan int {
	ch := make(chan int)
	go func() {
		defer close(ch)
		// for-select loop
		for i := 0; i < max; i++ {
			select {
			// ctx.Done()が送られてきたら処理を中断しチャンネルをクローズする
			case <-ctx.Done():
				return
			case ch <- i:
			}
		}
	}()
	return ch
}

func main() {
	ctx, cancel := context.WithCancel(context.Background())
	ch := countTo(ctx, 10)
	for i := range ch {
		if i > 5 {
			break
		}
		fmt.Println(i)
	}
	// ctx.Done()を発生させる
	cancel()
}
