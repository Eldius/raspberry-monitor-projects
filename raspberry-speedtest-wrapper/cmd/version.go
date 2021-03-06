package cmd

import (
	"fmt"

	"github.com/spf13/cobra"
)

// versionCmd represents the version command
var versionCmd = &cobra.Command{
	Use:   "version",
	Short: "Shows version and/or build time information",
	Long:  `Shows version and/or build time information.`,
	Run: func(cmd *cobra.Command, args []string) {
		msg := "---\nnetwork-monitor\n"
		msg += fmt.Sprintf("- version: %s\n", gitCommit)
		msg += fmt.Sprintf("- build date: %s\n---", buildDate)
		fmt.Println(msg)
	},
}

var (
	buildDate string
	gitCommit string
)

func init() {
	rootCmd.AddCommand(versionCmd)

	// Here you will define your flags and configuration settings.

	// Cobra supports Persistent Flags which will work for this command
	// and all subcommands, e.g.:
	// versionCmd.PersistentFlags().String("foo", "", "A help for foo")

	// Cobra supports local flags which will only run when this command
	// is called directly, e.g.:
	// versionCmd.Flags().BoolP("toggle", "t", false, "Help message for toggle")
}
