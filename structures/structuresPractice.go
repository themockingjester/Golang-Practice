package main

import "fmt"

type MyFunction func(int, int) int

// normal structure
type MyStruct struct {
	val1           string
	val2, val3     int
	functionInside MyFunction
}

func main() {

	// First way to initialize structure
	var val6 MyStruct = MyStruct{"Hello", 12, 56, func(val4 int, val5 int) int {
		return val4 * val5
	}}
	// another way to initialize structure
	val7 := MyStruct{functionInside: func(val4 int, val5 int) int {
		return val4 * val5
	}, val1: "Hello", val2: 10, val3: 5}

	_ = val7          // this line is needed because we dont want to use val7 anymore in program if we wont do this then program will throw error
	fmt.Println(val6) //    {Hello 12 56 0x9ddde0}

	fmt.Println(val6.val1)                                 // Hello
	fmt.Println(val6.functionInside(val6.val2, val6.val3)) // 672

	//updating fields inside structure
	val6.val2 = 456
	fmt.Println(val6.functionInside(val6.val2, val6.val3)) // 25536

	// Anonymous Structures
	anyStruct := struct {
		val1 string
		val2 int
	}{
		val1: "There",
		val2: 12,
	}
	fmt.Println(anyStruct) // {There 12}

}
