package main

import (
	"bytes"
	"encoding/json"
	"fmt"
	"github.com/spf13/cobra"
	"io/ioutil"
	"net/http"
	"net/url"
	"os"
)

var putCmd = &cobra.Command{
	Use:   "put",
	Short: "Store or Update a secret",
	Run:   putSecret,
}

func init() {
	app.AddCommand(putCmd)
	putCmd.Flags().StringVarP(&bufSecretKey, "key", "k", "", "Secret key (Required)")
	putCmd.Flags().StringVarP(&bufSecretValue, "value", "v", "", "Secret value (Required)")
	putCmd.Flags().StringVarP(&bufEncodingKey, "encoding-key", "e", "", "Decryption Key (Required)")

	if err := putCmd.MarkFlagRequired("key"); err != nil {
		fmt.Println(err)
		os.Exit(1)
	}
	if err := putCmd.MarkFlagRequired("value"); err != nil {
		fmt.Println(err)
		os.Exit(1)
	}
	if err := putCmd.MarkFlagRequired("encoding-key"); err != nil {
		fmt.Println(err)
		os.Exit(1)
	}
}

func putSecret(cmd *cobra.Command, args []string) {
	client := &http.Client{}
	reqUrl := fmt.Sprintf("http://%s:%d/secret/%s",
		serverBindAddr, serverBindPort, url.QueryEscape(bufSecretKey))

	payload, err := json.Marshal(map[string]string{
		"value": bufSecretValue,
	})
	if err != nil {
		fmt.Printf("Unable to create request payload: %v\n", err)
		os.Exit(1)
	}

	req, _ := http.NewRequest(http.MethodPost, reqUrl, bytes.NewBuffer(payload))
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
