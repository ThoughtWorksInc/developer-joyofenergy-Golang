package readings

import (
	"fmt"
	"time"

	"joyofenergy/src/meters"
)

type ReadingService struct {
	meterService *meters.MeterService
}

func NewReadingService(meterService *meters.MeterService) *ReadingService {
	return &ReadingService{
		meterService: meterService,
	}
}

func (rs *ReadingService) StoreReading(meterID string, reading float64) error {
	meter, err := rs.meterService.GetMeterByID(meterID)
	if err != nil {
		return err
	}

	newReading := meters.MeterReading{
		Time:  time.Now().UTC().Format(time.RFC3339),
		Value: reading,
	}

	meter.AddReading(newReading)
	return nil
}

func (rs *ReadingService) GetReadings(meterID string) ([]meters.MeterReading, error) {
	meter, err := rs.meterService.GetMeterByID(meterID)
	if err != nil {
		return nil, fmt.Errorf("failed to get readings: %v", err)
	}

	return meter.Readings, nil
}
