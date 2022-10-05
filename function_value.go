package main

import "fmt"

func main() {
	vpsAdit := getVpsInfo("127.0.0.1")
	fmt.Println(vpsAdit)
}

func getVpsInfo(ip string) string {
	if ip == "127.0.0.1" {
		return "lokal"
	} else {
		return "luar"
	}
}
