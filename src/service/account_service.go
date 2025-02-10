package service

import "fmt"

type AccountService struct {
	planIDsByMeter map[string]string
}

func NewAccountService() *AccountService {
	return &AccountService{
		planIDsByMeter: map[string]string{
			"smart-meter-0": "price-plan-0",
			"smart-meter-1": "price-plan-1",
			"smart-meter-2": "price-plan-0",
			"smart-meter-3": "price-plan-2",
			"smart-meter-4": "price-plan-1",
		},
	}
}

func (as *AccountService) GetPricePlan(smartMeterID string) (string, error) {
	pricePlan, exists := as.planIDsByMeter[smartMeterID]
	if !exists {
		return "", fmt.Errorf("no price plan found for smart meter ID: %s", smartMeterID)
	}
	return pricePlan, nil
}
