package main

import "fmt"

func main() {
	jumlahkan := sumAll(1, 1, 23, 12)
	fmt.Println(jumlahkan)
}

func sumAll(angka ...int) int {
	total := 0

	for _, value := range angka {
		total += value
	}
	return total
}
