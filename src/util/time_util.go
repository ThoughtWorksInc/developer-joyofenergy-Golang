package util

import "time"

func TimeElapsedInHours(startTime, endTime int64) float64 {
	start := time.Unix(startTime, 0)
	end := time.Unix(endTime, 0)
	duration := end.Sub(start)
	return duration.Hours()
}
