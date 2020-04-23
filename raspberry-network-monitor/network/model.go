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
