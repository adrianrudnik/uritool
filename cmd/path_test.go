package cmd

import (
	"github.com/adrianrudnik/cmd-uritools/cmd/cmdtest"
	"strings"
	"testing"
)

func TestPathHelp(t *testing.T) {
	setup()

	out, err := cmdtest.ExecuteCommand(rootCmd, "path")

	if err != nil {
		t.Errorf("Expected help, got error %s", err)
	}

	if !strings.Contains(out, "Usage:") {
		t.Errorf("Expected help, got %s", out)
	}
}

func TestPathEncodeHelp(t *testing.T) {
	setup()

	out, err := cmdtest.ExecuteCommand(rootCmd, "path", "encode")

	// throws error, missing arg
	if err == nil {
		t.Errorf("Expected help, got error %s", err)
	}

	// and shows help
	if !strings.Contains(out, "Usage:") {
		t.Errorf("Expected help, got %s", out)
	}
}

func TestPathDecodeHelp(t *testing.T) {
	setup()

	out, err := cmdtest.ExecuteCommand(rootCmd, "path", "decode")

	// throws error, missing arg
	if err == nil {
		t.Errorf("Expected help, got error %s", err)
	}

	// and shows help
	if !strings.Contains(out, "Usage:") {
		t.Errorf("Expected help, got %s", out)
	}
}

func TestPathEncodingWithValidValue(t *testing.T) {
	setup()

	in := "hello / world + %"
	expected := "hello%20%2F%20world%20+%20%25"
	out, _ := cmdtest.ExecuteCommand(rootCmd, "-n", "path", "encode", in)

	if out != expected {
		t.Errorf("Expected output is wrong: %s != %s", out, expected)
	}
}

func TestPathDecodingWithValidValue(t *testing.T) {
	setup()

	in := "hello%20%2F%20world%20%25"
	expected := "hello / world %"
	out, _ := cmdtest.ExecuteCommand(rootCmd, "-n", "path", "decode", in)

	if out != expected {
		t.Errorf("Expected output is wrong: %s != %s", out, expected)
	}
}

func TestPathDecodingWithInvalidValue(t *testing.T) {
	setup()

	in := "hello%%2Fworld"
	out, err := cmdtest.ExecuteCommand(rootCmd, "path", "decode", in)

	if err == nil {
		t.Errorf("Expected error, got %s", out)
	}
}
