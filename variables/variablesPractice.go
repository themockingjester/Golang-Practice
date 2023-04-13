package main

import "fmt"

func main() {
	// Normal Creation
	var var1 int

	// Short hand creation (it don't need datatype if you put you will get error)
	var2 := 2

	// Example for block creation
	var (
		var3 string = "dfgf"
		var4 int    = 34
		var5        = true
	)

	// Example for creating constant (here you can omit datatype also)
	const var6 int = 455

	// Compound Creation
	var7, var8 := 234, 'E'
	fmt.Println(
		var1, // 0 : default value for int is 0
		var2, // 2
		var3, //dfgf
		var4, //34
		var5, //true
		var6, //455
		var7, //234
		var8) //69  : Ascii code for E

	// variable having dynamic datatype
	var var9 any
	fmt.Println(var9) // <nil> : remember you can assign any type of value to this variable
}
