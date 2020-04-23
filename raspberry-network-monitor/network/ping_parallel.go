package network

import (
	"sync"

	"github.com/sparrc/go-ping"
)

/*
MultiplePingParallel pings some hosts
*/
func MultiplePingParallel(hosts []string, packets int) []PingResponse {
	ch := make(chan PingResponse, len(hosts))
	var wg sync.WaitGroup

	for _, h := range hosts {
		wg.Add(1)
		go SinglePingParallel(h, packets, &wg, ch)
	}
	wg.Wait()
	close(ch)

	s := make([]PingResponse, 0)
	for {
		value, ok := <-ch
		if ok == false {
			break
		}
		s = append(s, value)
	}
	return s
}

/*
SinglePingParallel executes a single ping call
*/
func SinglePingParallel(host string, packets int, wg *sync.WaitGroup, ch chan PingResponse) {
	defer wg.Done()

	pinger, err := ping.NewPinger(host)
	if err != nil {
		panic(err.Error())
	}

	pinger.Count = packets
	pinger.SetPrivileged(true)
	pinger.Run() // blocks until finished

	response := parseToPingResponse(pinger.Statistics(), host)
	ch <- response
}

func parseToPingResponse(stats *ping.Statistics, host string) PingResponse {
	return PingResponse{
		AvgTime:         stats.AvgRtt,
		MinTime:         stats.MinRtt,
		MaxTime:         stats.MaxRtt,
		Jitter:          stats.MaxRtt.Milliseconds() - stats.MinRtt.Milliseconds(),
		PacketsSent:     stats.PacketsSent,
		PacketsReceived: stats.PacketsRecv,
		ResponseTimes:   stats.Rtts,
		Host:            host,
	}
}
