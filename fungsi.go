package main

import "fmt"

func main() {
	loadServer("VPS Haxors", "SGP", "10.100.211")
}

func loadServer(nama string, lokasi string, ip string) {
	fmt.Println("Server name : ", nama, " - Lokasi : ", lokasi, " - IP : ", ip)
}
