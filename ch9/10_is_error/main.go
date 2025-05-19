// errors.Isでエラーを確認する。
package main

import (
	"errors"
	"fmt"
	"os"
)

// ファイルの存在を確認する。
func fileChecker(name string) error {
	f, err := os.Open(name)
	if err != nil {
		return fmt.Errorf("in fileChecker: %w", err)
	}
	f.Close()
	return nil
}

func main() {
	err := fileChecker("not_here.txt")
	if err != nil {
		// エラーがセンチネルエラーの型と一致するかIs関数で確認する。
		if errors.Is(err, os.ErrNotExist) {
			fmt.Println("That file doesn't exist")
		}
	}
}
