// 0から複数のエラーをラップする可能性のあるカスタムエラーの例
package main

import (
	"errors"
	"fmt"
)

type Status int

type StatusErr struct {
	Status  Status
	Message string
}

func (se StatusErr) Error() string {
	return se.Message
}

const (
	InvalidLogin Status = iota + 1
	NotFound
)

type MyError struct {
	Code   int
	Errors []error
}

// Error メソッドは MyError に含まれるエラーを結合し、単一の文字列として返します。
func (m MyError) Error() string {
	return errors.Join(m.Errors...).Error()
}

// Unwrap メソッドは MyError に含まれるエラーをスライスとして返します。
func (m MyError) Unwrap() []error {
	return m.Errors
}

// funcThatReturnsAnError はエラーを返す関数です。
func funcThatReturnsAnError() error {
	return MyError{
		Code: 12,
		Errors: []error{
			StatusErr{
				Status:  NotFound,
				Message: "file Not Found",
			},
			errors.New("a simple string error"),
		},
	}
}

func main() {
	var err error
	err = funcThatReturnsAnError()
	switch err := err.(type) {
	case interface{ Unwrap() error }: // 単発のエラー
		// handle single error
		innerErr := err.Unwrap()
		// process innerErr
		fmt.Println(innerErr)
	case interface{ Unwrap() []error }: // 複数エラー
		//handle multiple wrapped errors
		innerErrs := err.Unwrap()
		for _, innerErr := range innerErrs {
			// process each innerErr
			fmt.Println(innerErr)
		}
	default:
		// handle no wrapped error
	}
}
