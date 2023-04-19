package main

import "fmt"

func main() {

	const (
		val1 = iota
		val2
		val3
	)

	const (
		val4 = iota
		val5 = iota
	)

	const (
		val6 = iota + 3
		_    // skipping value
		val7
	)
	fmt.Println(val1, val2, val3) // 0 1 2
	fmt.Println(val4, val5)       // 0 1
	fmt.Println(val6, val7)       // 3 5
}
