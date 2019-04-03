package cmd

import (
	"encoding/json"
	"github.com/adrianrudnik/uritool/cmd/cmdtest"
	"strings"
	"testing"
)

func TestParseHelp(t *testing.T) {
	setup()

	out, err := cmdtest.ExecuteCommand(rootCmd, "parse")

	if err != nil {
		t.Errorf("Expected help, got error %s", err)
	}

	if !strings.Contains(out, "Usage:") {
		t.Errorf("Expected help, got %s", out)
	}
}

func TestParseUriHelp(t *testing.T) {
	setup()

	out, err := cmdtest.ExecuteCommand(rootCmd, "parse", "uri")

	// throws error, missing arg
	if err == nil {
		t.Errorf("Expected help, got error %s", err)
	}

	// and shows help
	if !strings.Contains(out, "Usage:") {
		t.Errorf("Expected help, got %s", out)
	}
}

func TestParseQueryHelp(t *testing.T) {
	setup()

	out, err := cmdtest.ExecuteCommand(rootCmd, "parse", "query")

	// throws error, missing arg
	if err == nil {
		t.Errorf("Expected help, got error %s", err)
	}

	// and shows help
	if !strings.Contains(out, "Usage:") {
		t.Errorf("Expected help, got %s", out)
	}
}

func TestParseValidUri(t *testing.T) {
	setup()

	in := "xyz://my:pass@the.example.com:8080/what/ ever?this=is&this=isnot#doing"
	out, _ := cmdtest.ExecuteCommand(rootCmd, "-n", "parse", "uri", in)

	var result URL
	err := json.Unmarshal([]byte(out), &result)

	if err != nil {
		t.Errorf("Could not validate JSON: %s", err)
	}

	expectedScheme := "xyz"
	if result.Scheme != expectedScheme {
		t.Errorf("Expected scheme is wrong: %s != %s", result.Scheme, expectedScheme)
	}

	expectedOpaque := ""
	if result.Opaque != expectedOpaque {
		t.Errorf("Expected opaque is not empty: %s", result.Opaque)
	}

	expectedUsername := "my"
	if result.Username != expectedUsername {
		t.Errorf("Expected username is wrong: %s != %s", result.Username, expectedUsername)
	}

	expectedPassword := "pass"
	if result.Password != expectedPassword {
		t.Errorf("Expected password is wrong: %s != %s", result.Password, expectedPassword)
	}

	expectedHost := "the.example.com:8080"
	if result.Host != expectedHost {
		t.Errorf("Expected host is wrong: %s != %s", result.Host, expectedHost)
	}

	expectedHostname := "the.example.com"
	if result.Hostname != expectedHostname {
		t.Errorf("Expected hostname is wrong: %s != %s", result.Hostname, expectedHostname)
	}

	expectedPort := 8080
	if result.Port != expectedPort {
		t.Errorf("Expected port is wrong: %d != %d", result.Port, expectedPort)
	}

	expectedPath := "/what/ ever"
	if result.Path != expectedPath {
		t.Errorf("Expected path is wrong: %s != %s", result.Path, expectedPath)
	}

	exptectedEscapedPath := "/what/%20ever"
	if result.PathEscaped != exptectedEscapedPath {
		t.Errorf("Expected escaped path is wrong: %s != %s", result.PathEscaped, exptectedEscapedPath)
	}

	expectedRawQuery := "this=is\u0026this=isnot"
	if result.RawQuery != expectedRawQuery {
		t.Errorf("Expected raw query is wrong: %s != %s", result.RawQuery, expectedRawQuery)
	}

	expectedFragment := "doing"
	if result.Fragment != expectedFragment {
		t.Errorf("Expected fragment is wrong: %s != %s", result.Fragment, expectedFragment)
	}

	if len(result.Query) != 1 {
		t.Errorf("Expected query parameter count is wrong: %d != %d", len(result.Query), 1)
	}

	if len(result.Query.Get("this")) != 2 {
		t.Errorf("Expected query parameter array count is wrong: %d != %d", len(result.Query.Get("this")), 2)
	}

	if result.Query["this"][0] != "is" {
		t.Errorf("Expected query parameter array element 0 is wrong: %s != %s", result.Query["this"][0], "is")
	}

	if result.Query["this"][1] != "isnot" {
		t.Errorf("Expected query parameter array element 0 is wrong: %s != %s", result.Query["this"][0], "isnot")
	}
}

