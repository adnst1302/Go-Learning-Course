package main

import "fmt"

func main() {
	var server1 = getInfoServer("CHR Maxim", "HKG", "192.168.11.22")
	fmt.Println(server1)
}

func getInfoServer(nama string, lokasi string, ip string) string {
	var namaServer string
	//var lokasiServer string
	if lokasi == "SGP" {
		namaServer = "Singapura"
	} else if lokasi == "HKG" {
		namaServer = "Hongkong"
	}
	return namaServer
}
