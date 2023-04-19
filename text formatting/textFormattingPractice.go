package main

import "fmt"

func main() {

	fmt.Printf("Hi %v, you are from %v", "Yash", "India") // Hi Yash, you are from India
	fmt.Println()
	fmt.Printf("This is a rune %c", 'q') // This is a rune q
	fmt.Println()
	fmt.Printf("Our number is %v", 44) // Our number is 44
	fmt.Println()
	fmt.Println(returnFormattedString("Yash", "India")) // Hi Again this is Yash From India
}

func returnFormattedString(name string, location string) string {
	// Sprint variants return strings rather than printing to console
	return fmt.Sprintf("Hi Again this is %v From %v", name, location)
}
