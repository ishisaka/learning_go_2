package main

import "fmt"

type MailCategory int

const (
	Uncategorized MailCategory = iota
	Personal
	Spam
	Social
	Advertisements
)

func main() {
	fmt.Println(Personal)
}
