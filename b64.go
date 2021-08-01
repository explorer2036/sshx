package main

import (
	"encoding/base64"
	"fmt"

	"github.com/spf13/cobra"
)

var base64Cmd = &cobra.Command{
	Use:   "b64",
	Short: "base64 encode and decode",
	Run: func(cmd *cobra.Command, args []string) {
		if encode {
			encoded := base64.StdEncoding.EncodeToString([]byte(source))
			fmt.Printf("after encode: \"%s\" -> \"%s\"\n", source, encoded)
		} else {
			decoded, err := base64.StdEncoding.DecodeString(source)
			if err != nil {
				panic(err)
			}
			fmt.Printf("after decode: \"%s\" -> \"%s\"\n", source, string(decoded))
		}
	},
}

func init() {
	// RootCmd.AddCommand(base64Cmd)
	base64Cmd.PersistentFlags().StringVarP(&source, "source", "s", "hi", "the source string for encoding/decoding")
	base64Cmd.PersistentFlags().BoolVarP(&encode, "encode", "e", false, "encode or decode")
}
