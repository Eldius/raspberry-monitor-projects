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
	AvgTime int64
	MinTime int64
	MaxTime int64
	Jitter  int64

	PacketsSent     int
	PacketsReceived int

	ResponseTimes []int64
	Host          string
	ExecutionTime time.Time
}

/*
Convert converts the respon se to a TO value
*/
func (r *PingResponse) Convert() PingResponseTO {
	return PingResponseTO{
		AvgTime:         convertToMili(r.AvgTime),
		MinTime:         convertToMili(r.MinTime),
		MaxTime:         convertToMili(r.MaxTime),
		Jitter:          r.Jitter,
		PacketsSent:     r.PacketsSent,
		PacketsReceived: r.PacketsReceived,
		ResponseTimes:   convertAllToMili(r.ResponseTimes),
		Host:            r.Host,
		ExecutionTime:   r.ExecutionTime,
	}
}


func convertToMili(d time.Duration) int64 {
	return int64(d / time.Millisecond)
}

func convertAllToMili(ds []time.Duration) []int64 {
	var parsedValues []int64
	for _, d := range ds {
		parsedValues = append(parsedValues, int64(d / time.Millisecond))
	}
	return parsedValues
}
