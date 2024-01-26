package cmd

import (
	"fmt"

	"github.com/silaselisha/coins/cmd/subcmd"
	"github.com/spf13/cobra"
)

var (
  rootCmd = &cobra.Command{
    Use: "coins",
    Short: "crypto coins fetcher",
    Long: "query to receive current live data",
    Run: func(cmd *cobra.Command, args []string) {
      fmt.Println("fetching...")
    },
  }
)

func init() {
  rootCmd.AddCommand(subcmd.FetchCmd)
  rootCmd.AddCommand(subcmd.VersionCmd)
}

func Execute() error {
  return rootCmd.Execute()
}
