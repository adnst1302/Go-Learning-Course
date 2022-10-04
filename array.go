package main

import "fmt"

func main() {
	var region [4]string
	region[0] = "SGP"
	region[1] = "JPY"
	region[2] = "JKT"
	region[3] = "AUS"
	//cetak semua
	fmt.Println(region)
	//cetak index ke
	fmt.Println(region[2])
}
