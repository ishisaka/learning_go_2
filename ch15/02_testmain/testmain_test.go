package testmain

import (
	"fmt"
	"os"
	"testing"
	"time"
)

var testTime time.Time

func TestMain(m *testing.M) {
	fmt.Println("テストのセットアップ開始")
	testTime = time.Now()
	exitVal := m.Run()
	fmt.Println("テスト終了")
	fmt.Println("テストの実行時間:", time.Since(testTime))
	os.Exit(exitVal)
}

func TestFirst(t *testing.T) {
	fmt.Println("テスト1", testTime)
}

func TestSecond(t *testing.T) {
	fmt.Println("テスト2", testTime)
}
