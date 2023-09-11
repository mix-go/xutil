package xhttp

import (
	"github.com/mix-go/xutil/xconv"
	"io"
	"net/http"
	"net/url"
	"strings"
)

type Response struct {
	*http.Response
	Body Body
}

type Body []byte

func (t Body) String() string {
	return xconv.BytesToString(t)
}

func newResponse(r *http.Response) *Response {
	if r == nil {
		return nil
	}
	resp := &Response{
		Response: r,
	}
	b, err := io.ReadAll(r.Body)
	defer r.Body.Close()
	if err != nil {
		return resp
	}
	resp.Body = b
	return resp
}

func Request(method string, u string, opts ...RequestOption) (*Response, error) {
	opt := getOptions(opts)
	cli := http.Client{
		Timeout: opt.Timeout,
	}
	URL, err := url.Parse(u)
	if err != nil {
		return nil, err
	}
	req := &http.Request{
		Method: method,
		URL:    URL,
		Header: opt.Header,
	}
	if opt.Body != nil {
		req.Body = io.NopCloser(strings.NewReader(opt.Body.String()))
	}
	r, err := cli.Do(req)
	resp := newResponse(r)
	return resp, err
}

func Do(req *http.Request, opts ...RequestOption) (*Response, error) {
	opt := getOptions(opts)
	cli := http.Client{
		Timeout: opt.Timeout,
	}
	r, err := cli.Do(req)
	resp := newResponse(r)
	return resp, err
}
