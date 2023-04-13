package main

import "fmt"

func main() {
	val := 534
	// Note: if switch is applying on x datatype (for below case its val - int datatype ) value then all cases must have values of same datatype else you will get error
	switch val {
	case 23:
		{
			fmt.Println("First Block")
		}
	case 234, 534, 56:
		fmt.Println("Second Block")
		// Here we have not closed this case in braces because we have fallthrough here
		fallthrough // we need fallthrough when we want to execute next case (no matter condition satifies or not)
	case 78, 34:
		fmt.Println("Third Block")
		fallthrough
	case 5412:
		{
			fmt.Println("Fourth Block")
		}
	case 67434554:
		{
			fmt.Println("Fifth Block")
		}
	default:
		{
			fmt.Println("Sixth Block")
		}
	}

	// above code outputs
	// 	Second Block
	// Third Block
	// Fourth Block
}
