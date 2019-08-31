package main

import "github.com/spf13/cobra"

func init() {
	app.AddCommand(listSecretKeysCmd)
}

var listSecretKeysCmd = &cobra.Command{
	Use: "list-secrets",
	Short: "List all secrets",
	Run: func(cmd *cobra.Command, args []string) {},
}
