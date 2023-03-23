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

var combineCmd = &cobra.Command{
	Use:   "combine",
	Short: "Combine n-shares to recreate a secret",
	Long: `
	Using HashiCorp's (Vault) implementation of Shamir Shared Secrets,
	this command will attempt to recreate a secret based on the provided shares inputted to.
	
	If you do not provide enough shares or provide one or more shares, the program will exit.
	`,
	Run: func(cmd *cobra.Command, args []string) {
		var err error
		var secret []byte

		combinePartsBytes = make([][]byte, len(combineParts))

		for i := 0; i < len(combineParts); i++ {
			combinePartsBytes[i], err = base64.StdEncoding.DecodeString(combineParts[i])
			if err != nil {
				log.Fatal(err)
			}
		}
		secret, err = shamir.Combine(combinePartsBytes)
		if err != nil {
			log.Fatal(err)
		}
		println(string(secret))
	},
}

var combineParts []string
var combinePartsBytes [][]byte

func init() {
	rootCmd.AddCommand(combineCmd)

	combineCmd.Flags().StringSliceVarP(&combineParts, "part", "p", []string{"part_one", "part_two"}, "Add the parts, `-p 'part_one' -p 'part_two'")
	combineCmd.MarkFlagRequired("part")
}
