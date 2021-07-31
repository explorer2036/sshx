package main

import (
	"fmt"

	"github.com/spf13/cobra"
)

var (
	input   string
	script  string
	timeout string
	encode  bool
	source  string

	local         = "MTI3LjAuMC4xOjEyMzQ="
	remote        = "MzEuMjIwLjEuMTY6MjIyNDg="
	server        = "ODkuMjQ5LjQ5LjEyODozMzMzMw=="
	user          = "Y29ubmVjdGVzdA=="
	password      = "bmFraTEyMzQ="
	code          = "cm9uYWxk"
	publicKeyFile = "fi8uc3NoL2lkX3JzYQ=="
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
