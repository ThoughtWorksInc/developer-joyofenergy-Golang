package tests

import (
	"testing"

	"joyofenergy/src/meters"
	"joyofenergy/src/readings"
)

func TestReadingService(t *testing.T) {
	meterService := meters.NewMeterService()
	meter := meters.NewMeter("test-meter", "electric")
	meterService.AddMeter(meter)

	readingService := readings.NewReadingService(meterService)

	err := readingService.StoreReading("test-meter", 10.0)
	if err != nil {
		t.Fatalf("Unexpected error storing reading: %v", err)
	}

	readings, err := readingService.GetReadings("test-meter")
	if err != nil {
		t.Fatalf("Unexpected error getting readings: %v", err)
	}

	if len(readings) != 1 {
		t.Errorf("Expected 1 reading, got %d", len(readings))
	}

	if readings[0].Value != 10.0 {
		t.Errorf("Expected reading value of 10.0, got %f", readings[0].Value)
	}

	// Test storing reading for non-existent meter
	err = readingService.StoreReading("non-existent", 15.0)
	if err == nil {
		t.Error("Expected error when storing reading for non-existent meter")
	}
}

func TestGetReadingsForNonExistentMeter(t *testing.T) {
	meterService := meters.NewMeterService()
	readingService := readings.NewReadingService(meterService)

	_, err := readingService.GetReadings("non-existent")
	if err == nil {
		t.Error("Expected error when getting readings for non-existent meter")
	}
}
