package main

import "fmt"

func main() {
	true := false
	fmt.Println(true)
}

/*
reviveで問題が検出されるケース
$ revive -config built_in.toml ./...
main.go:6:2: assignment creates a shadow of built-in identifier true
*/
