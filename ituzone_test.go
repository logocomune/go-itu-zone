package ituzone

import (
	"reflect"
	"testing"
)

func TestGetZone(t *testing.T) {
	tests := []struct {
		name     string
		coord    Coordinate
		expected ItuZone
		found    bool
	}{
		{
			name:  "Valid coordinate in zone 28",
			coord: Coordinate{Lat: 45.0, Lng: 10.0},
			expected: ItuZone{
				Number:  28,
				Infos:   zones[27].infos,
				Center:  zones[27].center,
				GeoJSON: zoneToGeoJSON(zones[27]),
			},
			found: true,
		},
		{
			name:  "Valid coordinate in zone 47",
			coord: Coordinate{Lat: 15.0, Lng: 25.0},
			expected: ItuZone{
				Number:  47,
				Infos:   zones[46].infos,
				Center:  zones[46].center,
				GeoJSON: zoneToGeoJSON(zones[46]),
			},
			found: true,
		},

		{
			name:     "Invalid latitude",
			coord:    Coordinate{Lat: -95.0, Lng: 10.0},
			expected: ItuZone{},
			found:    false,
		},
		{
			name:     "Invalid longitude",
			coord:    Coordinate{Lat: 45.0, Lng: -190.0},
			expected: ItuZone{},
			found:    false,
		},
		{
			name:     "Max latitude boundary",
			coord:    Coordinate{Lat: 90.0, Lng: 0.0},
			expected: ItuZone{},
			found:    false,
		},
		{
			name:     "Min latitude boundary",
			coord:    Coordinate{Lat: -90.1, Lng: 0.0},
			expected: ItuZone{},
			found:    false,
		},
		{
			name:     "Max longitude boundary",
			coord:    Coordinate{Lat: 0.0, Lng: 180.1},
			expected: ItuZone{},
			found:    false,
		},
		{
			name:     "Min longitude boundary",
			coord:    Coordinate{Lat: 0.0, Lng: -180.1},
			expected: ItuZone{},
			found:    false,
		},
	}

	for _, tt := range tests {
		t.Run(tt.name, func(t *testing.T) {
			result, found := GetZoneByCoordinate(tt.coord)
			if found != tt.found || (found && !reflect.DeepEqual(result, tt.expected)) {
				t.Errorf("GetZoneByCoordinate(%v) = %v, %v; want %v, %v", tt.coord, result, found, tt.expected, tt.found)
			}
		})
	}
}

func TestGetZoneByNumber(t *testing.T) {
	tests := []struct {
		name     string
		id       int
		expected ItuZone
		found    bool
	}{
		{
			name: "Valid zone number 28",
			id:   27,
			expected: ItuZone{
				Number:  28,
				Infos:   zones[27].infos,
				Center:  zones[27].center,
				GeoJSON: zoneToGeoJSON(zones[27]),
			},
			found: true,
		},
		{
			name: "Valid zone number 47",
			id:   46,
			expected: ItuZone{
				Number:  47,
				Infos:   zones[46].infos,
				Center:  zones[46].center,
				GeoJSON: zoneToGeoJSON(zones[46]),
			},
			found: true,
		},
		{
			name:     "Invalid zone number",
			id:       -1,
			expected: ItuZone{},
			found:    false,
		},
		{
			name:     "Zero zone number",
			id:       0,
			expected: ItuZone{},
			found:    false,
		},
		{
			name:     "Zone number out of range",
			id:       91,
			expected: ItuZone{},
			found:    false,
		},
	}

	for _, tt := range tests {
		t.Run(tt.name, func(t *testing.T) {
			result, found := GetZoneByNumber(tt.id)
			if found != tt.found || (found && !reflect.DeepEqual(result, tt.expected)) {
				t.Errorf("GetZoneByNumber(%d) = %v, %v; want %v, %v", tt.id, result, found, tt.expected, tt.found)
			}
		})
	}
}
