package cmd

import (
	"fmt"
	"github.com/spf13/cobra"
)

// traceCmd represents the trace command
var versionCmd = &cobra.Command{
	Use:   "version",
	Short: "app version",
	Run: func(cmd *cobra.Command, args []string) {
		if len(args) > 0 {
			for _, ip := range args {
				showData(ip)
			}
		} else {
			fmt.Println("v1.0.0")
		}
	},
}

func init() {
	rootCmd.AddCommand(versionCmd)
}
