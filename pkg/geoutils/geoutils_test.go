package geoutils

import (
	"testing"
)

func TestAntipode(t *testing.T) {
	tests := []struct {
		lat, lon                 float64
		expectedLat, expectedLon float64
	}{
		{0, 0, 0, 180},      // Экватор, нулевой меридиан
		{45, 90, -45, -90},  // Северное полушарие, восточная долгота
		{-30, -60, 30, 120}, // Южное полушарие, западная долгота
		{90, 180, -90, 0},   // Северный полюс
		{-90, -180, 90, 0},  // Южный полюс
	}

	for _, tt := range tests {
		lat, lon := Antipode(tt.lat, tt.lon)
		if lat != tt.expectedLat || lon != tt.expectedLon {
			t.Errorf("Antipode(%v, %v) = (%v, %v), expected (%v, %v)",
				tt.lat, tt.lon, lat, lon, tt.expectedLat, tt.expectedLon)
		}
	}
}
