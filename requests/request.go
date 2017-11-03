package requests

import (
	"fmt"
	"io"
	"io/ioutil"
	"net/http"
	"testing"
)

// Request sends a request using the options which are passed in. The Options
// provide only some simple assertions i.e. status code. Everything is expected
// to be matched with in your own tests.
func Request(t *testing.T, opts Options) []byte {
	req, err := http.NewRequest(opts.method, opts.url, opts.body)
	if err != nil {
		t.Fatal(err)
	}
	req.Header = opts.headers

	client := http.Client{}
	resp, err := client.Do(req)
	if err != nil {
		t.Error(err)
	}
	defer resp.Body.Close()

	if !opts.statusCode.Match(resp.StatusCode) {
		t.Errorf("expected status code: %s, received: %d", opts.statusCode.String(), resp.StatusCode)
	}

	bytes, err := ioutil.ReadAll(resp.Body)
	if err != nil {
		t.Error(err)
	}
	return bytes
}

// Options defines different options that can be used in the request.
type Options struct {
	method     string
	url        string
	body       io.Reader
	headers    http.Header
	statusCode Range
}

// Option defines a function alias for adding an option to the options
type Option func(*Options)

// Build a series of options and put them on to the request options
// configurations
func Build(opts ...Option) Options {
	var res Options
	for _, opt := range opts {
		opt(&res)
	}
	return res
}

// WithHeaders options
func WithHeaders(h http.Header) Option {
	return func(o *Options) {
		o.headers = h
	}
}

// WithMethod options
func WithMethod(m string) Option {
	return func(o *Options) {
		o.method = m
	}
}

// WithURL options
func WithURL(u string) Option {
	return func(o *Options) {
		o.url = u
	}
}

// WithBody options
func WithBody(b io.Reader) Option {
	return func(o *Options) {
		o.body = b
	}
}

// WithStatusCode options
func WithStatusCode(code int) Option {
	return func(o *Options) {
		o.statusCode = Range{code, code}
	}
}

// Range holds a min and max value of a valid window
type Range struct {
	Min, Max int
}

// Match a value with in a range
// Return true if it's valid or not
func (r Range) Match(s int) bool {
	return s >= r.Min && s <= r.Max
}

func (r Range) String() string {
	if r.Min == r.Max {
		return fmt.Sprintf("%d", r.Min)
	}
	return fmt.Sprintf("%d-%d", r.Min, r.Max)
}
