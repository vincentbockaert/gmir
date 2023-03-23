/*
Copyright © 2023 Vincent Bockaert <vincent.bockaert@pm.me>
*/
package cmd

import (
	"log"

	"encoding/base64"

	"github.com/hashicorp/vault/shamir"
	"github.com/spf13/cobra"
)

var splitCmd = &cobra.Command{
	Use:   "split",
	Short: "Split a secret into n shares, requiring the threshold amount of shares to recreate",
	Long: `Using HashiCorp's (Vault) implementation of Shamir Shared Secrets,
this command splits a secret into x so-called shares with y shares required 

to recreate the secret (i.e. threshold) with the combine command.
	`,
	Run: func(cmd *cobra.Command, args []string) {
		shares, err := shamir.Split([]byte(splitInput), splitParts, splitThreshold)
		if err != nil {
			log.Fatal(err)
		}
		for i := 0; i < len(shares); i++ {
			println(base64.StdEncoding.EncodeToString((shares[i])))
		}
	},
}

var splitInput string
var splitParts, splitThreshold int

func init() {
	rootCmd.AddCommand(splitCmd)

	splitCmd.Flags().StringVarP(&splitInput, "input", "i", "SomeSecretHere", "Enter a string as input `-i 'SuperSecret123!'`")
	splitCmd.Flags().IntVarP(&splitParts, "parts", "p", 5, "Number of parts to split the secret into (max 256)")
	splitCmd.Flags().IntVarP(&splitThreshold, "threshold", "t", 3, "Threshold of parts required in order to rebuild the secret (min 2)")
}
