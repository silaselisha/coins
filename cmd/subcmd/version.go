package subcmd

import (
	"fmt"

	"github.com/spf13/cobra"
)

var (
	VersionCmd = &cobra.Command{
		Use: "version",
    Short: "Application version",
    Long: "Log the program version for program backwards compatibility",
    Run: func(cmd *cobra.Command, args []string) {
      fmt.Println("coins version v1.0.0")
    },
	}
)

