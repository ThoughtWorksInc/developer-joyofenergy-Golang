package usage

import (
	"joyofenergy/src/readings"
)

type UsageService struct {
	readingService *readings.ReadingService
}

func NewUsageService(readingService *readings.ReadingService) *UsageService {
	return &UsageService{
		readingService: readingService,
	}
}

func (us *UsageService) GetTotalUsage(meterID string) (float64, error) {
	readings, err := us.readingService.GetReadings(meterID)
	if err != nil {
		return 0, err
	}

	totalUsage := 0.0
	for _, reading := range readings {
		totalUsage += reading.Value
	}

	return totalUsage, nil
}
