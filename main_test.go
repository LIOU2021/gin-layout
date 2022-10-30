package main

import (
	"io"
	"net/http"
	"net/http/httptest"
	"testing"

	"github.com/gin-gonic/gin"
	"github.com/stretchr/testify/assert"
)

var (
	router *gin.Engine
	w      *httptest.ResponseRecorder
)

func init() {
	router = GetRouter()

	w = httptest.NewRecorder()
}

func api(method string, url string, body io.Reader) *httptest.ResponseRecorder {
	req, _ := http.NewRequest(method, url, body)
	router.ServeHTTP(w, req)
	return w
}

func Test_Welcome(t *testing.T) {
	w := api("GET", "/", nil)

	assert.Equal(t, http.StatusOK, w.Code)
	assert.Contains(t, w.Body.String(), "welcome !")
}

func Test_Welcome2(t *testing.T) {
	w := api("GET", "/", nil)

	assert.Equal(t, http.StatusOK, w.Code)
	assert.Contains(t, w.Body.String(), "welcome !")
}
