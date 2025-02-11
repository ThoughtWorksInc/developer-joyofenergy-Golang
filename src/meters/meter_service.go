package meters

import (
	"fmt"
	"sort"
)

type MeterService struct {
	meters map[string]*Meter
}

func NewMeterService() *MeterService {
	return &MeterService{
		meters: make(map[string]*Meter),
	}
}

func (ms *MeterService) GetMeterByID(id string) (*Meter, error) {
	meter, exists := ms.meters[id]
	if !exists {
		return nil, fmt.Errorf("meter not found: %s", id)
	}
	return meter, nil
}

func (ms *MeterService) AddMeter(meter *Meter) {
	ms.meters[meter.ID] = meter
}

func (ms *MeterService) GetAllMeterIDs() []string {
	ids := make([]string, 0, len(ms.meters))
	for id := range ms.meters {
		ids = append(ids, id)
	}
	sort.Strings(ids)
	return ids
}
