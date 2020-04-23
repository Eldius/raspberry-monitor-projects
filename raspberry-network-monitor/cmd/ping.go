package cmd

import (
	"fmt"

	"github.com/Eldius/raspberry-monitor-projects/raspberry-network-monitor/config"
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
		pingResponse := network.MultiplePingParallel(cfg.PingHosts, cfg.QtdPackets)
		fmt.Printf("---\nping called\nresponses:\n%v\n---\n", pingResponse)
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
