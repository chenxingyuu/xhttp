package xhttp

import (
	"io"
	"net/http"
)

type Request struct {
	*http.Request
}

func NewRequest(method string, url string, body io.Reader, args ...any) (*Request, error) {
	request, err := http.NewRequest(method, url, body)

	// 设置请求头、Cookie等
	for _, arg := range args {
		switch v := arg.(type) {
		case http.Header:
			request.Header = v
		case http.Cookie:
			request.AddCookie(&v)
		}
	}

	return &Request{Request: request}, err
}

func (request *Request) Do() *Response {
	// 发送请求
	resp, err := Client.Do(request.Request)
	if err != nil {
		return &Response{Error: err}
	}

	return &Response{Error: err, RawResponse: resp, StatusCode: resp.StatusCode}
}

func Get(url string, args ...any) *Response {
	request, _ := NewRequest(http.MethodGet, url, nil, args)
	return request.Do()
}

func Head(url string, args ...any) *Response {
	request, _ := NewRequest(http.MethodHead, url, nil, args)
	return request.Do()
}

func Options(url string, args ...any) *Response {
	request, _ := NewRequest(http.MethodOptions, url, nil, args)
	return request.Do()
}

func Post(url string, Body io.Reader, args ...any) *Response {
	request, _ := NewRequest(http.MethodPost, url, Body, args)
	return request.Do()
}

func Put(url string, Body io.Reader, args ...any) *Response {
	request, _ := NewRequest(http.MethodPut, url, Body, args)
	return request.Do()
}

func Patch(url string, Body io.Reader, args ...any) *Response {
	request, _ := NewRequest(http.MethodPatch, url, Body, args)
	return request.Do()
}

func Delete(url string, args ...any) *Response {
	request, _ := NewRequest(http.MethodDelete, url, nil, args)
	return request.Do()
}
