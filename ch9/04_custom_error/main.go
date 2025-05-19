// エラーと値
package main

import (
	"errors"
	"fmt"
	"os"
)

// Status は状態を表現するための列挙型として使用されます。
// 主にエラーや状態管理の識別に利用されます。
// 特定の状態値は別途定義され、関連する処理で使用されます。
type Status int

const (
	InvalidLogin Status = iota + 1

	// NotFound は対象のリソースが見つからない場合を表す状態を示します。
	NotFound
)

// StatusErr はエラーの状態とメッセージを保持するための構造体です。
// エラーの詳細情報を表現する際に使用されます。
type StatusErr struct {
	Status  Status
	Message string
}

// Error は StatusErr 内に格納されているエラーメッセージを文字列として返します。
func (se StatusErr) Error() string {
	return se.Message
}

// LoginAndGetData はユーザーIDとパスワードを使用してログインを試行し、指定されたファイルのデータを取得します。
// ログインが失敗した場合は InvalidLogin エラーを返します。
// ファイルが見つからない場合は NotFound エラーを返します。
// 成功した場合はファイルデータをバイトスライスとして返します。
func LoginAndGetData(uid, pwd, file string) ([]byte, error) {
	token, err := login(uid, pwd)
	if err != nil {
		return nil, StatusErr{
			Status:  InvalidLogin,
			Message: fmt.Sprintf("invalid credentials for user %s", uid),
		}
	}
	data, err := getData(token, file)
	if err != nil {
		return nil, StatusErr{
			Status:  NotFound,
			Message: fmt.Sprintf("file %s not found", file),
		}
	}
	return data, nil
}

// login は指定されたユーザーIDとパスワードを使用してログインを試行します。
// 認証に成功した場合はトークンを返し、失敗した場合はエラーを返します。
func login(uid, pwd string) (string, error) {
	// world's worst login system
	if uid == "admin" && pwd == "admin" {
		return "user:admin", nil
	}
	return "", errors.New("bad user")
}

// getData はトークンとファイル名を受け取り、アクセス可能な場合はそのファイルのデータをバイトスライスとして返します。
// 指定されたファイルが存在しない場合、またはアクセス権がない場合はエラーを返します。
func getData(token, file string) ([]byte, error) {
	// world's worst access control
	if token == "user:admin" {
		switch file {
		case "secrets.txt":
			return []byte("passwords aplenty!"), nil
		case "payroll.csv":
			return []byte("everyone's salary"), nil
		}
	}
	return nil, os.ErrNotExist
}

func main() {
	data, err := LoginAndGetData("admin", "admin", "secrets.txt")
	fmt.Println(string(data), err)

	data, err = LoginAndGetData("admin", "admin", "chicken_recipe.txt")
	fmt.Println(string(data), err)

	data, err = LoginAndGetData("jon", "password", "secrets.txt")
	fmt.Println(string(data), err)
}
