// 複数エラーのラップ
package main

import (
	"errors"
	"fmt"
)

type Person struct {
	FirstName string
	LastName  string
	Age       int
}

// ValidatePerson 関数は、Person オブジェクトのフィールド値を検証し、無効な場合はエラーを返します。
// FirstName フィールドが空の場合はエラーを生成します。
// LastName フィールドが空の場合はエラーを生成します。
// Age フィールドが負の値の場合はエラーを生成します。
// エラーが複数存在する場合は、errors.Join を使用して複数のエラーを結合します。
func ValidatePerson(p Person) error {
	var errs []error
	if len(p.FirstName) == 0 {
		errs = append(errs, errors.New("field FirstName cannot be empty"))
	}
	if len(p.LastName) == 0 {
		errs = append(errs, errors.New("field LastName cannot be empty"))
	}
	if p.Age < 0 {
		errs = append(errs, errors.New("field Age cannot be negative"))
	}
	if len(errs) > 0 {
		// errors.Joinで複数のエラーをジョインする。
		return errors.Join(errs...)
	}
	return nil
}

func main() {
	err := ValidatePerson(Person{
		FirstName: "",
		LastName:  "",
		Age:       -1,
	})
	fmt.Println(err)
}
