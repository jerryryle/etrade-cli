package client

import (
	"github.com/stretchr/testify/mock"
	"io"
	"net/http"
	"strings"
)

type httpClientMock struct {
	mock.Mock
}

func (m *httpClientMock) Do(req *http.Request) (*http.Response, error) {
	args := m.Called(req.Method, req.URL.String())

	responseCode := args.Int(0)
	responseBody := args.String(1)
	err := args.Error(2)

	response := &http.Response{
		StatusCode: responseCode,
		Body:       io.NopCloser(strings.NewReader(responseBody)),
	}
	return response, err
}
