// ポインタレシーバと値レシーバ
package main

import (
	"fmt"
	"time"
)

type Counter struct {
	total       int
	lastUpdated time.Time
}

// ポインタレシーバ
func (c *Counter) Increment() {
	c.total++
	c.lastUpdated = time.Now()
}

// 型のインスタンスの値をメソッド内で変更する場合
// ポインタレシーバを用いる

// 値レシーバ
func (c Counter) String() string {
	return fmt.Sprintf("total: %d, last updated: %v", c.total, c.lastUpdated)
}

// 型のインスタンスの値をメソッド内で変更しない場合
// もしくは、変更したくない場合には
// 値レシーバを用いる

func doUpdateWrong(c Counter) {
	c.Increment()
	fmt.Println("in doUpdateWrong:", c.String())
}

func doUpdateRight(c *Counter) {
	c.Increment()
	fmt.Println("in doUpdateRight:", c.String())
}

func main() {
	var c Counter
	fmt.Println(c.String())
	c.Increment()
	fmt.Println(c.String())
	fmt.Println("----")
	// メソッド呼び出し時にcの値はコピーされるので元々のcの値は
	// doUpdateWrongで変更されない
	doUpdateWrong(c)
	fmt.Println("in main:", c.String())
	// メソッド呼び出し時にcのポインタがdoUpdateRightに渡されるので
	// doUpdateRight内でのcへの値変更は反映される
	doUpdateRight(&c)
	fmt.Println("in main:", c.String())
}
