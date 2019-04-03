package cmd

import (
	"fmt"
	"github.com/adrianrudnik/cmd-uritools/cmd/cmdtest"
	"strings"
	"testing"
)

func TestHelp(t *testing.T) {
	setup()

	out, err := cmdtest.ExecuteCommand(rootCmd)

	if err != nil {
		t.Errorf("Expected help, got error %s", err)
	}

	if !strings.Contains(out, "Usage:") {
		t.Errorf("Expected help, got %s", out)
	}
}

func TestVersionWithoutNewline(t *testing.T) {
	setup()

	expected := Version
	out, _ := cmdtest.ExecuteCommand(rootCmd, "-n", "version")

	if out != expected {
		t.Errorf("Expected output is wrong: %s != %s", out, expected)
	}
}

func TestVersionWithNewline(t *testing.T) {
	setup()

	expected := fmt.Sprintf("%s\n", Version)
	out, _ := cmdtest.ExecuteCommand(rootCmd, "version")

	if out != expected {
		t.Errorf("Expected output is wrong: %s != %s", out, expected)
	}
}
