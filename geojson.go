package ituzone

// FeatureCollection represents a collection of GeoJSON features.
type FeatureCollection struct {
	Type     string    `json:"type"`
	Features []GeoJSON `json:"features"`
}

// GeoJSON represents a GeoJSON feature with a type, properties, and geometry.
type GeoJSON struct {
	Type       string     `json:"type"`
	Properties Properties `json:"properties"`
	Geometry   Geometry   `json:"geometry"`
}

// Properties is a structure representing various attributes including infos, a number, and a central coordinate.
type Properties struct {
	Names  []string `json:"infos"`
	Number int      `json:"number"`
	Center Coordinate
}

// Geometry defines a GeoJSON geometry object, represented by a type and a nested array of coordinates.
type Geometry struct {
	Type        string        `json:"type"`
	Coordinates [][][]float64 `json:"coordinates"`
}

// zoneToGeoJSON converts a zone struct to its GeoJSON representation.
func zoneToGeoJSON(z zone) GeoJSON {
	coordinates := make([][]float64, len(z.polygon))
	for i, coord := range z.polygon {
		coordinates[i] = []float64{coord.Lng, coord.Lat} // GeoJSON uses [lng, lat]
	}

	// Close the polygon loop if needed
	if len(coordinates) > 0 && (coordinates[0][0] != coordinates[len(coordinates)-1][0] || coordinates[0][1] != coordinates[len(coordinates)-1][1]) {
		coordinates = append(coordinates, coordinates[0])
	}

	return GeoJSON{
		Type: "Feature",
		Properties: Properties{
			Names:  z.infos,
			Number: z.number,
			Center: z.center,
		},
		Geometry: Geometry{
			Type:        "Polygon",
			Coordinates: [][][]float64{coordinates},
		},
	}
}

// ZonesToGeoJSON converts a slice of zones into a GeoJSON FeatureCollection.
func zonesToGeoJSON(zList []zone) FeatureCollection {
	features := make([]GeoJSON, len(zList))
	for i, z := range zList {
		features[i] = zoneToGeoJSON(z)
	}

	featureCollection := FeatureCollection{
		Type:     "FeatureCollection",
		Features: features,
	}

	return featureCollection
}
