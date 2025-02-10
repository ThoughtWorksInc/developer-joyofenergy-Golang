package controller

import (
	"encoding/json"
	"io"
	"net/http"
	"strings"
)

type ElectricityReadingController struct {
	meterReadingManager *MeterReadingManager
}

func NewElectricityReadingController(meterReadingManager *MeterReadingManager) *ElectricityReadingController {
	return &ElectricityReadingController{
		meterReadingManager: meterReadingManager,
	}
}

func (c *ElectricityReadingController) StoreReading(w http.ResponseWriter, r *http.Request) {
	// Read request body
	body, err := io.ReadAll(r.Body)
	if err != nil {
		http.Error(w, "Error reading request body", http.StatusBadRequest)
		return
	}
	defer r.Body.Close()

	// Parse JSON body
	var jsonData map[string]interface{}
	if err := json.Unmarshal(body, &jsonData); err != nil {
		http.Error(w, "Invalid JSON", http.StatusBadRequest)
		return
	}

	// Store reading
	if err := c.meterReadingManager.StoreReading(jsonData); err != nil {
		http.Error(w, err.Error(), http.StatusBadRequest)
		return
	}

	w.WriteHeader(http.StatusCreated)
}

func (c *ElectricityReadingController) GetReadings(w http.ResponseWriter, r *http.Request) {
	// Extract smart meter ID from URL
	parts := strings.Split(r.URL.Path, "/")
	if len(parts) < 3 {
		http.Error(w, "Invalid smart meter ID", http.StatusBadRequest)
		return
	}
	smartMeterID := parts[len(parts)-1]

	// Get readings
	readings := c.meterReadingManager.ReadReadings(smartMeterID)

	// Send response
	w.Header().Set("Content-Type", "application/json")
	json.NewEncoder(w).Encode(readings)
}
