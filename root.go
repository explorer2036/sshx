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

	local      = "MTI3LjAuMC4xOjEyMzQ="
	remote     = "MzEuMjIwLjEuMTY6MjIyNDg="
	server     = "ODkuMjQ5LjQ5LjEyODozMzMzMw=="
	user       = "Y29ubmVjdGVzdA=="
	password   = "bmFraTEyMzQ="
	code       = "cm9uYWxk"
	privateKey = `
-----BEGIN OPENSSH PRIVATE KEY-----
b3BlbnNzaC1rZXktdjEAAAAABG5vbmUAAAAEbm9uZQAAAAAAAAABAAAAMwAAAAtzc2gtZW
QyNTUxOQAAACBpG77lm5Wb7KjK1mxkMHRfoRbN46pw6akdpfUFcrkInwAAAJhKyfNKSsnz
SgAAAAtzc2gtZWQyNTUxOQAAACBpG77lm5Wb7KjK1mxkMHRfoRbN46pw6akdpfUFcrkInw
AAAEC+BFzNTmVWPJ7HfjV/MSxvyYyw0Xs7cAoIEJl17JAA82kbvuWblZvsqMrWbGQwdF+h
Fs3jqnDpqR2l9QVyuQifAAAAEGNvbm5lY3Rlc3RAY2VudDgBAgMEBQ==
-----END OPENSSH PRIVATE KEY-----
`
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
