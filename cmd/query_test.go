package cmd

import (
	"github.com/adrianrudnik/cmd-uritools/cmd/cmdtest"
	"strings"
	"testing"
)

func TestQueryHelp(t *testing.T) {
	setup()

	out, err := cmdtest.ExecuteCommand(rootCmd, "query")

	if err != nil {
		t.Errorf("Expected help, got error %s", err)
	}

	if !strings.Contains(out, "Usage:") {
		t.Errorf("Expected help, got %s", out)
	}
}

func TestQueryEncodeHelp(t *testing.T) {
	setup()

	out, err := cmdtest.ExecuteCommand(rootCmd, "query", "encode")

	// throws error, missing arg
	if err == nil {
		t.Errorf("Expected help, got error %s", err)
	}

	// and shows help
	if !strings.Contains(out, "Usage:") {
		t.Errorf("Expected help, got %s", out)
	}
}

func TestQueryDecodeHelp(t *testing.T) {
	setup()

	out, err := cmdtest.ExecuteCommand(rootCmd, "query", "decode")

	// throws error, missing arg
	if err == nil {
		t.Errorf("Expected help, got error %s", err)
	}

	// and shows help
	if !strings.Contains(out, "Usage:") {
		t.Errorf("Expected help, got %s", out)
	}
}

func TestQueryEncodingWithValidValue(t *testing.T) {
	setup()

	in := "hello / world%"
	expected := "hello+%2F+world%25"
	out, _ := cmdtest.ExecuteCommand(rootCmd, "-n", "query", "encode", in)

	if out != expected {
		t.Errorf("Expected output is wrong: %s != %s", out, expected)
	}
}

func TestQueryDecodingWithValidValue(t *testing.T) {
	setup()

	in := "hello+%2F+world%25"
	expected := "hello / world%"
	out, _ := cmdtest.ExecuteCommand(rootCmd, "-n", "query", "decode", in)

	if out != expected {
		t.Errorf("Expected output is wrong: %s != %s", out, expected)
	}
}

func TestQueryDecodingWithInvalidValue(t *testing.T) {
	setup()

	in := "hello%%2Fworld"
	out, err := cmdtest.ExecuteCommand(rootCmd, "query", "decode", in)

	if err == nil {
		t.Errorf("Expected error, got %s", out)
	}
}
