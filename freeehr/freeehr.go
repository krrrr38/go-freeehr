package freeehr

import (
	"bytes"
	"context"
	"encoding/json"
	"fmt"
	"io"
	"io/ioutil"
	"net/http"
	"net/url"
	"strconv"
)

const (
	defaultBaseURL = "http://api.freee.co.jp"
	userAgent      = "go-freeehr"

	headerRateLimitLimit     = "X-Ratelimit-Limit"
	headerRateLimitRemaining = "X-Ratelimit-Remaining"
	headerRateLimitReset     = "X-Ratelimit-Reset"
)

// TODO
type Client struct {
	client    *http.Client
	BaseURL   *url.URL
	UserAgent string

	Users     *UserService
	Companies *CompanyService
	Employees *EmployeeService

	common service
}

// TODO
type service struct {
	client *Client
}

// TODO
func NewClient(httpClient *http.Client) (*Client, error) {
	if httpClient == nil {
		return nil, fmt.Errorf("httpClient required")
	}
	baseURL, _ := url.Parse(defaultBaseURL)

	c := &Client{client: httpClient, BaseURL: baseURL, UserAgent: userAgent}
	c.common.client = c
	c.Users = (*UserService)(&c.common)
	c.Companies = (*CompanyService)(&c.common)
	c.Employees = (*EmployeeService)(&c.common)
	return c, nil
}

// TODO
func (c *Client) NewRequest(method, urlStr string, body interface{}) (*http.Request, error) {
	u, err := c.BaseURL.Parse(urlStr)
	if err != nil {
		return nil, err
	}

	var buf io.ReadWriter
	if body != nil {
		buf = new(bytes.Buffer)
		enc := json.NewEncoder(buf)
		enc.SetEscapeHTML(false)
		err := enc.Encode(body)
		if err != nil {
			return nil, err
		}
	}

	req, err := http.NewRequest(method, u.String(), buf)
	if err != nil {
		return nil, err
	}

	if body != nil {
		req.Header.Set("Content-Type", "application/json")
	}
	req.Header.Set("Accept", "application/json")
	if c.UserAgent != "" {
		req.Header.Set("User-Agent", c.UserAgent)
	}
	return req, nil
}

// TODO
type Response struct {
	*http.Response
	RateLimit *RateLimit
}

// TODO
type RateLimit struct {
	Limit     int
	Remaining int
	Reset     string
}

// TODO
func (c *Client) Do(ctx context.Context, req *http.Request, v interface{}) (*Response, error) {
	resp, err := c.client.Do(req.WithContext(ctx))
	if err != nil {
		return nil, err
	}

	defer func() {
		io.CopyN(ioutil.Discard, resp.Body, 512)
		resp.Body.Close()
	}()

	response := newResponse(resp)

	if v != nil {
		if w, ok := v.(io.Writer); ok {
			io.Copy(w, resp.Body)
		} else {
			err = json.NewDecoder(resp.Body).Decode(v)
			if err == io.EOF {
				err = nil
			}
		}
	}

	return response, err
}

// TODO
func newResponse(res *http.Response) *Response {
	response := &Response{Response: res}
	response.RateLimit = parseRateLimit(res)
	return response
}

// TODO
func parseRateLimit(r *http.Response) *RateLimit {
	var rateLimit RateLimit
	if limit := r.Header.Get(headerRateLimitLimit); limit != "" {
		rateLimit.Limit, _ = strconv.Atoi(limit)
	}
	if remaining := r.Header.Get(headerRateLimitRemaining); remaining != "" {
		rateLimit.Remaining, _ = strconv.Atoi(remaining)
	}
	if reset := r.Header.Get(headerRateLimitReset); reset != "" {
		rateLimit.Reset = reset
	}
	return &rateLimit
}
