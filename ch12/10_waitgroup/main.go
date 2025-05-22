// WaitGroupの使用
package main

import (
	"fmt"
	"sync"
)

func main() {
	var wg sync.WaitGroup
	wg.Add(3)
	go func() {
		defer wg.Done() // 終了時にWaitGroupの数を減らす
		doThing1()
	}()
	go func() {
		defer wg.Done()
		doThing2()
	}()
	go func() {
		defer wg.Done()
		doThing3()
	}()
	// 3つのGoルーチンの終了待ち
	wg.Wait()
}

func doThing1() {
	fmt.Println("Thing 1 done!")
}

func doThing2() {
	fmt.Println("Thing 2 done!")
}

func doThing3() {
	fmt.Println("Thing 3 done!")
}
