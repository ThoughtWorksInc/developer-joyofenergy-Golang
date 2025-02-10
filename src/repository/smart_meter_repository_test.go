package repository

import (
	"testing"

	"joyofenergy/src/domain"
)

func TestSmartMeterRepository(t *testing.T) {
	// Create a new repository
	repo := NewSmartMeterRepository()

	// Test case 1: Save and Find a smart meter
	testMeterID := "smart-meter-1"
	testReadings := []*domain.ElectricityReading{
		{
			Time:    1614335400,
			Reading: 35.5,
		},
	}
	testMeter := domain.NewSmartMeter(nil, testReadings)

	// Save the smart meter
	repo.Save(testMeterID, testMeter)

	// Find the smart meter
	foundMeter := repo.FindByID(testMeterID)
	if foundMeter == nil {
		t.Errorf("Failed to find saved smart meter")
	}

	// Verify the readings
	if len(foundMeter.ElectricityReadings) != len(testReadings) {
		t.Errorf("Expected %d readings, got %d", len(testReadings), len(foundMeter.ElectricityReadings))
	}

	// Test case 2: Get all meters
	additionalMeterID := "smart-meter-2"
	additionalReadings := []*domain.ElectricityReading{
		{
			Time:    1614335500,
			Reading: 40.0,
		},
	}
	additionalMeter := domain.NewSmartMeter(nil, additionalReadings)
	repo.Save(additionalMeterID, additionalMeter)

	allMeters := repo.GetAllMeters()
	if len(allMeters) != 2 {
		t.Errorf("Expected 2 meters, got %d", len(allMeters))
	}

	// Test case 3: Find non-existent meter
	nonExistentMeter := repo.FindByID("non-existent-meter")
	if nonExistentMeter != nil {
		t.Errorf("Expected nil for non-existent meter, got a meter")
	}
}
