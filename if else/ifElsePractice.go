package main

import "fmt"

func main() {
	val1 := 32
	if val1 >= 10 && val1 <= 30 {
		fmt.Println("First Block")
	} else if val1 >= 31 {
		fmt.Println("Second Block")
	} else {
		fmt.Println("Third Block")
	}

	// above program outputs `Second Block`
}
