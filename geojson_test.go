package ituzone

import (
	"encoding/json"
	"reflect"
	"testing"
)

// TestZonesToGeoJSON tests the zonesToGeoJSON function
func TestZonesToGeoJSON(t *testing.T) {
	// Arrange: Define test zones
	zonesTest := []zone{
		{
			infos:  []string{"Zone A", "Zone 1"},
			number: 123,
			center: Coordinate{Lat: 37.7749, Lng: -122.4194},
			polygon: polygon{
				{Lat: 37.7749, Lng: -122.4194},
				{Lat: 37.7750, Lng: -122.4180},
				{Lat: 37.7740, Lng: -122.4170},
				{Lat: 37.7749, Lng: -122.4194},
			},
		},
		{
			infos:  []string{"Zone B", "Zone 2"},
			number: 456,
			center: Coordinate{Lat: 40.7128, Lng: -74.0060},
			polygon: polygon{
				{Lat: 40.7128, Lng: -74.0060},
				{Lat: 40.7130, Lng: -74.0050},
				{Lat: 40.7115, Lng: -74.0045},
				{Lat: 40.7128, Lng: -74.0060},
			},
		},
	}

	// Act: Generate GeoJSON
	actual := zonesToGeoJSON(zonesTest)

	// Assert: Check FeatureCollection properties
	if actual.Type != "FeatureCollection" {
		t.Errorf("Expected Type to be 'FeatureCollection', got '%s'", actual.Type)
	}

	// Assert: Check number of features
	expectedFeatureCount := len(zonesTest)
	if len(actual.Features) != expectedFeatureCount {
		t.Errorf("Expected %d features, got %d", expectedFeatureCount, len(actual.Features))
	}

	// Assert: Check the geometry and properties of each feature
	expectedFeatures := []GeoJSON{
		{
			Type: "Feature",
			Properties: Properties{
				Names:  []string{"Zone A", "Zone 1"},
				Number: 123,
				Center: Coordinate{Lat: 37.7749, Lng: -122.4194},
			},
			Geometry: Geometry{
				Type: "Polygon",
				Coordinates: [][][]float64{
					{
						{-122.4194, 37.7749},
						{-122.418, 37.775},
						{-122.417, 37.774},
						{-122.4194, 37.7749},
					},
				},
			},
		},
		{
			Type: "Feature",
			Properties: Properties{
				Names:  []string{"Zone B", "Zone 2"},
				Number: 456,
				Center: Coordinate{Lat: 40.7128, Lng: -74.0060},
			},
			Geometry: Geometry{
				Type: "Polygon",
				Coordinates: [][][]float64{
					{
						{-74.006, 40.7128},
						{-74.005, 40.713},
						{-74.0045, 40.7115},
						{-74.006, 40.7128},
					},
				},
			},
		},
	}

	for i, feature := range actual.Features {
		expected := expectedFeatures[i]

		if feature.Type != expected.Type {
			t.Errorf("Feature %d: Expected Type '%s', got '%s'", i, expected.Type, feature.Type)
		}

		if !reflect.DeepEqual(feature.Properties, expected.Properties) {
			t.Errorf("Feature %d: Expected Properties %+v, got %+v", i, expected.Properties, feature.Properties)
		}

		if !reflect.DeepEqual(feature.Geometry, expected.Geometry) {
			expectedJSON, _ := json.Marshal(expected.Geometry)
			actualJSON, _ := json.Marshal(feature.Geometry)
			t.Errorf("Feature %d: Expected Geometry %s, got %s", i, expectedJSON, actualJSON)
		}
	}
}
