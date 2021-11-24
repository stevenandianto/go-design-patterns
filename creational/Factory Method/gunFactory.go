// Simple Factory
package main

import "fmt"

func getGun(gunType string) (iGun, error) {
	if gunType == "ak47" {
		return newAk47(), nil
	}
	if gunType == "m4a1" {
		return newM4a1(), nil
	}
	return nil, fmt.Errorf("Wrong gun type passed")
}
