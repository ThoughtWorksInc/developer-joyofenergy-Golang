package tests

import (
	"testing"

	"joyofenergy/src/config"
	"joyofenergy/src/meters"
	priceplans "joyofenergy/src/pricePlans"
)

func TestPricePlanComparator(t *testing.T) {
	pricePlans := []config.PricePlan{
		{Name: "standard", UnitRate: 0.20, StandingCharge: 5.00},
		{Name: "night", UnitRate: 0.10, StandingCharge: 3.50},
	}

	meter := meters.NewMeter("test-meter", "electric")
	meter.AddReading(meters.MeterReading{Time: "2023-01-01T00:00:00Z", Value: 10.0})

	comparator := priceplans.NewPricePlanComparator(pricePlans)

	costs := comparator.GetCostForMeter(meter)
	if len(costs) != 2 {
		t.Errorf("Expected 2 price plan costs, got %d", len(costs))
	}

	recommendedPlans := comparator.GetRecommendedPricePlans(meter, 1)
	if len(recommendedPlans) != 1 {
		t.Errorf("Expected 1 recommended plan, got %d", len(recommendedPlans))
	}
}

func TestPricePlanCostCalculation(t *testing.T) {
	pricePlans := []config.PricePlan{
		{Name: "standard", UnitRate: 0.20, StandingCharge: 5.00},
		{Name: "night", UnitRate: 0.10, StandingCharge: 3.50},
	}

	meter := meters.NewMeter("test-meter", "electric")
	meter.AddReading(meters.MeterReading{Time: "2023-01-01T00:00:00Z", Value: 10.0})

	comparator := priceplans.NewPricePlanComparator(pricePlans)

	costs := comparator.GetCostForMeter(meter)

	expectedStandardCost := (10.0 * 0.20) + 5.00
	expectedNightCost := (10.0 * 0.10) + 3.50

	if costs["standard"] != expectedStandardCost {
		t.Errorf("Expected standard plan cost of %.2f, got %.2f", expectedStandardCost, costs["standard"])
	}

	if costs["night"] != expectedNightCost {
		t.Errorf("Expected night plan cost of %.2f, got %.2f", expectedNightCost, costs["night"])
	}
}
