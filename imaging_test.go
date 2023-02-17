package imaging

import (
	"image/color"
	"testing"
)

func TestGenerateZones(t *testing.T) {
	width := 100
	zones := generateZones(width)
	if len(zones) != len(zoneColors) {
		t.Errorf("expected %d zones, got %d", len(zoneColors), len(zones))
	}
	if expected := width / (len(zoneColors) + 1); zones[0] != expected {
		t.Errorf("expected first zone to be %d, got %d", expected, zones[0])
	}
}

func TestColorForZone(t *testing.T) {
	width := 100
	zones := generateZones(width)

	tests := []struct {
		position int
		expected color.RGBA
	}{
		{0, zoneColors[0]},
		{1, zoneColors[0]},
		{width - 1, zoneColors[len(zoneColors)-1]},
	}

	// test each zone, make sure we have len(zoneColors) different colors
	for i := 0; i < len(zoneColors); i++ {
		increment := width / (len(zoneColors) + 1)
		pos := int((float64(i) + 0.5) * float64(increment))
		tests = append(tests, struct {
			position int
			expected color.RGBA
		}{position: pos, expected: zoneColors[i]})
	}

	for _, test := range tests {
		if colorForZone(test.position, zones) != test.expected {
			t.Errorf("expected color %v for position %d, got %v", test.expected, test.position, colorForZone(test.position, zones))
		}
	}
}
