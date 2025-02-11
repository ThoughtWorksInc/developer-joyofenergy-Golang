package config

type PricePlan struct {
	Name           string  `json:"name"`
	UnitRate       float64 `json:"unitRate"`
	StandingCharge float64 `json:"standingCharge"`
}

var PricePlans = []PricePlan{
	{
		Name:           "standard",
		UnitRate:       0.20,
		StandingCharge: 5.00,
	},
	{
		Name:           "night",
		UnitRate:       0.10,
		StandingCharge: 3.50,
	},
	{
		Name:           "weekends",
		UnitRate:       0.15,
		StandingCharge: 4.00,
	},
}
