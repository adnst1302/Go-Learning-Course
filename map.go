package main

import "fmt"

func main() {
	var server map[string]string = map[string]string{
		"nama":    "vpshaxors",
		"ip":      "192.168.22.13",
		"os":      "Ubuntu 20.22",
		"storage": "1 Tb",
	}

	//tambahkan key baru
	server["ram"] = "4 Gb"

	//hapus key
	delete(server, "storage")

	fmt.Println(server["nama"])
	fmt.Println(server["ram"])
	fmt.Println(server)
}
