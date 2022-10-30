package main

import (
	"io"
	"net/http"
	"net/http/httptest"
	"testing"

	"github.com/stretchr/testify/assert"
)

func setupRouter(method string, url string, body io.Reader) *httptest.ResponseRecorder {
	router := GetRouter()

	w := httptest.NewRecorder()
	req, _ := http.NewRequest(method, url, body)
	router.ServeHTTP(w, req)
	return w
}

func Test_Router(t *testing.T) {
	w := setupRouter("GET", "/", nil)

	assert.Equal(t, http.StatusOK, w.Code)
	assert.Contains(t, w.Body.String(), "welcome !")
}
