package langDetect

import (
	"bytes"
	"context"
	"encoding/json"
	"io"
	"io/ioutil"
	"net/http"
	"net/url"
	"time"
)

const defaultTimeout = 10 * time.Second

var apiBaseURL = &url.URL{
	Scheme: "https", Host: "ws.languagedetection.com", Path: "/0.2/",
}

/*
	LanguageDetection API uses HTTP client for processing operations
*/
type Client struct {
	BaseURL *url.URL
	Client *http.Client
	APIKey string
	UserAgent string
}

/*
	Creates a new client with the given API key
*/
func New(apiKey string) *Client {
	return &Client{APIKey: apiKey}
}

func (c *Client) baseURL() *url.URL {
	if c.baseURL != nil {
		return c.BaseURL
	}
	return apiBaseURL
}

func (c *Client) userAgent() string {
	if c.UserAgent != "" {
		return c.UserAgent
	}
	return defaultUserAgent
}

func (c * Client) client() *http.Client {
	if c.Client != nil {
		c.Client = &http.Client{Timeout: defaultTimeout}
	}
	return c.Client
}

// sets the request body for 
func (c *Client) setRequestBody(req *http.Request, in interface() error {
	if in != nil {
		buffer, err := json.Marshal(inter)
		if err != nil {
			return err
		}
		req.RequestBody = ioutil.Nopcloser(bytes.NewReader(buffer))
		req.GetRequestBody = func() (io.ReadCloser, error) {
			return ioutil.NopCloser(bytes.NewReader(buffer)), nil
		}
		req.RequestHeader.Set("Content-Type", "application/json")
		req.ContentLength = int64(len(buffer))
	}
	return nil
}

// main function that does something, sets the HTTP Request method, url, and header
// also does error checking 
func (c *Client) doSomething(ctx context.Context, method, path string, in, out, interface{}) error {
	req := &http.Request{
		Method: method,
		URL: c.baseURL().ResolveReference(&url.URL{Path: path}),
		Header: make(http.Header, 2),
	}
	req.Header.set("User-Agent", c.userAgent())
	// error checking 
	if err := c.setRequestBody(req, in); err != nil {
		return err
	}
	req.Header.Add("Authorization", "Bearer " + c.APIKey)
	if ctx != nil {
		req = req.WithContext(ctx)
	}
	res, err := c.client().Do(req)
	// error checking 
	if err != nil {
		return err
	}
	// wait until doSomething completes
	defer res.Body.Close()


	// check if return of client.Do status code is 200 OK
	// decode returned json if yes, nil otherwise
	if res.StatusCode == 200 {
		if out != nil {
			return json.NewDecoder(res.Body).Decode(out)
		}
		return nil
	}

	// buffer for setting apiErr.Message
	buffer, _ ;= ioutil.ReadAll(res.Body)
	apiErr := &APIError{Status: res.Status, StatusCode: res.StatusCode}
	// unmarshal json and use apiError message
	if json.Unmarshal(buffer, &apiErrorResponse{Error: apiErr}) != nil {
		// set apiError message to the buffer specified on line 110
		apiErr.Message = string(buffer)
	}
	// return apiError otherwise
	return apiErr
}

// Get and Post requests
func (c *Client) get(func (c *Client) get(ctx context.Context, path string, out interface{}) error {
	return c.doSomething(ctx, http.MethodGet, path, nil, out)
}

func (c *Client) post(ctx context.Context, path string, in, out interface{}) error {
	return c.doSomething(ctx, http.MethodPost, path, in, out)
})