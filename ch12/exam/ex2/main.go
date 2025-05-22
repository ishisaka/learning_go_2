package main

import "fmt"

/*
2つのゴルーチンを起動する関数を作成します。各ゴルーチンは、独自のチャネルに10個の数値を書き込みます。
for-selectループを使用して両方のチャネルから読み取り、数値と値を書き込んだゴルーチンを出力します。
すべての値が読み取られた後に関数が終了し、どのゴルーチンもリークしないことを確認します。
*/

// ProcessData は2つのゴルーチンを使用してデータを生成し、チャネル経由で処理する非同期処理を行う関数です。
// 2つのチャネルからデータを受信し、受信されたデータを標準出力に表示します。
// チャネルが閉じられると、それ以上のデータ受信は行われず、処理が終了します。
// セレクトステートメントを活用して、複数のチャネルからのデータ受信を並行して処理します。
func ProcessData() {
	ch := make(chan int)
	ch2 := make(chan int)
	go func() {
		for i := 0; i < 10; i++ {
			ch <- i
		}
		close(ch)
	}()
	go func() {
		for i := 0; i < 10; i++ {
			ch2 <- i*100 + 1
		}
		close(ch2)
	}()
	// once a channel is closed, ok will return false. Use that information to set the channel variable to nil,
	// disabling the case. When both cases are disabled, you are done.
	count := 2
	for count != 0 {
		select {
		case v, ok := <-ch:
			if !ok {
				ch = nil
				count--
				break
			}
			fmt.Println(v)
		case v, ok := <-ch2:
			if !ok {
				ch2 = nil
				count--
				break
			}
			fmt.Println(v)
		}
	}
}

func main() {
	ProcessData()
}
