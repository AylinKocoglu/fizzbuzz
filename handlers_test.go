package main

import (
	"net/http"
	"net/http/httptest"
	"testing"

	"github.com/stretchr/testify/assert"
)

func TestHandler(t *testing.T) {
	route := &Route{
		"fizzbuzz",
		"GET",
		"/fizzbuzz",
		FizzbuzzHandler,
	}
	servertest := httptest.NewServer(route.HandlerFunc)
	req, err := http.NewRequest("GET", servertest.URL+"/fizzbuzz?String1=fizz&String2=buzz&Int1=3&Int2=5&Limit=10", nil)
	assert.NoError(t, err)
	resp, err := http.DefaultClient.Do(req)
	assert.NoError(t, err)
	assert.Equal(t, 200, resp.StatusCode)

}
