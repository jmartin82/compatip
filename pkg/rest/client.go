package rest

import (
	"errors"
	"fmt"
	"io/ioutil"
	"net/http"
	"strings"
)

type HTTPGetter interface {
	Get(string) (resp *http.Response, err error)
}

type RESTTransport struct {
	client HTTPGetter
}

func NewRESTTransport(client HTTPGetter) *RESTTransport {
	return &RESTTransport{client: client}
}

func (re *RESTTransport) DoCall(url string) (string, error) {

	resp, err := re.client.Get(url)
	if err != nil {
		return "", err
	}

	if resp.StatusCode != http.StatusOK {
		msg := fmt.Sprintf("Invalid server response (%d)", resp.StatusCode)
		return "", errors.New(msg)
	}

	defer resp.Body.Close()
	body, err := ioutil.ReadAll(resp.Body)

	if err != nil {
		return "", err
	}

	version := string(body)
	return strings.TrimSpace(version), nil

}
