package main

import "fmt"

func main() {
	emptyMap := make(map[string]int)

	filledMap := map[string]int{
		"key1": 1,
		"key2": 2,
	}

	fmt.Println(emptyMap)  // map[]
	fmt.Println(filledMap) //map[key1:1 key2:2]

	filledMap["key3"] = 3  // adding data to map
	fmt.Println(filledMap) // map[key1:1 key2:2 key3:3]

	delete(filledMap, "key2") // delete
	fmt.Println(filledMap)    // map[key1:1 key3:3]

	fmt.Println(filledMap["key3"]) // 3  :   reading values from map

	// checking key existence in map
	val1, found1 := filledMap["key2"] // 0 false
	fmt.Println(val1, found1)

	val2, found2 := filledMap["key3"] // 3 true
	fmt.Println(val2, found2)

	// iterating through map
	for key, val := range filledMap {
		fmt.Println(key, val)
	}

	// output for above loop
	// 	key1 1
	// key3 3

}
