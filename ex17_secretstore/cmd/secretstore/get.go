package main

import (
	"fmt"
	"github.com/spf13/cobra"
	"os"
)

func init() {
	app.AddCommand(getCmd)
	getCmd.Flags().StringVarP(&encodingKey, "key", "k", "", "Secret key (Required)")
	getCmd.Flags().StringVarP(&encodingKey, "encoding-key", "e", "", "Decryption Key (Required)")

	if err := getCmd.MarkFlagRequired("key"); err != nil {
		fmt.Println(err)
		os.Exit(1)
	}
	if err := getCmd.MarkFlagRequired("encoding-key"); err != nil {
		fmt.Println(err)
		os.Exit(1)
	}
}

var getCmd = &cobra.Command{
	Use: "get",
	Short: "Get a secret",
	Run: func(cmd *cobra.Command, args []string) {},
}
