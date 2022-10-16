package main

import "fmt"

func main() {
	type NoKTP string
	type Married bool

	var noKTPEko NoKTP = "1874930493049040300540"
	var marriedStatus Married = true
	fmt.Println(noKTPEko)
	fmt.Println(marriedStatus)
}
