// selectを使ったデッドロックの回避
package main

import (
	"fmt"
)

func main() {
	ch1 := make(chan int)
	ch2 := make(chan int)
	go func() {
		inGoroutine := 1
		ch1 <- inGoroutine
		fromMain := <-ch2
		fmt.Println("goroutine:", inGoroutine, fromMain)
	}()
	inMain := 2
	// select文で2つの通信操作のうちどちらかを実行
	// デッドロックは回避されるが期待通りには動いていない
	var fromGoroutine int
	select {
	case ch2 <- inMain:
	case fromGoroutine = <-ch1:
	}
	fmt.Println("main:", inMain, fromGoroutine) // main: 2 1
}
