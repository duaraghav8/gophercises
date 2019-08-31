package main

import (
	"fmt"
	"github.com/spf13/cobra"
)

func init() {
	app.AddCommand(versionCmd)
}

const version = "1.0.0"

var (
	app = &cobra.Command{
		Use:   "secretstore",
		Short: "Manage secrets and protect sensitive data",
		Long: `Secretstore is a secrets management engine that allows you
to quickly and securely store and retrieve small pieces of
data.`,
		Run: func(cmd *cobra.Command, args []string) {},
	}

	versionCmd = &cobra.Command{
		Use:   "version",
		Short: "Display app version",
		Run: func(cmd *cobra.Command, args []string) {
			fmt.Printf("Secretstore v%s\n", version)
		},
	}
)

var encodingKey string