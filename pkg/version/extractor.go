package version

import (
	"errors"
	"fmt"
	"net/url"
)


func assertInputUrl(u string) error {
	_, err := url.ParseRequestURI(u)
	if err != nil {
		msg := fmt.Sprintf("Invalid URL: %s", u)
		return errors.New(msg)
	}
	return nil
}

type Transport interface {
	DoCall(string) (string, error)
}

type PostProcessor interface {
	Process(string) string
}

type Extractor struct {
	url       string
	transport Transport
	post      PostProcessor
}

func NewExtractor() *Extractor {
	return &Extractor{}
}

func (e *Extractor) WithTransport(transport Transport) *Extractor {
	e.transport = transport

	return e
}

func (e *Extractor) FromURL(url string) *Extractor {
	e.url = url
	return e
}

func (e *Extractor) WithPostProcessor(post PostProcessor) *Extractor {
	e.post = post
	return e
}

func (e *Extractor) Extract() (string, error) {

	if err := assertInputUrl(e.url); err != nil {
		return "", err
	}

	v, err := e.transport.DoCall(e.url)
	if err != nil {
		return "", err
	}

	if e.post != nil {
		v = e.post.Process(v)
	}

	return v, nil
}
