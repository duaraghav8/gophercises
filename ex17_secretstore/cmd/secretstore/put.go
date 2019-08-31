package main

import (
	"fmt"
	"github.com/spf13/cobra"
	"os"
)

func init() {
	app.AddCommand(putCmd)
	putCmd.Flags().StringVarP(&encodingKey, "key", "k", "", "Secret key (Required)")
	putCmd.Flags().StringVarP(&encodingKey, "value", "v", "", "Secret value (Required)")
	putCmd.Flags().StringVarP(&encodingKey, "encoding-key", "e", "", "Decryption Key (Required)")

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

var putCmd = &cobra.Command{
	Use: "put",
	Short: "Store or Update a secret",
	Run: func(cmd *cobra.Command, args []string) {},
}
