// 関数を引数として渡す
package main

import (
	"fmt"
	"sort"
)

func main() {
	type Person struct {
		FirstName string
		LastName  string
		Age       int
	}

	people := []Person{
		{"Pat", "Patterson", 37},
		{"Tracy", "Bobdaughter", 23},
		{"Fred", "Fredson", 18},
	}
	fmt.Println(people)

	// sort by last name
	// 匿名関数を引数として渡している
	sort.Slice(people, func(i int, j int) bool {
		return people[i].LastName < people[j].LastName
	})
	fmt.Println(people)

	// sort by age
	sort.Slice(people, func(i int, j int) bool {
		return people[i].Age < people[j].Age
	})
	fmt.Println(people)
}
