package repository

import (
	"sync"

	"joyofenergy/src/domain"
)

type SmartMeterRepository struct {
	smartMeters map[string]*domain.SmartMeter
	mu          sync.RWMutex
}

var _ SmartMeterRepositoryInterface = &SmartMeterRepository{}

func NewSmartMeterRepository() *SmartMeterRepository {
	return &SmartMeterRepository{
		smartMeters: make(map[string]*domain.SmartMeter),
	}
}

func (r *SmartMeterRepository) FindByID(meterID string) *domain.SmartMeter {
	r.mu.RLock()
	defer r.mu.RUnlock()
	return r.smartMeters[meterID]
}

func (r *SmartMeterRepository) Save(smartMeterID string, meter *domain.SmartMeter) {
	r.mu.Lock()
	defer r.mu.Unlock()
	r.smartMeters[smartMeterID] = meter
}

func (r *SmartMeterRepository) GetAllMeters() []*domain.SmartMeter {
	r.mu.RLock()
	defer r.mu.RUnlock()

	meters := make([]*domain.SmartMeter, 0, len(r.smartMeters))
	for _, meter := range r.smartMeters {
		meters = append(meters, meter)
	}
	return meters
}
