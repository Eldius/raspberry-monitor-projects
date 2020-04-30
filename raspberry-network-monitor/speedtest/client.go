package speedtest

import (
	"fmt"

	"github.com/Eldius/raspberry-monitor-projects/raspberry-network-monitor/logger"
	"github.com/guifibages/speedtest/speedtest"
)

/*
Test to run the tests
*/
func Test() {
	c := new(speedtest.OoklaClient)

	if servers, err := FindServers(); err != nil {
		panic(err.Error())
	} else {
		//FindFasterServers(servers)
		for _, s := range servers.Servers.Server {
			logger.Debug(s)
		}
	}
	err := c.GetConfig()
	if err != nil {
		panic(err.Error())
	}

	fmt.Printf("Speedtest against %s with a %d seconds timeout\n", c.Server, c.Timeout)
	//fmt.Printf("IP: %s\nLon:%s\nLat:%s\n", c.Client.IP, c.Client.Lon, c.Client.Lat)

	//c.TestServer()
}
