package main

import "fmt"

func main() {
	provider := "digital_ocean"
	p1 := "digital_ocean"

	if provider == p1 {
		fmt.Println("Benar")
	} else {
		fmt.Println("Salah")
	}
}
