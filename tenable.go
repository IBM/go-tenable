/*
Copyright IBM Corp. 2022 All Rights Reserved.

SPDX-License-Identifier: Apache-2.0
*/

package tenable

import (
	"bytes"
	"context"
	"encoding/json"
	"fmt"
	"io"
	"net/http"
	"net/url"
	"os"
	"reflect"
	"strconv"
	"strings"

	"github.com/google/go-querystring/query"
)

// httpClient defines an interface for an http.Client implementation so that alternative
// http Clients can be passed in for making requests
type httpClient interface {
	Do(request *http.Request) (response *http.Response, err error)
}

// A Client manages communication with the Tenable API.
type Client struct {
	// HTTP client used to communicate with the API.
	client httpClient

	// Base URL for API requests.
	baseURL *url.URL

	apiKey string

	apiSecret string

	// Session storage if the user authenticates with a Session cookie
	//	session *Session

	// Services used for talking to different parts of the Tenable API.
	//	Authentication   *AuthenticationService
	CurrentUser *CurrentUserService

	Repository *RepositoryService
}

// NewClient returns a new Tenable API client.
// If a nil httpClient is provided, http.DefaultClient will be used.
// To use API methods which require authentication you can follow the preferred solution and
// provide an http.Client that will perform the authentication for you with OAuth and HTTP Basic (such as that provided by the golang.org/x/oauth2 library).
// As an alternative you can use Session Cookie based authentication provided by this package as well.
// See https://docs.tenable.com/tenablesc/api_best_practices/Content/ScApiBestPractices/APIKeyAuthorization.htm
// baseURL is the HTTP endpoint of your Tenable instance and should always be specified with a trailing slash.
func NewClient(httpClient httpClient, baseURL string) (*Client, error) {
	if httpClient == nil {
		httpClient = http.DefaultClient
	}

	// ensure the baseURL contains a trailing slash so that all paths are preserved in later calls
	if !strings.HasSuffix(baseURL, "/") {
		baseURL += "/"
	}

	parsedBaseURL, err := url.Parse(baseURL)
	if err != nil {
		return nil, err
	}
	apiKey := os.Getenv("SC05_ACCESS_KEY")
	apiSecret := os.Getenv("SC05_SECRET_KEY")
	if apiKey == "" || apiSecret == "" {
		return nil, fmt.Errorf("Requres env vars SC05_ACCESS_KEY and SC05_SECRET_KEY\n")
	}

	c := &Client{
		client:    httpClient,
		baseURL:   parsedBaseURL,
		apiKey:    apiKey,
		apiSecret: apiSecret,
	}
	//	c.Authentication = &AuthenticationService{client: c}
	c.CurrentUser = &CurrentUserService{client: c}
	c.Repository = &RepositoryService{client: c}
	return c, nil
}

// NewRawRequestWithContext creates an API request.
// A relative URL can be provided in urlStr, in which case it is resolved relative to the baseURL of the Client.
// Allows using an optional native io.Reader for sourcing the request body.
func (c *Client) NewRawRequestWithContext(ctx context.Context, method, urlStr string, body io.Reader) (*http.Request, error) {
	rel, err := url.Parse(urlStr)
	if err != nil {
		return nil, err
	}
	// Relative URLs should be specified without a preceding slash since baseURL will have the trailing slash
	rel.Path = strings.TrimLeft(rel.Path, "/")

	u := c.baseURL.ResolveReference(rel)

	req, err := newRequestWithContext(ctx, method, u.String(), body)
	if err != nil {
		return nil, err
	}
	req.Header.Set("x-apikey", fmt.Sprintf("accesskey=%s; secretkey=%s;", c.apiKey, c.apiSecret))
	req.Header.Set("Content-Type", "application/json")

	return req, nil
}

// NewRawRequest wraps NewRawRequestWithContext using the background context.
func (c *Client) NewRawRequest(method, urlStr string, body io.Reader) (*http.Request, error) {
	return c.NewRawRequestWithContext(context.Background(), method, urlStr, body)
}

func newRequestWithContext(ctx context.Context, method, url string, body io.Reader) (*http.Request, error) {
	return http.NewRequestWithContext(ctx, method, url, body)
}

