package cmd

import (
	"github.com/spf13/cobra"
	"net/url"
)

var queryCmd = &cobra.Command{
	Use: "query",
	Run: func(cmd *cobra.Command, args []string) {
		_ = cmd.Help()
	},
}

var queryEncodeCmd = &cobra.Command{
	Use:  "encode",
	Args: cobra.ExactArgs(1),
	Run: func(cmd *cobra.Command, args []string) {
		out := url.QueryEscape(args[0])
		print(cmd, out)
	},
}

var queryDecodeCmd = &cobra.Command{
	Use:  "decode",
	Args: cobra.ExactArgs(1),
	RunE: func(cmd *cobra.Command, args []string) error {
		out, err := url.QueryUnescape(args[0])

		if err != nil {
			return err
		}

		print(cmd, out)
		return nil
	},
}
