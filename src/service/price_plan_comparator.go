package service

import (
	"fmt"
	"sort"

	"joyofenergy/src/repository"
)

type PricePlanComparator struct {
	SmartMeterRepository repository.SmartMeterRepositoryInterface
	PricePlanRepository  repository.PricePlanRepositoryInterface
}

func NewPricePlanComparator(
	smartMeterRepo repository.SmartMeterRepositoryInterface,
	pricePlanRepo repository.PricePlanRepositoryInterface,
) *PricePlanComparator {
	return &PricePlanComparator{
		SmartMeterRepository: smartMeterRepo,
		PricePlanRepository:  pricePlanRepo,
	}
}

func (pc *PricePlanComparator) RecommendCheapestPricePlans(smartMeterID string, limit *int) ([]map[string]float64, error) {
	smartMeter := pc.SmartMeterRepository.FindByID(smartMeterID)
	if smartMeter == nil {
		return nil, fmt.Errorf("missing smart meter with ID: %s", smartMeterID)
	}

	consumptionForPricePlans := pc.PricePlanRepository.GetListOfSpendAgainstEachPricePlanFor(smartMeter, limit)

	// Sort the consumption list
	sort.Slice(consumptionForPricePlans, func(i, j int) bool {
		// Compare the first (and only) value in each map
		var valueI, valueJ float64
		for _, v := range consumptionForPricePlans[i] {
			valueI = v
		}
		for _, v := range consumptionForPricePlans[j] {
			valueJ = v
		}
		return valueI < valueJ
	})

	return consumptionForPricePlans, nil
}
