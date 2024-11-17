package ituzone

import (
	"testing"
)

func TestIsPointInZone(t *testing.T) {
	tests := []struct {
		name     string
		polygon  polygon
		point    Coordinate
		expected bool
	}{
		{
			name:     "InsideSimpleTriangle",
			polygon:  polygon{{Lat: 0, Lng: 0}, {Lat: 0, Lng: 10}, {Lat: 10, Lng: 0}},
			point:    Coordinate{Lat: 2, Lng: 2},
			expected: true,
		},
		{
			name:     "OutsideSimpleTriangle",
			polygon:  polygon{{Lat: 0, Lng: 0}, {Lat: 0, Lng: 10}, {Lat: 10, Lng: 0}},
			point:    Coordinate{Lat: 10, Lng: 10},
			expected: false,
		},
		{
			name:     "OnEdgeOfTriangle",
			polygon:  polygon{{Lat: 0, Lng: 0}, {Lat: 0, Lng: 10}, {Lat: 10, Lng: 0}},
			point:    Coordinate{Lat: 5, Lng: 5},
			expected: false,
		},
		{
			name:     "InsideComplexPolygon",
			polygon:  polygon{{Lat: 0, Lng: 0}, {Lat: 0, Lng: 10}, {Lat: 5, Lng: 15}, {Lat: 10, Lng: 10}, {Lat: 10, Lng: 0}},
			point:    Coordinate{Lat: 5, Lng: 5},
			expected: true,
		},
		{
			name:     "OutsideComplexPolygon",
			polygon:  polygon{{Lat: 0, Lng: 0}, {Lat: 0, Lng: 10}, {Lat: 5, Lng: 15}, {Lat: 10, Lng: 10}, {Lat: 10, Lng: 0}},
			point:    Coordinate{Lat: 10, Lng: 12},
			expected: false,
		},
		{
			name:     "OnVertexOfPolygon",
			polygon:  polygon{{Lat: 0, Lng: 0}, {Lat: 0, Lng: 10}, {Lat: 5, Lng: 15}, {Lat: 10, Lng: 10}, {Lat: 10, Lng: 0}},
			point:    Coordinate{Lat: 0, Lng: 0},
			expected: true,
		},
		{
			name:     "DegeneratePolygon",
			polygon:  polygon{{Lat: 0, Lng: 0}, {Lat: 0, Lng: 10}},
			point:    Coordinate{Lat: 5, Lng: 5},
			expected: false,
		},
		{
			name:     "Empty polygon",
			point:    Coordinate{Lat: 5.0, Lng: 5.0},
			polygon:  polygon{},
			expected: false,
		},
		{
			name:     "inside",
			polygon:  polygon{{Lat: 0, Lng: 0}, {Lat: 4, Lng: 0}, {Lat: 4, Lng: 4}, {Lat: 0, Lng: 4}},
			point:    Coordinate{Lat: 2, Lng: 2},
			expected: true,
		},
		{
			name:     "outside",
			polygon:  polygon{{Lat: 0, Lng: 0}, {Lat: 4, Lng: 0}, {Lat: 4, Lng: 4}, {Lat: 0, Lng: 4}},
			point:    Coordinate{Lat: 5, Lng: 5},
			expected: false,
		},
		{
			name:     "on edge",
			polygon:  polygon{{Lat: 0, Lng: 0}, {Lat: 4, Lng: 0}, {Lat: 4, Lng: 4}, {Lat: 0, Lng: 4}},
			point:    Coordinate{Lat: 4, Lng: 2},
			expected: false,
		},
		{
			name:     "on vertex",
			polygon:  polygon{{Lat: 0, Lng: 0}, {Lat: 4, Lng: 0}, {Lat: 4, Lng: 4}, {Lat: 0, Lng: 4}},
			point:    Coordinate{Lat: 4, Lng: 0},
			expected: false,
		},
		{
			name:     "small polygon",
			polygon:  polygon{{Lat: 0, Lng: 0}, {Lat: 1, Lng: 1}},
			point:    Coordinate{Lat: 0.5, Lng: 0.5},
			expected: false,
		},
		{
			name:     "concave inside",
			polygon:  polygon{{Lat: 0, Lng: 0}, {Lat: 4, Lng: 0}, {Lat: 4, Lng: 4}, {Lat: 2, Lng: 2}, {Lat: 0, Lng: 4}},
			point:    Coordinate{Lat: 1, Lng: 1},
			expected: true,
		},
		{
			name:     "concave outside",
			polygon:  polygon{{Lat: 0, Lng: 0}, {Lat: 4, Lng: 0}, {Lat: 4, Lng: 4}, {Lat: 2, Lng: 2}, {Lat: 0, Lng: 4}},
			point:    Coordinate{Lat: 3.0, Lng: 3.2},
			expected: false,
		},
	}

	for _, tt := range tests {
		t.Run(tt.name, func(t *testing.T) {
			result := tt.polygon.isPointInZone(tt.point)
			if result != tt.expected {
				t.Errorf("expected %v, got %v", tt.expected, result)
			}
		})
	}
}
