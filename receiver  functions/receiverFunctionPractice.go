package main

import "fmt"

type Coordinate struct {
	x int
	y int
}

func main() {
	currCoordinate := Coordinate{x: 12, y: 15}
	fmt.Println(currCoordinate) // {12 15}
	currCoordinate.increaseAxisByNum(1, 3)
	fmt.Println(currCoordinate) // {13 18}
	currCoordinate.uselessFunction(2, 45)
	fmt.Println(currCoordinate) // {13 18}
}

func (coor *Coordinate) increaseAxisByNum(x int, y int) {
	// This is is pointer receiver function which has capability to modify the type provided
	coor.x += x
	coor.y += y
}

func (coor Coordinate) uselessFunction(x int, y int) {
	// This is is value receiver function which won't be  able to modify the type provided

	coor.x += x
	coor.y += y
}
