// バックプレッシャーの実装
// バッファ付きチャネルとselectステートメントを利用して
// システム内の同時リクエストの数を制限できる
package main

import (
	"errors"
	"net/http"
	"time"
)

type PressureGauge struct {
	ch chan struct{}
}

func New(limit int) *PressureGauge {
	return &PressureGauge{
		ch: make(chan struct{}, limit),
	}
}

func (pg *PressureGauge) Process(f func()) error {
	select {
	// チャネルに書き込める場合には処理を実行
	case pg.ch <- struct{}{}:
		f()
		<-pg.ch
		return nil
	default:
		// （バッファがいっぱいで）チャネルに書き込めない場合にはエラーを返す
		return errors.New("no more capacity")
	}
}

func doThingThatShouldBeLimited() string {
	time.Sleep(2 * time.Second)
	return "done"
}

func main() {
	pg := New(10)
	http.HandleFunc("/request", func(w http.ResponseWriter, r *http.Request) {
		err := pg.Process(func() {
			w.Write([]byte(doThingThatShouldBeLimited()))
		})
		if err != nil {
			w.WriteHeader(http.StatusTooManyRequests)
			w.Write([]byte("Too many requests"))
		}
	})
	http.ListenAndServe(":8080", nil)
}
