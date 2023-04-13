package main

import "fmt"

func main() {

	fmt.Println(functionWithSingleReturn(2, 45)) //47
	returnedVal1, returnedVal2 := functionWithMultipleRetun(12, "user")
	fmt.Println(returnedVal1)                // 12
	fmt.Println(returnedVal2)                // Hi user
	fmt.Println(functionReturningAnyThing()) // true
	functionReturningNothing()               // wont return anything
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
