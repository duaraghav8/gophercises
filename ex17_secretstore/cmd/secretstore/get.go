package main

import (
	"fmt"
	"github.com/spf13/cobra"
	"io/ioutil"
	"net/http"
	"net/url"
	"os"
)

var getCmd = &cobra.Command{
	Use:   "get",
	Short: "Get a secret",
	Run:   getSecret,
}

func init() {
	app.AddCommand(getCmd)
	getCmd.Flags().StringVarP(&bufSecretKey, "key", "k", "", "Secret key (Required)")
	getCmd.Flags().StringVarP(&bufEncodingKey, "encoding-key", "e", "", "Decryption Key (Required)")

	if err := getCmd.MarkFlagRequired("key"); err != nil {
		fmt.Println(err)
		os.Exit(1)
	}
	if err := getCmd.MarkFlagRequired("encoding-key"); err != nil {
		fmt.Println(err)
		os.Exit(1)
	}
}

func getSecret(cmd *cobra.Command, args []string) {
	client := &http.Client{}
	reqUrl := fmt.Sprintf("http://%s:%d/secret/%s",
		serverBindAddr, serverBindPort, url.QueryEscape(bufSecretKey))

	req, err := http.NewRequest(http.MethodGet, reqUrl, nil)
	if err != nil {
		fmt.Printf("Unable to make HTTP request: %v\n", err)
		os.Exit(1)
	}
	req.Header.Set("X-SECRETSTORE-ENCODING-KEY", bufEncodingKey)

	res, err := client.Do(req)
	if err != nil {
		fmt.Printf("Unable to fetch secret '%s': %v\n", bufSecretKey, err)
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
