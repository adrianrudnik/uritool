package cmd

import (
	"github.com/spf13/cobra"
	"net/url"
)

var pathCmd = &cobra.Command{
	Use: "path",
	Run: func(cmd *cobra.Command, args []string) {
		_ = cmd.Help()
	},
}

var pathEncodeCmd = &cobra.Command{
	Use:  "encode",
	Args: cobra.ExactArgs(1),
	Run: func(cmd *cobra.Command, args []string) {
		out := url.PathEscape(args[0])
		output(cmd, out)
	},
}

var pathDecodeCmd = &cobra.Command{
	Use:  "decode",
	Args: cobra.ExactArgs(1),
	RunE: func(cmd *cobra.Command, args []string) error {
		out, err := url.PathUnescape(args[0])

		if err != nil {
			return err
		}

		output(cmd, out)
		return nil
	},
}
