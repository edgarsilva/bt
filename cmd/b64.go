/*
Copyright © 2024 NAME HERE <EMAIL ADDRESS>
*/
package cmd

import (
	"encoding/base64"
	"fmt"

	"github.com/spf13/cobra"
)

// b64Cmd represents the b64 command
var b64Cmd = &cobra.Command{
	Use:   "b64",
	Short: "Encodes/Decodes b64 strings",
	Long: `Example:

	$ jbot b64 some-string
	# c29tZS1zdHJpbmc=
	$ jbot b64 c29tZS1zdHJpbmc= -d
	# some-string`,
	Run: b64,
}

func init() {
	rootCmd.AddCommand(b64Cmd)

	// Here you will define your flags and configuration settings.

	// Cobra supports Persistent Flags which will work for this command
	// and all subcommands, e.g.:
	// b64Cmd.PersistentFlags().String("decode", "d", "-d ")

	// Cobra supports local flags which will only run when this command
	// is called directly, e.g.:
	b64Cmd.Flags().BoolP("decode", "d", false, "Decode b64 string instead")
}

func b64(cmd *cobra.Command, args []string) {
	if len(args) < 1 {
		fmt.Println("You must provide a string to encode in base64")
		return
	}

	data := args[0]
	decode, _ := cmd.Flags().GetBool("decode")

	sEnc := base64.StdEncoding.EncodeToString([]byte(data))
	if decode {
		sDec, err := base64.StdEncoding.DecodeString(data)
		if err != nil {
			fmt.Println("Error found while decoding '", data, "'", err)
			return
		}
		sEnc = string(sDec)
	}

	fmt.Println(sEnc)
}
