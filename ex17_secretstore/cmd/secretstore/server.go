package main

import (
	"fmt"
	"github.com/duaraghav8/gophercises/ex17_secretstore/cmd/secretstore/apiserver"
	"github.com/duaraghav8/gophercises/ex17_secretstore/pkg/secretstore"
	"github.com/spf13/cobra"
	"net/http"
	"os"
)

var serverCmd = &cobra.Command{
	Use:   "server",
	Short: "Start Secretstore server",
	Long: `Start the server.
This server serves the REST API to allow users to interact
with Secretsmanager to store and fetch data.
The server manages the storage backend and all cryptographic
tasks.`,
	Run: startServerProcess,
}

var (
	serverBindAddr = "127.0.0.1"
	serverBindPort = 8080
)

func init() {
	app.AddCommand(serverCmd)
}

func startServerProcess(cmd *cobra.Command, args []string) {
	serverAddr := fmt.Sprintf("%s:%d", serverBindAddr, serverBindPort)

	server, err := apiserver.NewAPIServer(secretstore.NewInMemoryKVStore())
	if err != nil {
		fmt.Printf("Failed to initialize server: %v\n", err)
	}

	fmt.Printf("Starting Secretstore API server at http://%s\n", serverAddr)

	if err := http.ListenAndServe(serverAddr, server); err != nil {
		fmt.Printf("Failed to start API server: %v\n", err)
		os.Exit(1)
	}
}
