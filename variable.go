package main

import "fmt"

func main() {
	//dengan kata kunci tipe data
	var name string
	name = "Aditia Darma Nst"

	// deklarasi langsung
	alamat := "Jakarta Barat"

	// multiple variable
	var (
		sd  = "SDN 101938 Adolina"
		smp = "SMPN 1 Perbaungan"
	)

	fmt.Println(name)
	fmt.Println(alamat)
	fmt.Println(sd)
	fmt.Println(smp)
}
