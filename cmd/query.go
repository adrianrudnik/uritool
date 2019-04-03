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

var queryEscapeCmd = &cobra.Command{
	Use:   "escape",
	Short: "escapes the given value and returns the valid query parameter value",
	Args:  cobra.ExactArgs(1),
	Run: func(cmd *cobra.Command, args []string) {
		out := url.QueryEscape(args[0])
		output(cmd, out)
	},
}

var queryUnescapeCmd = &cobra.Command{
	Use:   "unescape",
	Short: "unescapes the given escaped value and returns the original query parameter value",
	Args:  cobra.ExactArgs(1),
	RunE: func(cmd *cobra.Command, args []string) error {
		out, err := url.QueryUnescape(args[0])

		if err != nil {
			return err
		}

		output(cmd, out)
		return nil
	},
}
