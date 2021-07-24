package main

import (
	"fmt"
	"os"

	"github.com/spf13/cobra"
)

var (
	remote   string
	code     string
	user     string
	password string
	script   string
)

// RootCmd represents the base command when called without any subcommand
var RootCmd = &cobra.Command{
	Use:   "sshx",
	Short: "A replace tool for ssh features",
}

// Execute adds all child command to the root command sets flags appropriately.
func Execute() {
	if err := RootCmd.Execute(); err != nil {
		fmt.Println(err)
		os.Exit(-1)
	}
}

func init() {
	RootCmd.PersistentFlags().StringVarP(&remote, "remote", "r", "127.0.0.1:22", "the remote address")
	RootCmd.PersistentFlags().StringVarP(&code, "code", "c", "xxxx", "the pin code")
	RootCmd.PersistentFlags().StringVarP(&user, "user", "u", "root", "the user of remote server")
	RootCmd.PersistentFlags().StringVarP(&password, "password", "p", "123456", "the password of remote server")
}