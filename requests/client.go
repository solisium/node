/*
 * Copyright (C) 2017 The "MysteriumNetwork/node" Authors.
 *
 * This program is free software: you can redistribute it and/or modify
 * it under the terms of the GNU General Public License as published by
 * the Free Software Foundation, either version 3 of the License, or
 * (at your option) any later version.
 *
 * This program is distributed in the hope that it will be useful,
 * but WITHOUT ANY WARRANTY; without even the implied warranty of
 * MERCHANTABILITY or FITNESS FOR A PARTICULAR PURPOSE.  See the
 * GNU General Public License for more details.
 *
 * You should have received a copy of the GNU General Public License
 * along with this program.  If not, see <http://www.gnu.org/licenses/>.
 */

package requests

import (
	"encoding/json"
	"fmt"
	"io"
	"io/ioutil"
	"net/http"
	"net/url"
	"sync"
	"time"

	"github.com/mysteriumnetwork/node/logconfig/httptrace"
	"github.com/mysteriumnetwork/node/metadata"
)

const (
	// DefaultTimeout is a default HTTP client timeout.
	DefaultTimeout = 20 * time.Second
)

// NewHTTPClientWithTransport creates a new HTTP client with custom transport.
func NewHTTPClientWithTransport(transport *http.Transport, timeout time.Duration) *HTTPClient {
	c := &HTTPClient{
		clientFactory: func() *http.Client {
			return &http.Client{
				Timeout:   timeout,
				Transport: setUserAgent(transport, getUserAgent()),
			}
		},
	}
	// Create initial clean before any HTTP request is made.
	c.client = c.clientFactory()

	return c
}

func setUserAgent(transport http.RoundTripper, userAgent string) http.RoundTripper {
	return &userAgenter{
		transport: transport,
		Agent:     userAgent,
	}
}

func getUserAgent() string {
	return fmt.Sprintf("Mysterium node(%v; https://mysterium.network)", metadata.VersionAsString())
}

type userAgenter struct {
	transport http.RoundTripper
	Agent     string
}

func (ua *userAgenter) RoundTrip(r *http.Request) (*http.Response, error) {
	r.Header.Set("User-Agent", ua.Agent)
	return ua.transport.RoundTrip(r)
}

// NewHTTPClient creates a new HTTP client.
func NewHTTPClient(srcIP string, timeout time.Duration) *HTTPClient {
	return NewHTTPClientWithTransport(NewTransport(NewDialer(srcIP).DialContext), timeout)
}

// HTTPClient describes a client for performing HTTP requests.
type HTTPClient struct {
	client        *http.Client
	clientMu      sync.Mutex
	clientFactory func() *http.Client
}

// Do sends an HTTP request and returns an HTTP response.
func (c *HTTPClient) Do(req *http.Request) (*http.Response, error) {
	return c.resolveClient().Do(req)
}

// DoRequest performs HTTP requests and parses error without returning response.
func (c *HTTPClient) DoRequest(req *http.Request) error {
	response, err := c.Do(req)
	if err != nil {
		return err
	}
	defer response.Body.Close()

	return ParseResponseError(response)
}

// DoRequestAndParseResponse performs HTTP requests and response from JSON.
func (c *HTTPClient) DoRequestAndParseResponse(req *http.Request, resp interface{}) error {
	response, err := c.Do(req)
	if err != nil {
		return err
	}
	defer response.Body.Close()

	httptrace.TraceRequestResponse(req, response)

	err = ParseResponseError(response)
	if err != nil {
		return err
	}

	return ParseResponseJSON(response, &resp)
}

func (c *HTTPClient) resolveClient() *http.Client {
	c.clientMu.Lock()
	defer c.clientMu.Unlock()
	if c.client != nil {
		return c.client
	}
	c.client = c.clientFactory()
	return c.client
}

// ParseResponseJSON parses http.Response into given struct.
func ParseResponseJSON(response *http.Response, dto interface{}) error {
	err := json.NewDecoder(response.Body).Decode(dto)
	if err == io.EOF {
		return nil
	}
	return err
}

// ErrorHTTP represent HTTP error with metadata.
type ErrorHTTP struct {
	Code     int
	Status   string
	Url      *url.URL
	response []byte
}

// Error returns string equivalent for error.
func (e *ErrorHTTP) Error() string {
	return fmt.Sprintf("server response invalid: %s (%s)%s", e.Status, e.Url, e.response)
}

// ParseResponseError parses http.Response error.
func ParseResponseError(response *http.Response) error {
	if response.StatusCode < 200 || response.StatusCode >= 300 {
		body, _ := ioutil.ReadAll(response.Body)
		return &ErrorHTTP{
			Code:     response.StatusCode,
			Status:   response.Status,
			Url:      response.Request.URL,
			response: body,
		}
	}

	return nil
}
