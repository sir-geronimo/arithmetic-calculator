package external

import (
	"errors"
	"fmt"
	"io"
	"net/http"
)

const baseUrl = "https://www.random.org/strings/"

var (
	ErrUnableToGenerateString = errors.New("unable to generate random string")
)

type StringGenerator interface {
	Generate(opts *GenerateStringOptions) (string, error)
}

type GenerateStringOptions struct {
	Len    uint8
	Unique bool
}

type HTTPStringGenerator struct {
	URL string
}

func NewHTTPStringGenerator() *HTTPStringGenerator {
	return &HTTPStringGenerator{
		URL: fmt.Sprintf("%s?num=1&upperalpha=on&loweralpha=on&digits=on&format=plain", baseUrl),
	}
}

func (g *HTTPStringGenerator) Generate(opts *GenerateStringOptions) (string, error) {
	url := fmt.Sprintf("%s&len=%v&unique=%v",
		g.URL,
		opts.Len,
		opts.Unique,
	)
	req, err := http.NewRequest(http.MethodGet, url, nil)
	if err != nil {
		return "", ErrUnableToGenerateString
	}

	res, err := http.DefaultClient.Do(req)
	if err != nil {
		return "", ErrUnableToGenerateString
	}

	body, err := io.ReadAll(res.Body)
	if err != nil {
		return "", ErrUnableToGenerateString
	}
	defer res.Body.Close()

	return string(body), nil
}
