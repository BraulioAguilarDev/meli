package melihttp

import (
	"bytes"
	"net/http"
	"time"
)

type Options struct {
	Endpoint string
	Method   string
	Headers  map[string]string
}

func NewClient() *Request {
	return &Request{
		Client: &http.Client{
			Timeout: time.Duration(10) * time.Second,
		},
	}
}

type Request struct {
	Client *http.Client
}

func (r *Request) MakeRequest(o *Options) (*http.Response, error) {
	req, err := http.NewRequest(o.Method, o.Endpoint, bytes.NewBuffer(nil))
	if err != nil {
		return nil, err
	}

	if len(o.Headers) > 0 {
		for k, v := range o.Headers {
			req.Header.Set(k, v)
		}
	}

	q := req.URL.Query()
	req.URL.RawQuery = q.Encode()

	res, err := r.Client.Do(req)
	if err != nil {
		return nil, err
	}

	return res, nil
}
