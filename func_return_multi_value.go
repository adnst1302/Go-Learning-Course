package main

import "fmt"

func main() {
	fNilai1, fNilai2 := getInfoVps()
	fmt.Println(fNilai2, fNilai1)
}

func getInfoVps() (string, string) {
	return "Nilai1", "Nilai 2"
}
