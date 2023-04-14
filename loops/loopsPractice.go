package main

import "fmt"

func main() {

	// for loop : as you can see we have not wrapped conditons in braces because golang wont allow this
	for i := 0; i < 10; i++ {
		fmt.Println(i)
	}

	var val int = 0

	// while loop: in golang we don't have a while loop that can be easily implemented by for loops
	for val <= 5 {
		fmt.Println(val)
		val++
	}

	// infnite loop
	for {
		fmt.Println("Infinite loop")
	}

}
