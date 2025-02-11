package tests

import (
	"testing"

	"joyofenergy/src/meters"
	"joyofenergy/src/readings"
	"joyofenergy/src/usage"
)

func TestUsageService(t *testing.T) {
	meterService := meters.NewMeterService()
	meter := meters.NewMeter("test-meter", "electric")
	meterService.AddMeter(meter)

	readingService := readings.NewReadingService(meterService)
	usageService := usage.NewUsageService(readingService)

	err := readingService.StoreReading("test-meter", 10.0)
	if err != nil {
		t.Fatalf("Unexpected error storing reading: %v", err)
	}

	err = readingService.StoreReading("test-meter", 20.0)
	if err != nil {
		t.Fatalf("Unexpected error storing reading: %v", err)
	}

	totalUsage, err := usageService.GetTotalUsage("test-meter")
	if err != nil {
		t.Fatalf("Unexpected error getting total usage: %v", err)
	}

	if totalUsage != 30.0 {
		t.Errorf("Expected total usage of 30.0, got %f", totalUsage)
	}

	// Test getting total usage for non-existent meter
	_, err = usageService.GetTotalUsage("non-existent")
	if err == nil {
		t.Error("Expected error when getting total usage for non-existent meter")
	}
}

func TestUsageServiceWithNoReadings(t *testing.T) {
	meterService := meters.NewMeterService()
	meter := meters.NewMeter("test-meter", "electric")
	meterService.AddMeter(meter)

	readingService := readings.NewReadingService(meterService)
	usageService := usage.NewUsageService(readingService)

	totalUsage, err := usageService.GetTotalUsage("test-meter")
	if err != nil {
		t.Fatalf("Unexpected error getting total usage: %v", err)
	}

	if totalUsage != 0.0 {
		t.Errorf("Expected total usage of 0.0 for meter with no readings, got %f", totalUsage)
	}
}
