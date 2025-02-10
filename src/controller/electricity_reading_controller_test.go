package controller

import (
	"bytes"
	"encoding/json"
	"net/http"
	"net/http/httptest"
	"testing"

	"joyofenergy/src/repository"
)

func TestStoreReading(t *testing.T) {
	// Create a mock smart meter repository
	smartMeterRepo := repository.NewSmartMeterRepository()

	// Create a meter reading manager with the mock repository
	meterReadingManager := NewMeterReadingManager(smartMeterRepo)

	// Create the electricity reading controller
	electricityReadingController := NewElectricityReadingController(meterReadingManager)

	// Test cases
	testCases := []struct {
		name           string
		inputJSON      string
		expectedStatus int
	}{
		{
			name: "Valid Reading",
			inputJSON: `{
				"smartMeterId": "smart-meter-1",
				"electricityReadings": [
					{
						"time": 1614335400,
						"reading": 35.5
					}
				]
			}`,
			expectedStatus: http.StatusCreated,
		},
		{
			name: "Missing Smart Meter ID",
			inputJSON: `{
				"electricityReadings": [
					{
						"time": 1614335400,
						"reading": 35.5
					}
				]
			}`,
			expectedStatus: http.StatusBadRequest,
		},
		{
			name: "Missing Electricity Readings",
			inputJSON: `{
				"smartMeterId": "smart-meter-1"
			}`,
			expectedStatus: http.StatusBadRequest,
		},
	}

	for _, tc := range testCases {
		t.Run(tc.name, func(t *testing.T) {
			// Create a request body
			body := bytes.NewBufferString(tc.inputJSON)

			// Create a request
			req, err := http.NewRequest(http.MethodPost, "/readings", body)
			if err != nil {
				t.Fatalf("Failed to create request: %v", err)
			}

			// Create a response recorder
			w := httptest.NewRecorder()

			// Call the handler
			electricityReadingController.StoreReading(w, req)

			// Check the status code
			if w.Code != tc.expectedStatus {
				t.Errorf("Expected status %d, got %d", tc.expectedStatus, w.Code)
			}
		})
	}
}

func TestGetReadings(t *testing.T) {
	// Create a mock smart meter repository
	smartMeterRepo := repository.NewSmartMeterRepository()

	// Create a meter reading manager with the mock repository
	meterReadingManager := NewMeterReadingManager(smartMeterRepo)

	// Create the electricity reading controller
	electricityReadingController := NewElectricityReadingController(meterReadingManager)

	// Prepare test data
	testMeterID := "smart-meter-test"
	testReadings := []map[string]interface{}{
		{
			"smartMeterId": testMeterID,
			"electricityReadings": []map[string]interface{}{
				{
					"time":    1614335400,
					"reading": 35.5,
				},
			},
		},
	}

	// Store test readings
	for _, reading := range testReadings {
		err := meterReadingManager.StoreReading(reading)
		if err != nil {
			t.Fatalf("Failed to store test reading: %v", err)
		}
	}

	// Create a request
	req, err := http.NewRequest(http.MethodGet, "/readings/"+testMeterID, nil)
	if err != nil {
		t.Fatalf("Failed to create request: %v", err)
	}

	// Create a response recorder
	w := httptest.NewRecorder()

	// Call the handler
	electricityReadingController.GetReadings(w, req)

	// Check the status code
	if w.Code != http.StatusOK {
		t.Errorf("Expected status %d, got %d", http.StatusOK, w.Code)
	}

	// Check the response body
	var readings []map[string]interface{}
	err = json.Unmarshal(w.Body.Bytes(), &readings)
	if err != nil {
		t.Fatalf("Failed to parse response: %v", err)
	}

	if len(readings) == 0 {
		t.Errorf("Expected non-empty readings")
	}
}
