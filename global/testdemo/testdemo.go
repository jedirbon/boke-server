package global

import (
	"fmt"
	"github.com/oschwald/geoip2-golang"
	"log"
	"net"
)

func GetLocationByIP(ip string) {
	db, err := geoip2.Open("GeoLite2-City.mmdb")
	if err != nil {
		log.Fatal(err)
	}
	defer db.Close()

	ipObj := net.ParseIP(ip)
	record, err := db.City(ipObj)
	if err != nil {
		log.Fatal(err)
	}

	fmt.Printf("国家: %v\n", record.Country.Names["zh-CN"])
	fmt.Printf("城市: %v\n", record.City.Names["zh-CN"])
	fmt.Printf("经纬度: %v, %v\n", record.Location.Latitude, record.Location.Longitude)
}

func main() {
	GetLocationByIP("192.168.157.1")
}
