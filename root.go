package main

import (
	"fmt"

	"github.com/spf13/cobra"
)

var (
	code    string
	script  string
	timeout string
	tunnel  bool

	local    = "127.0.0.1:2222"
	remote   = "103.252.223.230:22"
	user     = "root"
	password = "07gn6aF7zfR1"
)

// RootCmd represents the base command when called without any subcommand
var RootCmd = &cobra.Command{
	Use:   "sshx",
	Short: "A replace tool for ssh features",
}

// Execute adds all child command to the root command sets flags appropriately.
func Execute() {
	if err := RootCmd.Execute(); err != nil {
		panic(fmt.Sprintf("root command execute: %v", err))
	}
}
