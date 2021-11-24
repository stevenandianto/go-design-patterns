package main

import "fmt"

func main() {
	// Buy from nike
	nikeFactory, _ := getSportsFactory("nike")
	nikeShoe := nikeFactory.makeShoe()
	nikeShirt := nikeFactory.makeShirt()

	// Buy from adidas
	adidasFactory, _ := getSportsFactory("adidas")
	adidasShoe := adidasFactory.makeShoe()
	adidasShirt := adidasFactory.makeShirt()

	fmt.Printf("Alfi love Nike")
	println()
	printShoeDetails(nikeShoe)
	printShirtDetails(nikeShirt)

	fmt.Printf("Okta love Adidas")
	println()
	printShoeDetails(adidasShoe)
	printShirtDetails(adidasShirt)
}

func printShoeDetails(s iShoe) {
	fmt.Printf("Shoe Logo: %s", s.getLogo())
	fmt.Println()
	fmt.Printf("Shoe Size: %d", s.getSize())
	fmt.Println()
}

func printShirtDetails(s iShirt) {
	fmt.Printf("Shirt Logo: %s", s.getLogo())
	fmt.Println()
	fmt.Printf("Shirt Size: %d", s.getSize())
	fmt.Println()
}
