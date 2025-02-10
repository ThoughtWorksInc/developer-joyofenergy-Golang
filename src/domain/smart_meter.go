package domain

type SmartMeter struct {
	PricePlan           *PricePlan
	ElectricityReadings []*ElectricityReading
}

func (sm *SmartMeter) AddReadings(readings []*ElectricityReading) {
	sm.ElectricityReadings = append(readings, sm.ElectricityReadings...)
}

func NewSmartMeter(pricePlan *PricePlan, readings []*ElectricityReading) *SmartMeter {
	return &SmartMeter{
		PricePlan:           pricePlan,
		ElectricityReadings: readings,
	}
}
