package freeehr

import (
	"bytes"
	"encoding/json"
	"fmt"
	"io"
	"io/ioutil"
	"net/http"
	"net/url"
	"strconv"
	"strings"
)

const (
	defaultBaseURL = "https://api.freee.co.jp"
	userAgent      = "go-freeehr"

	headerRateLimitLimit     = "X-Ratelimit-Limit"
	headerRateLimitRemaining = "X-Ratelimit-Remaining"
	headerRateLimitReset     = "X-Ratelimit-Reset"
)

// Client is freee hr api client
type Client struct {
	client    *http.Client
	BaseURL   *url.URL
	UserAgent string

	Users     *UserService
	Companies *CompanyService
	Employees *EmployeeService

	common service
}

type service struct {
	client *Client
}

// NewClient generate freee hr api client
// generally httpClient requires OAuth2 context
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

// NewRequest build freee hr api http request
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

// Response represents raw freee hr api response
type Response struct {
	*http.Response
	RateLimit *RateLimit
}

// RateLimit represents freee hr api rate limit
type RateLimit struct {
	Limit     int
	Remaining int
	Reset     string
}

// Do execute http request call
func (c *Client) Do(req *http.Request, v interface{}) (*Response, error) {
	if DebugEnable() {
		fmt.Printf("> Send request: %v\n", req)
	}

	resp, err := c.client.Do(req)
	if err != nil {
		return nil, err
	}

	defer func() {
		io.CopyN(ioutil.Discard, resp.Body, 512)
		resp.Body.Close()
	}()

	err = checkResponse(resp)
	if err != nil {
		return nil, err
	}

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

func checkResponse(r *http.Response) error {
	if c := r.StatusCode; 200 <= c && c <= 299 {
		return nil
	}

	data, err := ioutil.ReadAll(r.Body)
	if err == nil && data != nil {
		return fmt.Errorf("api error: status=%v, body=%v", r.StatusCode, string(data))
	}

	return fmt.Errorf("api error: status=%v", r.StatusCode)
}

func newResponse(res *http.Response) *Response {
	response := &Response{Response: res}
	response.RateLimit = parseRateLimit(res)
	return response
}

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

// PagingOption manages freee hr api paging
type PagingOption struct {
	Per  int
	Page int
}

// AddPagingQueryParam handles url path with paging query parameters
func AddPagingQueryParam(path string, pagingOption *PagingOption) string {
	if pagingOption == nil {
		return path
	}

	if strings.Contains(path, "?") {
		return fmt.Sprintf("%s&per=%d&page=%d", path, pagingOption.Per, pagingOption.Page)
	}
	return fmt.Sprintf("%s?per=%d&page=%d", path, pagingOption.Per, pagingOption.Page)
}
