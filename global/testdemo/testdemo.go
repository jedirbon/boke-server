package main

import (
	"boke-server/core"
	"fmt"
)

func main() {
	core.InitIPDB()
	addr := core.GetIpAddr("175.0.201.207")
	fmt.Println(addr)
}
