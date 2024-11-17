# Go ITU Zone Info Package




This repository contains a Go package to retrieve information about ITU Zones (International Telecommunication Union
Zones) using geographic coordinates (latitude and longitude) or the ITU Zone number.

## Features

**Search by coordinates**: Get the ITU Zone corresponding to specific geographic coordinates.

**Search by ITU Zone number**: Retrieve details about an ITU Zone by its identifier number.

## Installation

```console
go get  github.com/logocomune/go-itu-zone
```

## Usage

```golang
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

```

## Acknowledgments

Special thanks to @HB9HIL for providing the GeoJSON files of ITU Zones via their repository: hamradio-zones-geojson.

## License

This project is distributed under the MIT License. See the LICENSE file for more details.