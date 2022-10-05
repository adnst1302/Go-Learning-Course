package main

import "fmt"

func main() {
	var server1 = getInfoServer("JPY")
	fmt.Println(server1)
}

func getInfoServer(lokasi string) string {
	var namaServer string
	//var lokasiServer string
	if lokasi == "SGP" {
		namaServer = "Singapura"
	} else if lokasi == "HKG" {
		namaServer = "Hongkong"
	} else if lokasi == "JPY" {
		namaServer = "Jepang"
	} else {
		namaServer = "Tidak diketahui"
	}
	return namaServer
}
