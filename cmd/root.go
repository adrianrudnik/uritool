package cmd

import (
	"fmt"
	"github.com/spf13/cobra"
	"os"
)

var Version = "1.0.0"
var NoTrailingNewlineFlag bool

func Execute() {
	if err := rootCmd.Execute(); err != nil {
		fmt.Println(err)
		os.Exit(1)
	}
}

func setup() {
	rootCmd.ResetFlags()
	rootCmd.ResetCommands()

	queryCmd.AddCommand(queryEncodeCmd)
	queryCmd.AddCommand(queryDecodeCmd)

	rootCmd.AddCommand(versionCmd)
	rootCmd.AddCommand(queryCmd)

	// add global n
	rootCmd.PersistentFlags().BoolVarP(&NoTrailingNewlineFlag, "no-newline", "n", false, "do not output the trailing newline")
}

func init() {
	setup()
}

func print(cmd *cobra.Command, val string) {
	if !NoTrailingNewlineFlag {
		cmd.Println(val)
	} else {
		cmd.Printf("%s", val)
	}
}

var rootCmd = &cobra.Command{
	Use:   "uritool",
	Short: "uritool is a helper command to process uris on the command line",
	Run: func(cmd *cobra.Command, args []string) {
		_ = cmd.Help()
	},
}

var versionCmd = &cobra.Command{
	Use:   "version",
	Short: "prints the current version",
	Run: func(cmd *cobra.Command, args []string) {
		print(cmd, "1.0.0")
	},
}
