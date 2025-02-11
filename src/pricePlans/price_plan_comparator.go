package priceplans

import (
	"math"
	"sort"

	"joyofenergy/src/config"
	"joyofenergy/src/meters"
)

type PricePlanComparator struct {
	pricePlans []config.PricePlan
}

func NewPricePlanComparator(pricePlans []config.PricePlan) *PricePlanComparator {
	return &PricePlanComparator{
		pricePlans: pricePlans,
	}
}

func (ppc *PricePlanComparator) GetCostForMeter(meter *meters.Meter) map[string]float64 {
	costs := make(map[string]float64)

	for _, plan := range ppc.pricePlans {
		totalCost := ppc.calculateTotalCost(meter, plan)
		costs[plan.Name] = math.Round(totalCost*100) / 100
	}

	return costs
}

func (ppc *PricePlanComparator) calculateTotalCost(meter *meters.Meter, plan config.PricePlan) float64 {
	totalUsage := 0.0
	for _, reading := range meter.Readings {
		totalUsage += reading.Value
	}

	return (totalUsage * plan.UnitRate) + plan.StandingCharge
}

func (ppc *PricePlanComparator) GetRecommendedPricePlans(meter *meters.Meter, numberOfPlans int) []string {
	costs := ppc.GetCostForMeter(meter)

	type planCost struct {
		name string
		cost float64
	}

	var planCosts []planCost
	for name, cost := range costs {
		planCosts = append(planCosts, planCost{name, cost})
	}

	sort.Slice(planCosts, func(i, j int) bool {
		return planCosts[i].cost < planCosts[j].cost
	})

	var recommendedPlans []string
	for i := 0; i < numberOfPlans && i < len(planCosts); i++ {
		recommendedPlans = append(recommendedPlans, planCosts[i].name)
	}

	return recommendedPlans
}
