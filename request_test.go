package xhttp

import (
	"github.com/stretchr/testify/assert"
	"io"
	"net/http"
	"net/http/httptest"
	"testing"
)

func SendRequest(url, method string, body io.Reader) *Response {
	switch method {
	case "GET":
		return Get(url)
	case "OPTIONS":
		return Options(url)
	case "HEAD":
		return Head(url)
	case "POST":
		return Post(url, body)
	case "PUT":
		return Put(url, body)
	case "PATCH":
		return Patch(url, body)
	case "DELETE":
		return Delete(url)
	default:
		panic("unsupported method")
	}
}

func TestRequest(t *testing.T) {
	server := httptest.NewServer(http.HandlerFunc(func(rw http.ResponseWriter, req *http.Request) {
		_, _ = rw.Write([]byte(`OK`))
	}))
	defer server.Close()

	runTest := func(t *testing.T, method string) {
		res := SendRequest(server.URL, method, nil)
		assert.NoError(t, res.Error)
		if method != "HEAD" {
			body, _ := res.Body()
			assert.Equal(t, []byte(`OK`), body)
		}
	}

	for _, method := range []string{"GET", "OPTIONS", "HEAD", "POST", "PUT", "PATCH", "DELETE"} {
		t.Run("Test"+method, func(t *testing.T) {
			runTest(t, method)
		})
	}

}

func TestNewRequest(t *testing.T) {
	server := httptest.NewServer(http.HandlerFunc(func(rw http.ResponseWriter, req *http.Request) {
		_, _ = rw.Write([]byte(`OK`))
	}))
	defer server.Close()

	_, err := NewRequest(
		http.MethodGet,
		server.URL,
		nil,
		http.Header{
			"token": []string{"123"},
		},
		http.Cookie{
			Name:  "test",
			Value: "test",
		},
	)

	assert.NoError(t, err)

}
