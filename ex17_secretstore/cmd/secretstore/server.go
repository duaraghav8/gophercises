package main

import "github.com/spf13/cobra"

var serverCmd = &cobra.Command{
	Use:   "server",
	Short: "Start Secretstore server",
	Long: `Start the server.
This server serves the REST API to allow users to interact
with Secretsmanager to store and fetch data.
The server manages the storage backend and all cryptographic
tasks.`,
	Run: func(cmd *cobra.Command, args []string) {},
}

var (
	serverBindAddr = "127.0.0.1"
	serverBindPort = 8080
)

func init() {
	app.AddCommand(serverCmd)
}
