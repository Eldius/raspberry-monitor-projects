package network

import (
	"time"
)

/*
PingResponse represents the ping execution response
*/
type PingResponse struct {
	AvgTime time.Duration
	MinTime time.Duration
	MaxTime time.Duration
	Jitter  int64

	PacketsSent     int
	PacketsReceived int

	ResponseTimes []time.Duration
	Host          string
	ExecutionTime time.Time
}

/*
PingResponseTO is a converted type to publish
*/
type PingResponseTO struct {
	AvgTime float64
	MinTime float64
	MaxTime float64
	Jitter  int64

	PacketsSent     int
	PacketsReceived int

	ResponseTimes []float64
	Host          string
	ExecutionTime time.Time
}

/*
Convert converts the respon se to a TO value
*/
func (r *PingResponse) Convert() PingResponseTO {
	return PingResponseTO{
		AvgTime:         ConvertToMili(r.AvgTime),
		MinTime:         ConvertToMili(r.MinTime),
		MaxTime:         ConvertToMili(r.MaxTime),
		Jitter:          r.Jitter,
		PacketsSent:     r.PacketsSent,
		PacketsReceived: r.PacketsReceived,
		ResponseTimes:   convertAllToMili(r.ResponseTimes),
		Host:            r.Host,
		ExecutionTime:   r.ExecutionTime,
	}
}

/*
ConvertToMili converts time.Duration to milisseconds int64
*/
func ConvertToMili(d time.Duration) float64 {
	return float64(d / time.Millisecond)
}

func convertAllToMili(ds []time.Duration) []float64 {
	var parsedValues []float64
	for _, d := range ds {
		parsedValues = append(parsedValues, ConvertToMili(d))
	}
	return parsedValues
}
