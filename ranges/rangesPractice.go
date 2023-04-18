package main

// Ranges  used for iterating over elements in various data structures like an array, slice, map, or even a string, etc
import "fmt"

func main() {
	mySlice := []string{"Hi", "There"}

	// irerating through elments of slices
	for index, element := range mySlice {
		fmt.Println(index, " ", element)

		// iterating through strings (present as elements for slice)
		for innerIndex, innerElement := range element {
			fmt.Printf("%v %q", innerIndex, innerElement) // Here we need %q because to show everything as character of string else it will treat every character as rune
			fmt.Println()
		}
	}

	// Output of above program
	// 0   Hi
	// 0 'H'
	// 1 'i'
	// 1   There
	// 0 'T'
	// 1 'h'
	// 2 'e'
	// 3 'r'
	// 4 'e'
}
