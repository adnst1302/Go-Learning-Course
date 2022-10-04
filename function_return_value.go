package main

import "fmt"

func main() {
	infoVpsHaxors := getInfoServer("VPS Haxors", "SGP", "105.122.201.111")
	fmt.Println(infoVpsHaxors)
}

func getInfoServer(nama string, lokasi string, ip string) string {
	return "Server name : " + nama + " - Lokasi : " + lokasi + " - IP : " + ip
}
