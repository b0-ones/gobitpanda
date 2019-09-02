package gobitpanda

import (
	"bytes"
	"encoding/json"
	"errors"
	"io"
	"io/ioutil"
	"net/http"
)

// NewClient returns a new Client struct
func NewClient(APIBase string, APIToken string) (*Client, error) {

	if APIBase == "" {
		return nil, errors.New("APIBase is required to create a Client")
	}

	return &Client{
		Client:   &http.Client{},
		APIBase:  APIBase,
		APIToken: APIToken,
	}, nil
}

// Send makes a request to the API and tries to unmarshal the response
func (c *Client) Send(req *http.Request, i interface{}) error {
	var (
		err  error
		resp *http.Response
		data []byte
	)

	req.Header.Set("Accept", "application/json")

	if req.Header.Get("Content-type") == "" {
		req.Header.Set("Content-type", "application/json")
	}

	resp, err = c.Client.Do(req)

	if err != nil {
		return err
	}
	defer resp.Body.Close()

	if resp.StatusCode < 200 || resp.StatusCode > 299 {
		errResp := &ErrorResponse{Response: resp}
		data, err = ioutil.ReadAll(resp.Body)

		if err == nil && len(data) > 0 {
			json.Unmarshal(data, errResp)
		}

		return errors.New(errResp.Error)
	}
	if i == nil {
		return nil
	}

	if w, ok := i.(io.Writer); ok {
		io.Copy(w, resp.Body)
		return nil
	}

	return json.NewDecoder(resp.Body).Decode(i)
}

// SendWithAuth makes a request to the API with an Auth header
func (c *Client) SendWithAuth(req *http.Request, v interface{}) error {

	if c.APIToken != "" {
		req.Header.Set("Authorization", "Bearer "+c.APIToken)
	} else {
		return errors.New("Client has no API Key, please first set an API Key")
	}

	return c.Send(req, v)
}

// NewRequest creates a new request and convert and add data as JSON if given
func (c *Client) NewRequest(method, url string, data interface{}) (*http.Request, error) {
	var buffer io.Reader
	if data != nil {
		b, err := json.Marshal(&data)
		if err != nil {
			return nil, err
		}
		buffer = bytes.NewBuffer(b)
	}
	return http.NewRequest(method, url, buffer)
}
