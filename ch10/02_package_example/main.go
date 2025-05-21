package main

import (
	"fmt"

	"exampla/package_example/do-format"
	"exampla/package_example/math"
)

func main() {
	num := math.Double(2)
	output := format.Number(num)
	fmt.Println(output)
}
