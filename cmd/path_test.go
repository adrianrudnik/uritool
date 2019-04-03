package cmd

import (
	"github.com/adrianrudnik/uritool/cmd/cmdtest"
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

func TestPathEscapeHelp(t *testing.T) {
	setup()

	out, err := cmdtest.ExecuteCommand(rootCmd, "path", "escape")

	// throws error, missing arg
	if err == nil {
		t.Errorf("Expected help, got error %s", err)
	}

	// and shows help
	if !strings.Contains(out, "Usage:") {
		t.Errorf("Expected help, got %s", out)
	}
}

func TestPathUnescapeHelp(t *testing.T) {
	setup()

	out, err := cmdtest.ExecuteCommand(rootCmd, "path", "unescape")

	// throws error, missing arg
	if err == nil {
		t.Errorf("Expected help, got error %s", err)
	}

	// and shows help
	if !strings.Contains(out, "Usage:") {
		t.Errorf("Expected help, got %s", out)
	}
}

func TestPathEscapeWithValidValue(t *testing.T) {
	setup()

	in := "hello / world + %"
	expected := "hello%20%2F%20world%20+%20%25"
	out, _ := cmdtest.ExecuteCommand(rootCmd, "-n", "path", "escape", in)

	if out != expected {
		t.Errorf("Expected output is wrong: %s != %s", out, expected)
	}
}

func TestPathUnescapeWithValidValue(t *testing.T) {
	setup()

	in := "hello%20%2F%20world%20%25"
	expected := "hello / world %"
	out, _ := cmdtest.ExecuteCommand(rootCmd, "-n", "path", "unescape", in)

	if out != expected {
		t.Errorf("Expected output is wrong: %s != %s", out, expected)
	}
}

func TestPathUnescapeWithInvalidValue(t *testing.T) {
	setup()

	in := "hello%%2Fworld"
	out, err := cmdtest.ExecuteCommand(rootCmd, "path", "unescape", in)

	if err == nil {
		t.Errorf("Expected error, got %s", out)
	}
}
