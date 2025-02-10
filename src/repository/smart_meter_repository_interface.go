package repository

import "joyofenergy/src/domain"

type SmartMeterRepositoryInterface interface {
	FindByID(meterID string) *domain.SmartMeter
	Save(smartMeterID string, meter *domain.SmartMeter)
	GetAllMeters() []*domain.SmartMeter
}
