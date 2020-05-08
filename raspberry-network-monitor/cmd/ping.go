package cmd

import (
	"fmt"

	"github.com/Eldius/raspberry-monitor-projects/raspberry-network-monitor/config"
	"github.com/Eldius/raspberry-monitor-projects/raspberry-network-monitor/mqttclient"
	"github.com/Eldius/raspberry-monitor-projects/raspberry-network-monitor/network"
	"github.com/spf13/cobra"
)

// pingCmd represents the ping command
var pingCmd = &cobra.Command{
	Use:   "ping",
	Short: "It just ping hosts",
	Long:  `It just ping hosts.`,
	Run: func(cmd *cobra.Command, args []string) {
		fmt.Println("executing ping...")
		cfg := config.AppConfig()
		pingResponses := network.MultiplePingParallel(cfg.PingHosts, cfg.QtdPackets)
		if *publish {
			mqttclient.SendPingResponses(pingResponses, cfg.MQTT)
		}
		out := "---\nping called\nresponses:\n"
		for _, r := range pingResponses {
			out += fmt.Sprintf("- host: %s\n", r.Host)
			out += fmt.Sprintf("    avg: %d ms\n", r.AvgTime)
			out += fmt.Sprintf("    min: %d ms\n", r.MinTime)
			out += fmt.Sprintf("    max: %d ms\n", r.MaxTime)
			out += fmt.Sprintf("    jitter: %d ms\n", r.Jitter)
		}
		out += "---\n"
		fmt.Println(out)
	},
}

func init() {
	rootCmd.AddCommand(pingCmd)
	// Here you will define your flags and configuration settings.

	// Cobra supports Persistent Flags which will work for this command
	// and all subcommands, e.g.:
	// pingCmd.PersistentFlags().String("foo", "", "A help for foo")

	// Cobra supports local flags which will only run when this command
	// is called directly, e.g.:
	// pingCmd.Flags().BoolP("toggle", "t", false, "Help message for toggle")
}
