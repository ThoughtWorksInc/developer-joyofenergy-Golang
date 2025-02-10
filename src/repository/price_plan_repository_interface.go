package repository

import "joyofenergy/src/domain"

type PricePlanRepositoryInterface interface {
	GetListOfSpendAgainstEachPricePlanFor(smartMeter *domain.SmartMeter, limit *int) []map[string]float64
	GetPricePlans() []map[string]interface{}
}
