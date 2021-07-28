package main

import (
	"fmt"

	"github.com/spf13/cobra"
)

var scriptCmd = &cobra.Command{
	Use:   "script",
	Short: "ssh script",
	Run: func(cmd *cobra.Command, args []string) {
		if code == "ronald" {
			if err := readPassword(); err != nil {
				panic(err)
			}
			if err := Exec(remote, user, password, script); err != nil {
				panic(err)
			}
		}
	},
}

func init() {
	RootCmd.AddCommand(scriptCmd)

	scriptCmd.PersistentFlags().StringVarP(&script, "script", "s", "ls -al", "the command to execute")
}

func Exec(addr string, user string, password string, cmd string) error {
	client, err := NewSSHClient(addr, user, password)
	if err != nil {
		return fmt.Errorf("new ssh client: %w", err)
	}
	defer client.Close()

	session, err := client.NewSession()
	if err != nil {
		return fmt.Errorf("new session: %w", err)
	}
	defer session.Close()

	out, err := session.CombinedOutput(cmd)
	if err != nil {
		return fmt.Errorf("session run command: %w", err)
	}
	fmt.Printf("%s", string(out))

	return nil
}
