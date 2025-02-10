package controller

import (
	"encoding/json"
	"net/http"
	"net/http/httptest"
	"testing"

	"joyofenergy/src/domain"
	"joyofenergy/src/repository"
)

func TestGetPricePlans(t *testing.T) {
	// Create mock repositories
	smartMeterRepo := repository.NewSmartMeterRepository()
	pricePlanRepo := repository.NewPricePlanRepository()

	// Prepare test price plans
	testPlans := []*domain.PricePlan{
		{
			Name:     "green",
			UnitRate: 0.3,
		},
		{
			Name:     "standard",
			UnitRate: 0.2,
		},
	}
	pricePlanRepo.Store(testPlans)

	// Create price plan comparator
	pricePlanComparator := NewPricePlanComparator(smartMeterRepo, pricePlanRepo)

	// Create price plan comparator controller
	accountService := NewAccountService()
	pricePlanComparatorController := NewPricePlanComparatorController(
		pricePlanComparator,
		accountService,
	)

	// Create a request
	req, err := http.NewRequest(http.MethodGet, "/price-plans", nil)
	if err != nil {
		t.Fatalf("Failed to create request: %v", err)
	}

	// Create a response recorder
	w := httptest.NewRecorder()

	// Call the handler
	pricePlanComparatorController.GetPricePlans(w, req)

	// Check the status code
	if w.Code != http.StatusOK {
		t.Errorf("Expected status %d, got %d", http.StatusOK, w.Code)
	}

	// Check the response body
	var pricePlans []map[string]interface{}
	err = json.Unmarshal(w.Body.Bytes(), &pricePlans)
	if err != nil {
		t.Fatalf("Failed to parse response: %v", err)
	}

	if len(pricePlans) != len(testPlans) {
		t.Errorf("Expected %d price plans, got %d", len(testPlans), len(pricePlans))
	}
}

func TestComparePricePlans(t *testing.T) {
	// Create mock repositories
	smartMeterRepo := repository.NewSmartMeterRepository()
	pricePlanRepo := repository.NewPricePlanRepository()

	// Prepare test price plans
	testPlans := []*domain.PricePlan{
		{
			Name:     "green",
			UnitRate: 0.3,
		},
		{
			Name:     "standard",
			UnitRate: 0.2,
		},
	}
	pricePlanRepo.Store(testPlans)

	// Prepare test smart meter with readings
	testMeterID := "smart-meter-test"
	testReadings := []*domain.ElectricityReading{
		{
			Time:    1614335400,
			Reading: 35.5,
		},
	}
	smartMeter := domain.NewSmartMeter(nil, testReadings)
	smartMeterRepo.Save(testMeterID, smartMeter)

	// Create price plan comparator
	pricePlanComparator := NewPricePlanComparator(smartMeterRepo, pricePlanRepo)

	// Create price plan comparator controller
	accountService := NewAccountService()
	pricePlanComparatorController := NewPricePlanComparatorController(
		pricePlanComparator,
		accountService,
	)

	// Create a request
	req, err := http.NewRequest(http.MethodGet, "/price-plans/compare/"+testMeterID, nil)
	if err != nil {
		t.Fatalf("Failed to create request: %v", err)
	}

	// Create a response recorder
	w := httptest.NewRecorder()

	// Call the handler
	pricePlanComparatorController.ComparePricePlans(w, req)

	// Check the status code
	if w.Code != http.StatusOK {
		t.Errorf("Expected status %d, got %d", http.StatusOK, w.Code)
	}

	// Check the response body
	var comparisons []map[string]float64
	err = json.Unmarshal(w.Body.Bytes(), &comparisons)
	if err != nil {
		t.Fatalf("Failed to parse response: %v", err)
	}

	if len(comparisons) == 0 {
		t.Errorf("Expected non-empty price plan comparisons")
	}

	// Verify the order of comparisons (cheapest first)
	var prevCost float64
	for i, comparison := range comparisons {
		var cost float64
		for _, v := range comparison {
			cost = v
			break
		}
		if i > 0 && cost < prevCost {
			t.Errorf("Comparisons not sorted from cheapest to most expensive")
		}
		prevCost = cost
	}
}
