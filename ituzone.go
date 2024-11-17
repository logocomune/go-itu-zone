package ituzone

type ItuZone struct {
	Number  int
	Infos   []string
	Center  Coordinate
	GeoJSON GeoJSON
}

func GetZoneByCoordinate(c Coordinate) (ItuZone, bool) {
	if c.Lat < -90 || c.Lat > 90 || c.Lng < -180 || c.Lng > 180 {
		return ItuZone{}, false
	}

	for i := range zones {
		if zones[i].polygon.isPointInZone(c) {
			return zoneToItuZone(zones[i]), true
		}
	}

	return ItuZone{}, false
}

func GetZoneByNumber(id int) (ItuZone, bool) {
	if id < 1 || id > len(zones) {
		return ItuZone{}, false
	}

	return zoneToItuZone(zones[id]), true

}

func GetGoeJson() FeatureCollection {
	return zonesToGeoJSON(zones)
}

func zoneToItuZone(z zone) ItuZone {
	return ItuZone{
		Number:  z.number,
		Infos:   z.infos,
		Center:  z.center,
		GeoJSON: zoneToGeoJSON(z),
	}
}
