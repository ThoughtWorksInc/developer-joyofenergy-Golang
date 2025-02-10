package domain

import "time"

type DayOfWeek int

const (
	Sunday    DayOfWeek = 0
	Monday    DayOfWeek = 1
	Tuesday   DayOfWeek = 2
	Wednesday DayOfWeek = 3
	Thursday  DayOfWeek = 4
	Friday    DayOfWeek = 5
	Saturday  DayOfWeek = 6
)

type PeakTimeMultiplier struct {
	DayOfWeek  DayOfWeek
	Multiplier float64
}

type PricePlan struct {
	Name                string
	Supplier            string
	UnitRate            float64
	PeakTimeMultipliers []PeakTimeMultiplier
}

func (pp *PricePlan) GetPrice(dateTime time.Time) float64 {
	matchingMultipliers := []PeakTimeMultiplier{}
	for _, m := range pp.PeakTimeMultipliers {
		if m.DayOfWeek == DayOfWeek(dateTime.Weekday()) {
			matchingMultipliers = append(matchingMultipliers, m)
		}
	}

	if len(matchingMultipliers) > 0 {
		return pp.UnitRate * matchingMultipliers[0].Multiplier
	}
	return pp.UnitRate
}

func NewPricePlan(name, supplier string, unitRate float64, peakTimeMultipliers []PeakTimeMultiplier) *PricePlan {
	return &PricePlan{
		Name:                name,
		Supplier:            supplier,
		UnitRate:            unitRate,
		PeakTimeMultipliers: peakTimeMultipliers,
	}
}
