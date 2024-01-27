package cmd

import (
	"fmt"
	"log"
	"os"

	"github.com/silaselisha/coins/cmd/subcmd"
	"github.com/spf13/cobra"
	"github.com/spf13/viper"
)

var (
	rootCmd = &cobra.Command{
		Use:     "coins",
		Short:   "A CLI for crypto coins",
		Long:    "CLI for querying crypto api to receive current live data",
		Version: "v1.0.0",
		Run: func(cmd *cobra.Command, args []string) {
			version, err := cmd.Flags().GetBool("version")
			if err != nil {
				log.Print(err)
				os.Exit(1)
			}
			if version {
				fmt.Println("coins version: v1.0.0")
			}
		},
	}
)

func init() {
	subcmd.FetchCmd.Flags().StringVarP(&subcmd.Listings, "listings", "l", "latest", "Returns the latest market quote for 1 or more cryptocurrencies.")
  subcmd.FetchCmd.Flags().StringVarP(&subcmd.Watch, "watch", "w", "10","number of crypto coins to watch")
	rootCmd.Flags().BoolP("version", "v", false, "coins current version")
	rootCmd.AddCommand(subcmd.FetchCmd)
	rootCmd.AddCommand(subcmd.VersionCmd)
	viper.BindPFlags(rootCmd.Flags())

}

func Execute() error {
	return rootCmd.Execute()
}
