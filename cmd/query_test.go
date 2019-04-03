package cmd

import (
	"github.com/adrianrudnik/uritool/cmd/cmdtest"
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

func TestQueryEscapeHelp(t *testing.T) {
	setup()

	out, err := cmdtest.ExecuteCommand(rootCmd, "query", "escape")

	// throws error, missing arg
	if err == nil {
		t.Errorf("Expected help, got error %s", err)
	}

	// and shows help
	if !strings.Contains(out, "Usage:") {
		t.Errorf("Expected help, got %s", out)
	}
}

func TestQueryUnescapeHelp(t *testing.T) {
	setup()

	out, err := cmdtest.ExecuteCommand(rootCmd, "query", "unescape")

	// throws error, missing arg
	if err == nil {
		t.Errorf("Expected help, got error %s", err)
	}

	// and shows help
	if !strings.Contains(out, "Usage:") {
		t.Errorf("Expected help, got %s", out)
	}
}

func TestQueryEscapeWithValidValue(t *testing.T) {
	setup()

	in := "hello / world%"
	expected := "hello+%2F+world%25"
	out, _ := cmdtest.ExecuteCommand(rootCmd, "-n", "query", "escape", in)

	if out != expected {
		t.Errorf("Expected output is wrong: %s != %s", out, expected)
	}
}

func TestQueryUnescapeWithValidValue(t *testing.T) {
	setup()

	in := "hello+%2F+world%25"
	expected := "hello / world%"
	out, _ := cmdtest.ExecuteCommand(rootCmd, "-n", "query", "unescape", in)

	if out != expected {
		t.Errorf("Expected output is wrong: %s != %s", out, expected)
	}
}

func TestQueryUnescapeWithInvalidValue(t *testing.T) {
	setup()

	in := "hello%%2Fworld"
	out, err := cmdtest.ExecuteCommand(rootCmd, "query", "unescape", in)

	if err == nil {
		t.Errorf("Expected error, got %s", out)
	}
}
