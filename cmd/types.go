package cmd

import (
	"net/url"
	"strconv"
)

type URL struct {
	Scheme          string
	Opaque          string
	Username        string
	Password        string
	PasswordIsGiven bool
	Host            string
	Hostname        string
	Port            int
	Path            string
	PathEscaped     string
	RawQuery        string
	Fragment        string
	Query           url.Values
}

func NewUrl(in *url.URL) *URL {
	u := &URL{
		Scheme:      in.Scheme,
		Opaque:      in.Opaque,
		Username:    in.User.Username(),
		Host:        in.Host,
		Path:        in.Path,
		PathEscaped: in.EscapedPath(),
		Hostname:    in.Hostname(),
		RawQuery:    in.RawQuery,
		Fragment:    in.Fragment,
		Query:       in.Query(),
	}

	if in.Port() != "" {
		port, err := strconv.Atoi(in.Port())

		if err == nil {
			u.Port = port
		}
	}

	u.Password, u.PasswordIsGiven = in.User.Password()

	return u
}
