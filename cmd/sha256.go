/*
Copyright © 2024 NAME HERE <EMAIL ADDRESS>
*/
package cmd

import (
	"crypto/sha256"
	"fmt"

	"github.com/spf13/cobra"
)

// sha256Cmd represents the sha256 command
var sha256Cmd = &cobra.Command{
	Use:   "sha256",
	Short: "hashes a string with sha256",
	Long: `
sha256 hashing algorithm implemented in the go crypto/sha256 library,
it hashes the string passed as an arg.`,
	Run: hashString,
}

func hashString(cmd *cobra.Command, args []string) {
	if len(args) < 1 {
		fmt.Println("You must provide a string to use for the Sha256 checksum")
		return
	}

	s := args[0]
	sum := sha256.Sum256([]byte(s))
	fmt.Printf("%x\n", sum)
}

func init() {
	rootCmd.AddCommand(sha256Cmd)

	// Here you will define your flags and configuration settings.

	// Cobra supports Persistent Flags which will work for this command
	// and all subcommands, e.g.:
	// sha256Cmd.PersistentFlags().String("foo", "", "A help for foo")

	// Cobra supports local flags which will only run when this command
	// is called directly, e.g.:
	// sha256Cmd.Flags().BoolP("toggle", "t", false, "Help message for toggle")
}
