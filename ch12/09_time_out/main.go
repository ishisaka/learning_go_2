// タイムアウト処理
package main

import (
	"context"
	"errors"
	"fmt"
	"math/rand"
	"time"
)

func main() {
	result, err := timeLimit(doSomeWork, 2*time.Second)
	fmt.Println(result, err)
}

// timeLimitはworkerを呼び出し、limit時間内に結果を返すか、
// タイムアウトした場合にエラーを返す。
func timeLimit[T any](worker func() T, limit time.Duration) (T, error) {
	out := make(chan T, 1)
	// タイムアウトのコンテキストをセット
	ctx, cancel := context.WithTimeout(context.Background(), limit)
	defer cancel()
	go func() {
		out <- worker()
	}()
	select {
	case result := <-out:
		return result, nil
	case <-ctx.Done():
		// タイムアウトした場合の処理
		var zero T
		return zero, errors.New("work timed out")
	}
}

func doSomeWork() int {
	if x := rand.Int(); x%2 == 0 {
		return x
	} else {
		time.Sleep(100 * time.Second)
		return 100
	}
}
