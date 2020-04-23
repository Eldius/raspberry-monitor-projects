package network

import (
	"github.com/sparrc/go-ping"
)

/*
MultiplePing pings some hosts
*/
func MultiplePing(hosts []string, packets int) []PingResponse {
	var responseList []PingResponse

	for _, h := range hosts {
		responseList = append(responseList, SinglePing(h, packets))
	}
	return responseList
}

/*
SinglePing executes a single ping call
*/
func SinglePing(host string, packets int) PingResponse {
	pinger, err := ping.NewPinger(host)
	if err != nil {
		panic(err.Error())
	}
	pinger.Count = packets
	pinger.SetPrivileged(true)
	pinger.Run() // blocks until finished

	return parseToPingResponse(pinger.Statistics(), host)
}
