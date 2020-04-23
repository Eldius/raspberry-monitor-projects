package network

import (
	"fmt"
	"sync"

	"github.com/Eldius/raspberry-monitor-projects/raspberry-network-monitor/logger"
	"github.com/sparrc/go-ping"
)

/*
MultiplePingParallel pings some hosts
*/
func MultiplePingParallel(hosts []string, packets int) []PingResponse {
	ch := make(chan PingResponse, len(hosts))
	var wg sync.WaitGroup

	for _, h := range hosts {
		logger.Debug(fmt.Sprintf(" -> Starting a new ping request to %s", h))
		wg.Add(1)
		go SinglePingParallel(h, packets, &wg, ch)
	}
	logger.Debug("waiting for ping responses...")
	wg.Wait()
	logger.Debug("closing channel...")
	close(ch)

	s := make([]PingResponse, 0)
	for {
		value, ok := <-ch
		if ok == false {
			break
		}
		logger.Debug(fmt.Sprintf("appending %v response to slice...", value))
		s = append(s, value)
	}
	logger.Debug("returning ping responses...")
	return s
}

/*
SinglePingParallel executes a single ping call
*/
func SinglePingParallel(host string, packets int, wg *sync.WaitGroup, ch chan PingResponse) {
	defer wg.Done()
	logger.Println("pinging", host)

	pinger, err := ping.NewPinger(host)
	if err != nil {
		panic(err.Error())
	}

	pinger.Count = packets
	pinger.SetPrivileged(true)
	pinger.Run() // blocks until finished
	logger.Println("Finished pinging:", host)

	response := parseToPingResponse(pinger.Statistics(), host)
	logger.Println("Finished pinging 2:", host)
	ch <- response
	logger.Println("Finished pinging 3:", host)
	//wg.Done()
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