// NewRequestWithContext creates an API request.
// A relative URL can be provided in urlStr, in which case it is resolved relative to the baseURL of the Client.
// If specified, the value pointed to by body is JSON encoded and included as the request body.
func (c *Client) NewRequestWithContext(ctx context.Context, method, urlStr string, body interface{}) (*http.Request, error) {
	rel, err := url.Parse(urlStr)
	if err != nil {
		return nil, err
	}
	// Relative URLs should be specified without a preceding slash since baseURL will have the trailing slash
	rel.Path = strings.TrimLeft(rel.Path, "/")

	u := c.baseURL.ResolveReference(rel)

	var buf io.ReadWriter
	if body != nil {
		buf = new(bytes.Buffer)
		err = json.NewEncoder(buf).Encode(body)
		if err != nil {
			return nil, err
		}
	}

	req, err := newRequestWithContext(ctx, method, u.String(), buf)
	if err != nil {
		return nil, err
	}
	req.Header.Set("Content-Type", "application/json")
	req.Header.Set("X-Apikey", fmt.Sprintf("accesskey=%s; secretkey=%s;", c.apiKey, c.apiSecret))

	return req, nil
}

// NewRequest wraps NewRequestWithContext using the background context.
func (c *Client) NewRequest(method, urlStr string, body interface{}) (*http.Request, error) {
	return c.NewRequestWithContext(context.Background(), method, urlStr, body)
}

// addOptions adds the parameters in opt as URL query parameters to s.  opt
// must be a struct whose fields may contain "url" tags.
func addOptions(s string, opt interface{}) (string, error) {
	v := reflect.ValueOf(opt)
	if v.Kind() == reflect.Ptr && v.IsNil() {
		return s, nil
	}

	u, err := url.Parse(s)
	if err != nil {
		return s, err
	}

	qs, err := query.Values(opt)
	if err != nil {
		return s, err
	}

	u.RawQuery = qs.Encode()
	return u.String(), nil
}

// CheckResponse checks the API response for errors, and returns them if present.
// A response is considered an error if it has a status code outside the 200 range.
// The caller is responsible to analyze the response body.
// The body can contain JSON (if the error is intended) or xml (sometimes Tenable just failes).
func CheckResponse(r *http.Response) error {
	if c := r.StatusCode; 200 <= c && c <= 299 {
		return nil
	}

	err := fmt.Errorf("request failed. Please analyze the request body for more details. Status code: %d", r.StatusCode)
	return err
}

// Response represents Tenable API response. It wraps http.Response returned from
// API and provides information about paging.
type Response struct {
	*http.Response

	StartAt    int
	MaxResults int
	Total      int
}

// Do sends an API request and returns the API response.
// The API response is JSON decoded and stored in the value pointed to by v, or returned as an error if an API error has occurred.
func (c *Client) Do(req *http.Request, v interface{}) (*Response, error) {
	httpResp, err := c.client.Do(req)
	if err != nil {
		return nil, err
	}

	err = CheckResponse(httpResp)
	if err != nil {
		// Even though there was an error, we still return the response
		// in case the caller wants to inspect it further
		return newResponse(httpResp, nil), err
	}

	if v != nil {
		// Open a NewDecoder and defer closing the reader only if there is a provided interface to decode to
		defer httpResp.Body.Close()
		// body, err := ioutil.ReadAll(httpResp.Body)
		// fmt.Println(string(body))
		err = json.NewDecoder(httpResp.Body).Decode(v)
		if err != nil {
			return nil, err
		}
	}

	resp := newResponse(httpResp, v)
	return resp, err
}

func newResponse(r *http.Response, v interface{}) *Response {
	resp := &Response{Response: r}
	//resp.populatePageValues(v)
	return resp
}

func interface2Int(v interface{}) (int, error) {
	switch v := v.(type) {
	case float64:
		return int(v), nil
	case string:
		c, err := strconv.Atoi(v)
		if err != nil {
			return 0, err
		}
		return c, nil
	default:
		return 0, fmt.Errorf("conversion to int from %T not supported", v)
	}
}
