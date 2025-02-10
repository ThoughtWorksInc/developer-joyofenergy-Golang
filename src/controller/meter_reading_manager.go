package controller

import (
	"joyofenergy/src/domain"
	"joyofenergy/src/repository"
)

type IllegalArgumentError struct {
	message string
}

func (e *IllegalArgumentError) Error() string {
	return e.message
}

type MeterReadingManager struct {
	SmartMeterRepository repository.SmartMeterRepositoryInterface
}

func NewMeterReadingManager(smartMeterRepo repository.SmartMeterRepositoryInterface) *MeterReadingManager {
	return &MeterReadingManager{
		SmartMeterRepository: smartMeterRepo,
	}
}

func (mrm *MeterReadingManager) StoreReading(json map[string]interface{}) error {
	smartMeterID, ok := json["smartMeterId"].(string)
	if !ok || smartMeterID == "" {
		return &IllegalArgumentError{"Smart meter id must be provided!"}
	}

	electricityReadingsData, ok := json["electricityReadings"].([]interface{})
	if !ok {
		return &IllegalArgumentError{"Electricity readings must be provided!"}
	}

	electricityReadings := make([]*domain.ElectricityReading, 0, len(electricityReadingsData))
	for _, readingData := range electricityReadingsData {
		readingMap, ok := readingData.(map[string]interface{})
		if !ok {
			return &IllegalArgumentError{"Invalid electricity reading format"}
		}
		reading := domain.NewElectricityReading(readingMap)
		electricityReadings = append(electricityReadings, reading)
	}

	smartMeter := mrm.SmartMeterRepository.FindByID(smartMeterID)
	if smartMeter != nil {
		smartMeter.AddReadings(electricityReadings)
	} else {
		newMeter := domain.NewSmartMeter(nil, electricityReadings)
		mrm.SmartMeterRepository.Save(smartMeterID, newMeter)
	}

	return nil
}

func (mrm *MeterReadingManager) ReadReadings(smartMeterID string) []*domain.ElectricityReading {
	smartMeter := mrm.SmartMeterRepository.FindByID(smartMeterID)
	if smartMeter == nil {
		return []*domain.ElectricityReading{}
	}
	return smartMeter.ElectricityReadings
}
