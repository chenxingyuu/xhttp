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
}

func (r *Response) Body() (body []byte, err error) {
	body, err = io.ReadAll(r.RawResponse.Body)
	defer r.RawResponse.Body.Close()
	return body, err
}

func (r *Response) Json(v any) error {
	body, err := r.Body()
	if err != nil {
		return err
	}
	return json.Unmarshal(body, v)
}