func TestParseValidOpaqueUri(t *testing.T) {
	setup()

	in := "xyz:my:pass@the.example.com:8080/what/ ever?this=is&this=isnot#doing"
	out, _ := cmdtest.ExecuteCommand(rootCmd, "-n", "parse", "uri", in)

	var result URL
	err := json.Unmarshal([]byte(out), &result)

	if err != nil {
		t.Errorf("Could not validate JSON: %s", err)
	}

	expectedScheme := "xyz"
	if result.Scheme != expectedScheme {
		t.Errorf("Expected scheme is wrong: %s != %s", result.Scheme, expectedScheme)
	}

	expectedOpaque := "my:pass@the.example.com:8080/what/ ever"
	if result.Opaque != expectedOpaque {
		t.Errorf("Expected opaque is not empty: %s", result.Opaque)
	}

	expectedUsername := ""
	if result.Username != expectedUsername {
		t.Errorf("Expected username is wrong: %s != %s", result.Username, expectedUsername)
	}

	expectedPassword := ""
	if result.Password != expectedPassword {
		t.Errorf("Expected password is wrong: %s != %s", result.Password, expectedPassword)
	}

	expectedHost := ""
	if result.Host != expectedHost {
		t.Errorf("Expected host is wrong: %s != %s", result.Host, expectedHost)
	}

	expectedHostname := ""
	if result.Hostname != expectedHostname {
		t.Errorf("Expected hostname is wrong: %s != %s", result.Hostname, expectedHostname)
	}

	expectedPort := 0
	if result.Port != expectedPort {
		t.Errorf("Expected port is wrong: %d != %d", result.Port, expectedPort)
	}

	expectedPath := ""
	if result.Path != expectedPath {
		t.Errorf("Expected path is wrong: %s != %s", result.Path, expectedPath)
	}

	exptectedEscapedPath := ""
	if result.PathEscaped != exptectedEscapedPath {
		t.Errorf("Expected escaped path is wrong: %s != %s", result.PathEscaped, exptectedEscapedPath)
	}

	expectedRawQuery := "this=is\u0026this=isnot"
	if result.RawQuery != expectedRawQuery {
		t.Errorf("Expected raw query is wrong: %s != %s", result.RawQuery, expectedRawQuery)
	}

	expectedFragment := "doing"
	if result.Fragment != expectedFragment {
		t.Errorf("Expected fragment is wrong: %s != %s", result.Fragment, expectedFragment)
	}

	if len(result.Query) != 1 {
		t.Errorf("Expected query parameter count is wrong: %d != %d", len(result.Query), 1)
	}

	if len(result.Query.Get("this")) != 2 {
		t.Errorf("Expected query parameter array count is wrong: %d != %d", len(result.Query.Get("this")), 2)
	}

	if result.Query["this"][0] != "is" {
		t.Errorf("Expected query parameter array element 0 is wrong: %s != %s", result.Query["this"][0], "is")
	}

	if result.Query["this"][1] != "isnot" {
		t.Errorf("Expected query parameter array element 0 is wrong: %s != %s", result.Query["this"][0], "isnot")
	}
}

func TestParseValidUriWithFormat(t *testing.T) {
	setup()

	var format strings.Builder
	format.WriteString("--format=")
	format.WriteString("\"")
	format.WriteString("{{.Scheme}}#")
	format.WriteString("{{.Opaque}}#")
	format.WriteString("{{.Username}}#")
	format.WriteString("{{.Password}}#")
	format.WriteString("{{.PasswordIsGiven}}#")
	format.WriteString("{{.Host}}#")
	format.WriteString("{{.Hostname}}#")
	format.WriteString("{{.Port}}#")
	format.WriteString("{{.Path}}#")
	format.WriteString("{{.PathEscaped}}#")
	format.WriteString("{{.RawQuery}}#")
	format.WriteString("{{.Fragment}}#")
	format.WriteString("{{index .Query.this 0}}#")
	format.WriteString("{{index .Query.this 1}}#")
	format.WriteString("\"")

	in := "xyz://my:pass@the.example.com:8080/what/ ever?this=is&this=isnot#doing"
	expected := "\"xyz##my#pass#true#the.example.com:8080#the.example.com#8080#/what/ ever#/what/%20ever#this=is&amp;this=isnot#doing#is#isnot#\""
	out, _ := cmdtest.ExecuteCommand(rootCmd, "-n", format.String(), "parse", "uri", in)

	if out != expected {
		t.Errorf("Expected format is wrong: %s != %s", out, expected)
	}
}

func TestParseStringWithFormat(t *testing.T) {
	setup()

	in := "http://www.example.com/"
	expected := "www.example.com"
	out, _ := cmdtest.ExecuteCommand(rootCmd, "-n", "--format={{.Hostname}}", "parse", "uri", in)

	if out != expected {
		t.Errorf("Expected format is wrong: %s != %s", out, expected)
	}
}
