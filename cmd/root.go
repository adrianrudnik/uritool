package cmd

import (
	"fmt"
	"github.com/spf13/cobra"
	"os"
)

var Version = "1.0.0"
var NoTrailingNewlineFlag bool

func Execute() {
	rootCmd.SetOutput(os.Stdout)

	if err := rootCmd.Execute(); err != nil {
		fmt.Println(err)
		os.Exit(1)
	}
}

func setup() {
	// query
	queryCmd.ResetFlags()
	queryCmd.ResetCommands()

	// query escape
	queryEscapeCmd.ResetFlags()
	queryEscapeCmd.ResetCommands()
	queryCmd.AddCommand(queryEscapeCmd)

	// query unescape
	queryUnescapeCmd.ResetFlags()
	queryUnescapeCmd.ResetCommands()
	queryCmd.AddCommand(queryUnescapeCmd)

	// path
	pathCmd.ResetFlags()
	pathCmd.ResetCommands()

	// path escape
	pathEscapeCmd.ResetFlags()
	pathEscapeCmd.ResetCommands()
	pathCmd.AddCommand(pathEscapeCmd)

	pathUnescapeCmd.ResetFlags()
	pathUnescapeCmd.ResetCommands()
	pathCmd.AddCommand(pathUnescapeCmd)

	// parse
	parseCmd.ResetFlags()
	parseCmd.ResetCommands()

	// parse uri
	parseUriCmd.ResetFlags()
	parseUriCmd.ResetCommands()
	parseUriCmd.Flags().String("format", "", "use go template for formatted output")
	parseCmd.AddCommand(parseUriCmd)

	// parse query
	parseQueryCmd.ResetFlags()
	parseQueryCmd.ResetCommands()
	parseQueryCmd.Flags().String("format", "", "use go template for formatted output")
	parseCmd.AddCommand(parseQueryCmd)

	// root
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
		cmd.OutOrStdout()
		cmd.Println(val)
	} else {
		cmd.Printf("%s", val)
	}
}

var rootCmd = &cobra.Command{
	Use:   "uritool",
	Short: "process uris and extract information",
	Run: func(cmd *cobra.Command, args []string) {
		_ = cmd.Help()
	},
}

var versionCmd = &cobra.Command{
	Use:   "version",
	Short: "prints the current version",
	Run: func(cmd *cobra.Command, args []string) {
		output(cmd, Version)
	},
}
