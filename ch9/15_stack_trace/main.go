// github.com/cockroachdb/errorsパッケージを使ってスタックトレースを表示
package main

import (
	"fmt"
	"github.com/cockroachdb/errors"
)

func innerFunc() error {
	// return errors.New("エラーが発生しました")
	// errors.WithStack() を使用すると、明示的にスタックトレースを付与できます。
	return errors.WithStack(errors.New("エラーが発生しました"))
}

func outerFunc() error {
	if err := innerFunc(); err != nil {
		return errors.WithMessage(err, "outerFuncでinnerFuncが失敗")
	}
	return nil
}

func main() {
	err := outerFunc()
	if err != nil {
		// %+v でスタックトレースを表示
		fmt.Printf("%+v\n", err)
	}
}
