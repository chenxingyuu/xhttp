package xhttp

import (
	"encoding/json"
	"io"
	"net/http"
)

type Response struct {
	Error       error
	RawResponse *http.Response
	StatusCode  int
	ResponseReader
}

type ResponseReader interface {
	Body() (body []byte, err error)
	String() (str string, err error)
	Json(v any) (err error)
	Map(m map[string]string, err error)
}

func (r *Response) Body() (body []byte, err error) {
	body, err = io.ReadAll(r.RawResponse.Body)
	defer r.RawResponse.Body.Close()
	return body, err
}

func (r *Response) String() (str string, err error) {
	body, err := r.Body()
	if err != nil {
		return "", err
	}
	return string(body), err
}

func (r *Response) Json(v any) error {
	body, err := r.Body()
	if err != nil {
		return err
	}
	return json.Unmarshal(body, v)
}

func (r *Response) Map() (m map[string]string, err error) {
	body, err := r.Body()
	if err != nil {
		return nil, err
	}
	err = json.Unmarshal(body, &m)
	if err != nil {
		return nil, err
	}
	return m, err
}
