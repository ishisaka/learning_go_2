// センチネルエラーの例
package main

import (
	"archive/zip"
	"bytes"
	"fmt"
)

func main() {
	data := []byte("This is not a zip file")
	notAZipFile := bytes.NewReader(data)
	_, err := zip.NewReader(notAZipFile, int64(len(data)))
	// センチネルエラーの処理
	if err == zip.ErrFormat {
		fmt.Println("Told you so")
	}
}
