package ituzone

import "log/slog"

type zone struct {
	infos   []string
	number  int
	center  Coordinate
	polygon polygon
}

type polygon []Coordinate

type Coordinate struct {
	Lat float64
	Lng float64
}

func (p polygon) isPointInZone(c Coordinate) bool {

	n := len(p)
	if n < 3 {
		slog.Debug("Polygon is too small")
		return false
	}

	inside := false

	for i, v1 := range p {

		v2 := p[(i+1)%n]

		// Check if the point's lat-coordinate is within the lat-bounds of the edge
		if (v1.Lat > c.Lat) != (v2.Lat > c.Lat) {
			// Calculate the lng-coordinate where the edge intersects the horizontal line at p.Lat
			intersectLng := (v2.Lng-v1.Lng)*(c.Lat-v1.Lat)/(v2.Lat-v1.Lat) + v1.Lng

			// Toggle `inside` if p.Lng is less than the intersection point
			if c.Lng < intersectLng {
				inside = !inside
			}
		}
	}

	return inside
}
