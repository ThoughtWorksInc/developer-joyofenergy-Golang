package tests

import (
	"fmt"
	"testing"
	"time"

	"joyofenergy/src/meters"
)

func TestStoreAndFetchReadings(t *testing.T) {
	meterService := meters.NewMeterService()
	meter := meters.NewMeter("testMeterId", "electric")
	meterService.AddMeter(meter)

	readingService := &mockReadingService{
		meterService: meterService,
	}

	// Store readings matching JavaScript structure
	readings := []map[string]interface{}{
		{"time": float64(1), "reading": 1.0},
		{"time": float64(2), "reading": 2.0},
	}

	err := readingService.StoreReadings("testMeterId", readings)
	if err != nil {
		t.Fatalf("Unexpected error storing readings: %v", err)
	}

	// Retrieve and verify readings
	storedMeter, err := meterService.GetMeterByID("testMeterId")
	if err != nil {
		t.Fatalf("Unexpected error getting meter: %v", err)
	}

	if len(storedMeter.Readings) != 2 {
		t.Errorf("Expected 2 readings, got %d", len(storedMeter.Readings))
	}
}

func TestInvalidReadings(t *testing.T) {
	testCases := []struct {
		name        string
		readings    []map[string]interface{}
		expectError bool
	}{
		{
			name:        "Nil readings",
			readings:    nil,
			expectError: true,
		},
		{
			name:        "Empty readings slice",
			readings:    []map[string]interface{}{},
			expectError: true,
		},
		{
			name: "Invalid time (not a number)",
			readings: []map[string]interface{}{
				{"time": "invalid", "reading": 1.0},
			},
			expectError: true,
		},
		{
			name: "Invalid reading (not a number)",
			readings: []map[string]interface{}{
				{"time": 1, "reading": "invalid"},
			},
			expectError: true,
		},
		{
			name: "Time less than or equal to 0",
			readings: []map[string]interface{}{
				{"time": 0, "reading": 1.0},
			},
			expectError: true,
		},
		{
			name: "Reading less than or equal to 0",
			readings: []map[string]interface{}{
				{"time": 1, "reading": 0},
			},
			expectError: true,
		},
	}

	for _, tc := range testCases {
		t.Run(tc.name, func(t *testing.T) {
			meterService := meters.NewMeterService()
			meter := meters.NewMeter("testMeterId", "electric")
			meterService.AddMeter(meter)

			readingService := &mockReadingService{
				meterService: meterService,
			}

			err := readingService.StoreReadings("testMeterId", tc.readings)

			if tc.expectError && err == nil {
				t.Errorf("Expected error for %s, got nil", tc.name)
			} else if !tc.expectError && err != nil {
				t.Errorf("Unexpected error for %s: %v", tc.name, err)
			}
		})
	}
}

// mockReadingService mimics the behavior of the readings.ReadingService
type mockReadingService struct {
	meterService *meters.MeterService
}

func (mrs *mockReadingService) StoreReadings(meterID string, readings []map[string]interface{}) error {
	if !readingsValid(readings) {
		return fmt.Errorf("Invalid readings")
	}

	meter, err := mrs.meterService.GetMeterByID(meterID)
	if err != nil {
		return err
	}

	// Convert readings to MeterReading
	var meterReadings []meters.MeterReading
	for _, reading := range readings {
		meterReadings = append(meterReadings, meters.MeterReading{
			Time:  time.Unix(int64(reading["time"].(float64)), 0).UTC().Format(time.RFC3339),
			Value: reading["reading"].(float64),
		})
	}

	meter.Readings = meterReadings
	return nil
}

func readingsValid(readings []map[string]interface{}) bool {
	if readings == nil || len(readings) == 0 {
		return false
	}

	for _, reading := range readings {
		if !isValidReading(reading) {
			return false
		}
	}

	return true
}

func isValidReading(reading map[string]interface{}) bool {
	return isValidTime(reading["time"]) && isValidValue(reading["reading"])
}

func isValidTime(timeVal interface{}) bool {
	switch t := timeVal.(type) {
	case float64:
		return t > 0 && t < float64(time.Now().Unix())
	default:
		return false
	}
}

func isValidValue(value interface{}) bool {
	switch v := value.(type) {
	case float64:
		return v > 0
	default:
		return false
	}
}
