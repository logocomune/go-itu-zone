package main

import (
	"fmt"
	"github.com/logocomune/go-itu-zone"
)

func main() {
	ituZone, found := ituzone.GetZoneByCoordinate(ituzone.Coordinate{Lat: 43.71, Lng: 11.75})

	if !found {
		fmt.Println("Zone not found")
	}
	fmt.Println("ITU Zone: ", ituZone.Number)
	fmt.Printf("ITU Zone Infos: %+v\n", ituZone.Infos)
	fmt.Printf("ITU Zone Center: %+v\n", ituZone.Center)
	fmt.Printf("ITU Zone GeoJSON: %+v\n", ituZone.GeoJSON)
}
