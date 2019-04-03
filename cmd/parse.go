package cmd

import (
	"bytes"
	"encoding/json"
	"github.com/spf13/cobra"
	"html/template"
	"net/url"
)

var parseCmd = &cobra.Command{
	Use: "parse",
	Run: func(cmd *cobra.Command, args []string) {
		_ = cmd.Help()
	},
}

func doFormat(cmd *cobra.Command, data interface{}) (bool, error) {
	format, err := cmd.Flags().GetString("format")

	if err != nil {
		return true, err
	}

	if format != "" {
		tmpl, err := template.New("input").Parse(format)

		if err != nil {
			return true, err
		}

		var out bytes.Buffer

		if err := tmpl.Execute(&out, data); err != nil {
			return true, err
		}

		output(cmd, out.String())

		return true, nil
	}

	return false, nil
}

func doJson(cmd *cobra.Command, data interface{}) error {
	b, err := json.MarshalIndent(data, "", "  ")

	if err != nil {
		return err
	}

	output(cmd, string(b))

	return nil
}

var parseUriCmd = &cobra.Command{
	Use:  "uri",
	Args: cobra.ExactArgs(1),
	RunE: func(cmd *cobra.Command, args []string) error {
		parsed, err := url.Parse(args[0])

		if err != nil {
			return err
		}

		wrapped := NewUrl(parsed)

		done, err := doFormat(cmd, wrapped)

		if err != nil {
			return err
		}

		if done {
			return nil
		}

		if err := doJson(cmd, wrapped); err != nil {
			return err
		}

		return nil
	},
}

var parseQueryCmd = &cobra.Command{
	Use:  "query",
	Args: cobra.ExactArgs(1),
	RunE: func(cmd *cobra.Command, args []string) error {
		_, err := url.ParseQuery(args[0])

		if err != nil {
			return err
		}

		// @todo output(cmd, out)

		return nil
	},
}
