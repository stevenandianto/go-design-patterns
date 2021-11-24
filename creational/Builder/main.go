package main

import "fmt"

func main() {
	// Without director
	normal := newNormalBuilder()
	normal.setWindowType()
	normal.setDoorType()
	house := normal.getHouse()

	fmt.Printf("Without Director\n")
	fmt.Printf("Normal House Door Type: %s\n", house.doorType)
	fmt.Printf("Normal House Window Type: %s\n", house.windowType)
	fmt.Printf("Normal House Num Floor: %d\n", house.floor)

	// Get Builder
	normalBuilder := getBuilder("normal")
	iglooBuilder := getBuilder("igloo")

	// Assign to director
	director := newDirector(normalBuilder)
	normalHouse := director.buildHouse()

	fmt.Printf("\nWith Director\n")
	fmt.Printf("Normal House Door Type: %s\n", normalHouse.doorType)
	fmt.Printf("Normal House Window Type: %s\n", normalHouse.windowType)
	fmt.Printf("Normal House Num Floor: %d\n", normalHouse.floor)

	// Assign to director
	director.setBuilder(iglooBuilder)
	iglooHouse := director.buildHouse()

	fmt.Printf("\nIgloo House Door Type: %s\n", iglooHouse.doorType)
	fmt.Printf("Igloo House Window Type: %s\n", iglooHouse.windowType)
	fmt.Printf("Igloo House Num Floor: %d\n", iglooHouse.floor)

}
