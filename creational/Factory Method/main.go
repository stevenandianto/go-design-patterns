// Client code
package main

import "fmt"

func main() {
	ak47, _ := getGun("ak47")
	m4a1, _ := getGun("m4a1")

	printDetails(ak47)
	printDetails(m4a1)
}

func printDetails(g iGun) {
	fmt.Printf("Gun: %s", g.getName())
	fmt.Println()
	fmt.Printf("Power: %d", g.getPower())
	fmt.Println()
}
