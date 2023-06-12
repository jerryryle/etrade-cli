package etradelib

import "net/http"

type RoundTripFunc func(req *http.Request) *http.Response

func (f RoundTripFunc) RoundTrip(req *http.Request) (*http.Response, error) {
	return f(req), nil
}

func NewHttpClientFake(fn RoundTripFunc) *http.Client {
	return &http.Client{
		Transport: fn,
	}
}
