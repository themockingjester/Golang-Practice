package main

import "fmt"

func main() {

	// ways of creating arrays
	var arr1 [3]int                     // When we want to assign values later (by default values are defaults as per datatypes)
	arr2 := [...]int{23, 45, 67, 34, 7} // When we don't have idea about array but we have  elements
	arr3 := [3]int{34, 56, 7}           // normal creation
	arr4 := [4]int{1, 56, 3}            // when some indices values are missing ( you can only miss last values first and middle values can't be missed)
	fmt.Println(arr1)                   //[0 0 0]
	fmt.Println(arr2)                   // [23 45 67 34 7]
	fmt.Println(arr3)                   //[34 56 7]
	fmt.Println(arr4)                   //[1 56 3 0]

	for i := 0; i < len(arr2); i++ {
		fmt.Println(arr2[i])
	}

	fmt.Println(arr2[1:3]) // [45 67] : using slicing operation
}
