package main

import (
	"fmt"

	"github.com/spf13/cobra"
)

var (
	local    string
	code     string
	remote   = "103.252.223.230:22"
	user     = "root"
	password = "07gn6aF7zfR1"
	script   string
	timeout  string
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
