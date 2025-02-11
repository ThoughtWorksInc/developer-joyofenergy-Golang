package tests

import (
	"testing"

	"joyofenergy/src/meters"
)

func TestNewMeter(t *testing.T) {
	meter := meters.NewMeter("test-meter", "electric")

	if meter.ID != "test-meter" {
		t.Errorf("Expected meter ID to be 'test-meter', got %s", meter.ID)
	}

	if meter.Type != "electric" {
		t.Errorf("Expected meter type to be 'electric', got %s", meter.Type)
	}
}

func TestMeterAddReading(t *testing.T) {
	meter := meters.NewMeter("test-meter", "electric")
	reading := meters.MeterReading{Time: "2023-01-01T00:00:00Z", Value: 10.0}

	meter.AddReading(reading)

	if len(meter.Readings) != 1 {
		t.Errorf("Expected 1 reading, got %d", len(meter.Readings))
	}

	if meter.Readings[0].Value != 10.0 {
		t.Errorf("Expected reading value of 10.0, got %f", meter.Readings[0].Value)
	}
}
