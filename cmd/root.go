package cmd

import (
	"fmt"

	"github.com/silaselisha/coins/cmd/subcmd"
	"github.com/spf13/cobra"
  "github.com/spf13/viper"
)

var (
  rootCmd = &cobra.Command{
    Use: "coins",
    Short: "crypto coins fetcher",
    Long: "query to receive current live data",
    Run: func(cmd *cobra.Command, args []string) {
      fmt.Println(`
      
 ___  ___  _   _  ___  _____  __     ___  __   _  _   _  ___  
(  _)(   )( \_/ )(   )(_   _)(  )   (  _)(  ) ( )( \ ( )/  _) 
| |  | O  |\   / | O  | | |  /  \   | |  /  \ | || \\| |\_"-. 
( )_ ( _ (  ( )  ( __/  ( ) ( O  )  ( )_( O  )( )( )\\ ) __) )
/___\/_\\_| |_|  /_\    /_\  \__/   /___\\__/ /_\/_\ \_\/___/ 
                                                                                                                                                   `)
    },
  }
)

func init() {
  subcmd.FetchCmd.Flags().StringVarP(&subcmd.Listing, "listing", "l", "latest", "Returns the latest market quote for 1 or more cryptocurrencies.")
  rootCmd.AddCommand(subcmd.FetchCmd)
  rootCmd.AddCommand(subcmd.VersionCmd)
  viper.BindPFlags(subcmd.FetchCmd.Flags())
}

func Execute() error {
  return rootCmd.Execute()
}
