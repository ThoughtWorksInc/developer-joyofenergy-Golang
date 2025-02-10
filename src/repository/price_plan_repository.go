package repository

import (
	"math"
	"sort"
	"sync"

	"joyofenergy/src/domain"
	"joyofenergy/src/util"
)

var _ PricePlanRepositoryInterface = &PricePlanRepository{}

type PricePlanRepository struct {
	pricePlans [](*domain.PricePlan)
	mu         sync.RWMutex
}

func NewPricePlanRepository() *PricePlanRepository {
	return &PricePlanRepository{
		pricePlans: [](*domain.PricePlan){},
	}
}

func (r *PricePlanRepository) Store(newPricePlans []*domain.PricePlan) {
	r.mu.Lock()
	defer r.mu.Unlock()
	r.pricePlans = append(r.pricePlans, newPricePlans...)
}

func (r *PricePlanRepository) GetPricePlans() []map[string]interface{} {
	r.mu.RLock()
	defer r.mu.RUnlock()

	plans := make([]map[string]interface{}, len(r.pricePlans))
	for i, plan := range r.pricePlans {
		plans[i] = map[string]interface{}{
			"planName": plan.Name,
			"unitRate": plan.UnitRate,
		}
	}
	return plans
}

func (r *PricePlanRepository) Clear() {
	r.mu.Lock()
	defer r.mu.Unlock()
	r.pricePlans = [](*domain.PricePlan){}
}

func (r *PricePlanRepository) GetListOfSpendAgainstEachPricePlanFor(smartMeter *domain.SmartMeter, limit *int) []map[string]float64 {
	readings := smartMeter.ElectricityReadings
	if len(readings) < 1 {
		return []map[string]float64{}
	}

	average := r.calculateAverageReading(readings)
	timeElapsed := calculateTimeElapsed(readings)
	consumedEnergy := average / timeElapsed

	pricePlans := r.pricePlans
	cheapestPlans := r.cheapestPlansFirst(pricePlans)

	listOfSpend := make([]map[string]float64, 0, len(cheapestPlans))
	for _, plan := range cheapestPlans {
		cost := map[string]float64{
			plan.Name: consumedEnergy * plan.UnitRate,
		}
		listOfSpend = append(listOfSpend, cost)
	}

	if limit != nil && *limit < len(listOfSpend) {
		return listOfSpend[:*limit]
	}
	return listOfSpend
}

func (r *PricePlanRepository) cheapestPlansFirst(pricePlans []*domain.PricePlan) []*domain.PricePlan {
	sortedPlans := append([]*domain.PricePlan{}, pricePlans...)
	sort.Slice(sortedPlans, func(i, j int) bool {
		return sortedPlans[i].UnitRate < sortedPlans[j].UnitRate
	})
	return sortedPlans
}

func (r *PricePlanRepository) calculateAverageReading(readings []*domain.ElectricityReading) float64 {
	total := 0.0
	for _, reading := range readings {
		total += reading.Reading
	}
	return total / float64(len(readings))
}

func calculateTimeElapsed(readings []*domain.ElectricityReading) float64 {
	minTime := math.Inf(1)
	maxTime := math.Inf(-1)

	for _, reading := range readings {
		minTime = math.Min(minTime, float64(reading.Time))
		maxTime = math.Max(maxTime, float64(reading.Time))
	}

	return util.TimeElapsedInHours(int64(minTime), int64(maxTime))
}
