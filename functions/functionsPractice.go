package main

import "fmt"

func main() {

	fmt.Println(functionWithSingleReturn(2, 45)) //47
	returnedVal1, returnedVal2 := functionWithMultipleRetun(12, "user")
	fmt.Println(returnedVal1)                // 12
	fmt.Println(returnedVal2)                // Hi user
	fmt.Println(functionReturningAnyThing()) // true
	functionReturningNothing()               // wont return anything

	fmt.Println(addNumbers(3, 45, 67, 7)) // this function can acccept any number of integers
}
func functionWithMultipleRetun(val1 int, val2 string) (int, string) {
	return val1, "Hi " + val2
}
func functionWithSingleReturn(val1 int, val2 int) int {
	return val1 + val2
}

func functionReturningAnyThing() any {
	fmt.Println("This function can return any thing based on the return statement")
	return true
}

func functionReturningNothing() {
	fmt.Println(`This function wont return anything (if you try to return you will get an error)`)
}

func addNumbers(nums ...int) int {
	sum := 0
	for index, num := range nums {
		fmt.Println(index, num)
		sum += num
	}
	// above loops prints below on screen
	// 	0 3
	// 1 45
	// 2 67
	// 3 7
	return sum
}
