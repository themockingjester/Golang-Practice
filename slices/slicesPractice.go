package main

import "fmt"

func main() {
	mySlice := []int{1, 2, 3, 4, 5}
	fmt.Println(mySlice)      // [1 2 3 4 5]
	fmt.Println(mySlice[:])   // [1 2 3 4 5]
	fmt.Println(mySlice[1:3]) // [2 3]

	newSlice := append(mySlice, 67, 78, 64) // [1 2 3 4 5 67 78 64]

	thirdSlice := append(mySlice, newSlice...) //[1 2 3 4 5 1 2 3 4 5 67 78 64]   : Joining two slices
	fmt.Println(newSlice)
	fmt.Println(thirdSlice)

	preAllocatedSlice := make([]int, 10) // this is preallocation of slice it is useful when we have total number of elements but don't have their values
	fmt.Println(preAllocatedSlice)       // [0 0 0 0 0 0 0 0 0 0]

	iterate(thirdSlice) // 1234512345677864
	fmt.Println()
	iterate(newSlice) // 12345677864
	fmt.Println()

	// Below is the case of multidimentional array , remember you are only allowed to keep values according to datatype specified (for full slice)
	multiDimesionalSlice := [][]int{

		// type declartion is optional
		[]int{12, 23, 45},
		{45, 67, 89},
	}

	fmt.Println(multiDimesionalSlice[0][2]) // 45
}

// see below function can iterate over slice of any size
func iterate(slice []int) {
	for i := 0; i < len(slice); i++ {
		fmt.Print(slice[i])

	}
}
