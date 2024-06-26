package xhttp

import (
	"bytes"
	"encoding/json"
	"github.com/stretchr/testify/assert"
	"io"
	"net/http"
	"testing"
)

func TestResponse_Body(t *testing.T) {
	responseBody := "Hello, world!"
	resp := &Response{
		RawResponse: &http.Response{
			StatusCode: http.StatusOK,
			Body:       io.NopCloser(bytes.NewBufferString(responseBody)),
		},
	}

	body, err := resp.Body()
	assert.NoError(t, err)
	assert.Equal(t, body, []byte(responseBody))
}

func TestResponse_String(t *testing.T) {
	responseBody := "Hello, world!"
	resp := &Response{
		RawResponse: &http.Response{
			StatusCode: http.StatusOK,
			Body:       io.NopCloser(bytes.NewBufferString(responseBody)),
		},
	}

	body, err := resp.String()
	assert.NoError(t, err)
	assert.Equal(t, body, responseBody)
}

func TestResponse_Json(t *testing.T) {
	// 创建一个模拟的 JSON 响应
	type TestData struct {
		Message string `json:"message"`
	}
	expectedData := TestData{Message: "Hello, world!"}
	jsonBody, _ := json.Marshal(expectedData)

	resp := &Response{
		RawResponse: &http.Response{
			StatusCode: http.StatusOK,
			Body:       io.NopCloser(bytes.NewBuffer(jsonBody)),
		},
	}

	// 调用 Json 方法
	var actualData TestData
	err := resp.Json(&actualData)
	assert.NoError(t, err)
	assert.Equal(t, expectedData, actualData)
}

func TestResponse_Map(t *testing.T) {
	// 创建一个模拟的 JSON 响应
	testData := map[string]string{"message": "hello world"}
	jsonBody, _ := json.Marshal(testData)

	resp := &Response{
		RawResponse: &http.Response{
			StatusCode: http.StatusOK,
			Body:       io.NopCloser(bytes.NewBuffer(jsonBody)),
		},
	}

	// 调用 Json 方法
	actualData, err := resp.Map()
	assert.NoError(t, err)
	assert.Equal(t, testData, actualData)
}
