package main

import (
	"fmt"
	"github.com/spf13/cobra"
	"io/ioutil"
	"net/http"
	"os"
)

var listSecretsCmd = &cobra.Command{
	Use:   "list-secrets",
	Short: "List all secrets",
	Run:   listSecrets,
}

func init() {
	app.AddCommand(listSecretsCmd)
}

func listSecrets(cmd *cobra.Command, args []string) {
	reqUrl := fmt.Sprintf("http://%s:%d/keys", serverBindAddr, serverBindPort)

	res, err := http.Get(reqUrl)
	if err != nil {
		fmt.Printf("Unable to fetch the list of secrets: %v\n", err)
		os.Exit(1)
	}
	defer res.Body.Close()

	body, err := ioutil.ReadAll(res.Body)
	if err != nil {
		fmt.Printf("Problem read response: %v\n", err)
		os.Exit(1)
	}

	fmt.Println(string(body))
}
