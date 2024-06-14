package client

import (
	"encoding/json"
	"io"
	"net/http"
	"net/url"
	"path"
	"strconv"
)

type Client struct {
	host     string
	BasePath string
	client   http.Client
}

func New(host string, token string) *Client {
	return &Client{
		host:     host,
		BasePath: "bot" + token,
		client:   http.Client{},
	}
}

func (c *Client) Updates(offset int, limit int) ([]Update, error) {

	query := url.Values{}
	query.Add("offset", strconv.Itoa(offset))
	query.Add("limit", strconv.Itoa(limit))

	data, err := c.Request("getUpdates", query)
	if err != nil {
		return nil, err
	}

	var result UpdatesResponse

	if err := json.Unmarshal(data, &result); err != nil {
		return nil, err
	}

	return result.Result, nil
}

func (c *Client) Request(metod string, query url.Values) ([]byte, error) {

	url := url.URL{
		Scheme: "https",
		Host:   c.host,
		Path:   path.Join(c.BasePath, metod),
	}

	request, err := http.NewRequest(http.MethodGet, url.String(), nil)
	if err != nil {
		return nil, err
	}

	request.URL.RawQuery = query.Encode()

	response, err := c.client.Do(request)
	if err != nil {
		return nil, err
	}

	body, err := io.ReadAll(response.Body)
	if err != nil {
		return nil, err
	}

	return body, nil
}
