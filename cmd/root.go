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
	queryCmd.AddCommand(queryEncodeCmd)
	queryCmd.AddCommand(queryDecodeCmd)

	pathCmd.AddCommand(pathEncodeCmd)
	pathCmd.AddCommand(pathDecodeCmd)

	// parse uri
	ParseUriCmd.ResetFlags()
	ParseUriCmd.ResetCommands()
	ParseUriCmd.Flags().String("format", "", "use go template for formatted output")

	// parse query
	parseQueryCmd.ResetFlags()
	parseQueryCmd.ResetCommands()
	parseQueryCmd.Flags().String("format", "", "use go template for formatted output")

	parseCmd.AddCommand(parseQueryCmd)
	parseCmd.AddCommand(ParseUriCmd)

	rootCmd.ResetFlags()
	rootCmd.ResetCommands()
	rootCmd.PersistentFlags().BoolVarP(&NoTrailingNewlineFlag, "no-newline", "n", false, "do not output the trailing newline")
	rootCmd.AddCommand(versionCmd)
	rootCmd.AddCommand(queryCmd)
	rootCmd.AddCommand(pathCmd)
	rootCmd.AddCommand(parseCmd)
}

func init() {
	setup()
}

func output(cmd *cobra.Command, val string) {
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
		output(cmd, "1.0.0")
	},
}
