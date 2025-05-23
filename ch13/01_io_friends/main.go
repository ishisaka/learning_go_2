package main

import (
	"compress/gzip"
	"fmt"
	"io"
	"log/slog"
	"os"
	"strings"
)

// countLetters 関数は io.Reader を入力として英字の出現回数をカウントするマップを返します。
// io.EOF で正常終了し、それ以外のエラーが発生した場合はエラーを返します。
// 大文字と小文字を区別してカウントします。
func countLetters(r io.Reader) (map[string]int, error) {
	buf := make([]byte, 2048)
	out := map[string]int{}
	for {
		n, err := r.Read(buf)
		for _, b := range buf[:n] {
			if (b >= 'A' && b <= 'Z') || (b >= 'a' && b <= 'z') {
				out[string(b)]++
			}
		}
		if err == io.EOF {
			return out, nil
		}
		if err != nil {
			return nil, err
		}
	}
}

// simpleCountLetters 関数は文字列から英字の出現回数をカウントし、結果を標準出力に表示します。
// countLetters 関数を利用して処理し、処理中にエラーが発生した場合はエラーを返します。
func simpleCountLetters() error {
	s := "The quick brown fox jumped over the lazy dog"
	sr := strings.NewReader(s)
	counts, err := countLetters(sr)
	if err != nil {
		return err
	}
	fmt.Println(counts)
	return nil
}

// buildGZipReader 関数はファイル名を入力として、gzip.Reader を返します。
// 処理中にエラーが発生した場合はエラーを返します。
func buildGZipReader(fileName string) (*gzip.Reader, func(), error) {
	r, err := os.Open(fileName)
	if err != nil {
		return nil, nil, err
	}
	gr, err := gzip.NewReader(r)
	if err != nil {
		return nil, nil, err
	}
	return gr, func() {
		gr.Close()
		r.Close()
	}, nil
}

// gzipCountLetters 関数はファイル名を入力として、gzip.Reader を利用して、
// 英字の出現回数をカウントし、結果を標準出力に表示します。
// buildGZipReader 関数を利用して処理し、処理中にエラーが発生した場合はエラーを返します。
func gzipCountLetters() error {
	r, closer, err := buildGZipReader("my_data.txt.gz")
	if err != nil {
		return err
	}
	defer closer()
	counts, err := countLetters(r)
	if err != nil {
		return err
	}
	fmt.Println(counts)
	return nil
}

func main() {
	err := simpleCountLetters()
	if err != nil {
		slog.Error("error with simpleCountLetters", "msg", err)
	}

	err = gzipCountLetters()
	if err != nil {
		slog.Error("error with gzipCountLetters", "msg", err)
	}
}
