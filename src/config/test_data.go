package config

type MeterReading struct {
	Time  string  `json:"time"`
	Value float64 `json:"value"`
}

var TestMeterReadings = map[string][]MeterReading{
	"smart-meter-0": {
		{Time: "2023-01-01T00:00:00Z", Value: 10.0},
		{Time: "2023-01-02T00:00:00Z", Value: 20.0},
	},
	"smart-meter-1": {
		{Time: "2023-01-01T00:00:00Z", Value: 15.0},
		{Time: "2023-01-02T00:00:00Z", Value: 25.0},
	},
}
