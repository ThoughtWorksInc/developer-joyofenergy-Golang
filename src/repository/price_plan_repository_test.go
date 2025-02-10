package repository

import (
	"testing"

	"joyofenergy/src/domain"
)

func TestPricePlanRepository(t *testing.T) {
	// Create a new repository
	repo := NewPricePlanRepository()

	// Test case 1: Store and Get Price Plans
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

	// Store the price plans
	repo.Store(testPlans)

	// Get the price plans
	storedPlans := repo.GetPricePlans()

	// Verify the number of stored plans
	if len(storedPlans) != len(testPlans) {
		t.Errorf("Expected %d price plans, got %d", len(testPlans), len(storedPlans))
	}

	// Verify the content of stored plans
	for i, plan := range storedPlans {
		expectedPlanName := testPlans[i].Name
		expectedUnitRate := testPlans[i].UnitRate

		if planName, ok := plan["planName"].(string); !ok || planName != expectedPlanName {
			t.Errorf("Expected plan name %s, got %v", expectedPlanName, plan["planName"])
		}

		if unitRate, ok := plan["unitRate"].(float64); !ok || unitRate != expectedUnitRate {
			t.Errorf("Expected unit rate %f, got %v", expectedUnitRate, plan["unitRate"])
		}
	}

	// Test case 2: Clear repository
	repo.Clear()
	clearedPlans := repo.GetPricePlans()
	if len(clearedPlans) != 0 {
		t.Errorf("Expected 0 price plans after clearing, got %d", len(clearedPlans))
	}

	// Test case 3: Store additional plans
	additionalPlans := []*domain.PricePlan{
		{
			Name:     "white",
			UnitRate: 0.4,
		},
	}
	repo.Store(additionalPlans)

	finalPlans := repo.GetPricePlans()
	if len(finalPlans) != len(additionalPlans) {
		t.Errorf("Expected %d price plans, got %d", len(additionalPlans), len(finalPlans))
	}
}

func TestPricePlanComparison(t *testing.T) {
	// Create a new repository
	repo := NewPricePlanRepository()

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
	repo.Store(testPlans)

	// Prepare a test smart meter with readings
	testReadings := []*domain.ElectricityReading{
		{
			Time:    1614335400,
			Reading: 35.5,
		},
	}
	testMeter := domain.NewSmartMeter(nil, testReadings)

	// Test price plan comparison
	comparisonLimit := 1
	comparisons := repo.GetListOfSpendAgainstEachPricePlanFor(testMeter, &comparisonLimit)

	// Verify the number of comparisons
	if len(comparisons) != comparisonLimit {
		t.Errorf("Expected %d comparisons, got %d", comparisonLimit, len(comparisons))
	}

	// Verify the comparison values
	for _, comparison := range comparisons {
		for planName, cost := range comparison {
			// Verify the cost is calculated correctly
			expectedCost := testReadings[0].Reading * testPlans[1].UnitRate // standard plan
			if cost != expectedCost {
				t.Errorf("Expected cost %f for plan %s, got %f", expectedCost, planName, cost)
			}
		}
	}
}
