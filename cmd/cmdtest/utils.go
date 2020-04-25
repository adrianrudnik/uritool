package cmdtest

import (
	"bytes"
	"github.com/spf13/cobra"
)

// ExecuteCommand executes and captures the output, returning the output and error
func ExecuteCommand(root *cobra.Command, args ...string) (output string, err error) {
	_, output, err = ExecuteCommandC(root, args...)
	return output, err
}

// ExecuteCommand executes and captures the output, returning the used command and output and error
func ExecuteCommandC(root *cobra.Command, args ...string) (c *cobra.Command, output string, err error) {
	buf := new(bytes.Buffer)
	root.SetOutput(buf)
	root.SetArgs(args)

	c, err = root.ExecuteC()

	return c, buf.String(), err
}
