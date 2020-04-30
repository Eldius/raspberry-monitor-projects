package speedtest

import (
	"encoding/xml"
	"fmt"
	"io/ioutil"
	"net/http"
	"net/url"
	"sort"

	"github.com/Eldius/raspberry-monitor-projects/raspberry-network-monitor/network"
)

const (
	// ServerListURL URL to fetch server list
	ServerListURL = "https://c.speedtest.net/speedtest-servers-static.php"
)

/*
ServersSettings is a server list configuration
*/
type ServersSettings struct {
	XMLName xml.Name `xml:"settings"`
	Text    string   `xml:",chardata"`
	Servers struct {
		Text   string `xml:",chardata"`
		Server []struct {
			Text    string  `xml:",chardata"`
			URL     string  `xml:"url,attr"`
			Lat     float64 `xml:"lat,attr"`
			Lon     float64 `xml:"lon,attr"`
			Name    string  `xml:"name,attr"`
			Country string  `xml:"country,attr"`
			Cc      string  `xml:"cc,attr"`
			Sponsor string  `xml:"sponsor,attr"`
			ID      string  `xml:"id,attr"`
			Host    string  `xml:"host,attr"`
		} `xml:"server"`
	} `xml:"servers"`
}

/*
FindServers finds some servers
*/
func FindServers() (servers ServersSettings, err error) {
	//servers = make([]ServerSpec, 0)
	res, err := http.Get(ServerListURL)
	if err != nil {
		return servers, err
	}
	configxml, err := ioutil.ReadAll(res.Body)
	res.Body.Close()
	if err != nil {
		return servers, err
	}

	err = xml.Unmarshal(configxml, &servers)
	if err != nil {
		fmt.Printf("error: %v", err)
		return servers, err
	}
	return servers, nil

}

/*
FindFasterServers execute ping test
*/
func FindFasterServers(servers ServersSettings) {
	var serverList []string
	for i, s := range servers.Servers.Server {
		if u, err := url.Parse(fmt.Sprintf("http://%s", s.Host)); err == nil {
			serverList = append(serverList, u.Hostname())
		}
		if i > 10 {
			break
		}
	}

	pingResults := network.MultiplePingParallel(serverList, 5)

	//slice.Sort(pingResults[:], func(i, j int) bool {
	//	return pingResults[i].AvgTime < pingResults[j].AvgTime
	//})

	sort.Slice(pingResults, func(i, j int) bool {
		return pingResults[i].AvgTime < pingResults[j].AvgTime
	})
	fmt.Println("first: %v\nlast: %v", pingResults[0], pingResults[len(pingResults)-1])

}
