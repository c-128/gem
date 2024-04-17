package gem

import (
	"errors"
	"net/url"
)

// A gemini URL.
type URL struct {
	Scheme string
	Host   string
	Path   string
	Query  string
}

// Parse a URL.
// Only supports the "gemini" scheme.
func ParseURL(raw string) (*URL, error) {
	parsedURL, err := url.Parse(raw)
	if err != nil {
		return nil, err
	}

	if parsedURL.Scheme != geminiScheme {
		return nil, errors.New("not a gemini URL")
	}

	query, err := url.QueryUnescape(parsedURL.RawQuery)
	if err != nil {
		return nil, err
	}

	url := &URL{
		Scheme: parsedURL.Scheme,
		Host:   parsedURL.Host,
		Path:   parsedURL.Path,
		Query:  query,
	}
	return url, nil
}
