package main

// Pointers are used to update variables out of the function (because function passes parameters as call by value)
import "fmt"

func main() {

	val := 12
	var point *int
	point = &val // creating pointer

	fmt.Println(point) // 0xc000016098 : location of val variable in memory

	fmt.Println(" Value before updating: ", val) // Value before updating:  12
	updateVariable(&val, 45)
	fmt.Println(" Value after updating: ", val) //  Value after updating:  45

}

func updateVariable(oldVal *int, newVal int) {
	*oldVal = newVal
}
