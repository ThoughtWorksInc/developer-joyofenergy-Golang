package domain

type ElectricityReading struct {
	Time    int64   `json:"time"`
	Reading float64 `json:"reading"`
}

func NewElectricityReading(jsonData map[string]interface{}) *ElectricityReading {
	return &ElectricityReading{
		Time:    int64(jsonData["time"].(float64)),
		Reading: jsonData["reading"].(float64),
	}
}

func (e *ElectricityReading) ToJSON() map[string]interface{} {
	return map[string]interface{}{
		"time":    e.Time,
		"reading": e.Reading,
	}
}
