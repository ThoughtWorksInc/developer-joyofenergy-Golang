package meters

type Meter struct {
	ID       string
	Readings []MeterReading
	Type     string
}

type MeterReading struct {
	Time  string
	Value float64
}

func NewMeter(id, meterType string) *Meter {
	return &Meter{
		ID:   id,
		Type: meterType,
	}
}

func (m *Meter) AddReading(reading MeterReading) {
	m.Readings = append(m.Readings, reading)
}
