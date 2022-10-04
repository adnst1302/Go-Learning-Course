package main

import "fmt"

func main() {
	var region = [...]string{
		"SGP",
		"JKT",
		"NDW",
		"JPY",
		"KLM",
		"KNO",
		"AUS",
		"HKG",
	}
	var slice1 = region[3:6]
	fmt.Println(slice1)
	fmt.Println(cap(slice1))
}
