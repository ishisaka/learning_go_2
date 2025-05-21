package main

import (
	"errors"
	"fmt"
)

func main() {
	err := returnErr(false)
	if err != nil {
		fmt.Println(err)
	}
	err = returnErr(true)
	fmt.Println("end of program")
}

func returnErr(b bool) error {
	if b {
		return errors.New("err")
	}
	return nil
}

// staticcheckで問題が検出される例
//
// staticcheck ./...
// main.go:13:2: this value of err is never used (SA4006)
// main.go:13:8: returnErr doesn't have side effects and its return value is ignored (SA4017)
