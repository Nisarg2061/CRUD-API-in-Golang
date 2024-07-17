package main

import (
	"net/http"
	"net/http/httptest"
	"testing"

	"github.com/stretchr/testify/assert"
)

func testGetMovies(t *testing.T){
  testServer := httptest.NewServer((http.HandlerFunc(getMovies)))
  defer testServer.Close()

  testclient := testServer.Client()
  res, err := testclient.Get(testServer.URL)
  if err != nil {
    t.Error(err)
  }
  assert.Equal(t, 200, res.StatusCode)
}
