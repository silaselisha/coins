package cmd

import (
 "github.com/spf13/cobra"
 "fmt"
)

var rootCmd = &cobra.Command{
  Use: "hugo",
  Short: "hugo is a very fast static site generator",
  Long: `A Fast and Flexible Static Site Generator built with
              love by spf13 and friends in Go.
              Complete documentation is available at https://gohugo.io/documentation/`,
  Run: func(cmd *cobra.Command, args []string) {
    // Do stuff here
    fmt.Println("World!")
  },
}

func Execute() error {
  return rootCmd.Execute()
}
