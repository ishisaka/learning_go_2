// deferを使ったエラーのラップ
package main

import "fmt"

//func DoSomeThings(val1 int, val2 string) (string, error) {
//	val3, err := doThing1(val1)
//	if err != nil {
//		return "", fmt.Errorf("in DoSomeThings: %w", err)
//	}
//	val4, err := doThing2(val2)
//	if err != nil {
//		return "", fmt.Errorf("in DoSomeThings: %w", err)
//	}
//	result, err := doThing3(val3, val4)
//	if err != nil {
//		return "", fmt.Errorf("in DoSomeThings: %w", err)
//	}
//	return result, nil
//}

// 上のコードをdeferを使うことで以下のように簡潔に記述できる

func DoSomeThings(val1 int, val2 string) (_ string, err error) {
	defer func() {
		if err != nil {
			err = fmt.Errorf("in DoSomeThings: %w", err)
		}
	}()
	val3, err := doThing1(val1)
	if err != nil {
		return "", err
	}
	val4, err := doThing2(val2)
	if err != nil {
		return "", err
	}
	return doThing3(val3, val4)
}
