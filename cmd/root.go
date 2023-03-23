/*
Copyright © 2023 Vincent Bockaert <vincent.bockaert@pm.me>
*/
package cmd

import (
	"os"

	"github.com/spf13/cobra"
)

var rootCmd = &cobra.Command{
	Use:   "gmir",
	Short: "Small cli tool to split/combine secrets based on Shamir Shared Secrets (HashiCorp's Vault implementation)",
	Long: `
	A small command line tool which allows you split a securely secret up into x number of shares,
	with only y (>x) number of shares being required to recreate to the secret.

	All credit for the crypto-work goes to the HashiCorp, this tool is merely a tiny wrapper around it.
	`,
}

func Execute() {
	err := rootCmd.Execute()
	if err != nil {
		os.Exit(1)
	}
}

func init() {
	rootCmd.Flags().BoolP("toggle", "t", false, "Help message for toggle")
}
