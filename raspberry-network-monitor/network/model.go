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
MinTimeInMili returns MinTime as milisseconds
*/
func (p *PingResponse) MinTimeInMili() int64 {
	return p.MinTime.Milliseconds()
}

/*
MaxTimeInMili returns MaxTime as milisseconds
*/
func (p *PingResponse) MaxTimeInMili() int64 {
	return p.MaxTime.Milliseconds()
}

/*
AvgTimeInMili returns AvgTime as milisseconds
*/
func (p *PingResponse) AvgTimeInMili() int64 {
	return p.AvgTime.Milliseconds()
}
