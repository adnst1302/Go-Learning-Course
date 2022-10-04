package main

import "fmt"

func main() {
	for i := 0; i < 15; i++ {
		fmt.Println("Perulangan ke ", i)
		if i == 5 {
			fmt.Println("Stop")
			break
		}
	}
}
