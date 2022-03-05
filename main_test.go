package main

import (
	"encoding/json"
	"net/http"
	"net/http/httptest"
	"testing"

	"github.com/gin-gonic/gin"
	"github.com/stretchr/testify/assert"
)

func runRequest(req *http.Request) *httptest.ResponseRecorder {
	w := httptest.NewRecorder()
	getRouter().ServeHTTP(w, req)
	return w
}

func TestHealthCheckAPI(t *testing.T) {
	testPath := "/health-check"
	testResponse, _ := json.Marshal(gin.H{"data": "ok"})
	testHeader := "test-value"
	req := httptest.NewRequest("GET", testPath, nil)
	response := runRequest(req)
	assert.Equal(t, string(testResponse), response.Body.String(), "response not same")
	assert.Equal(t, http.StatusOK, response.Code, "status response is not same")
	assert.Equal(t, testHeader, response.Header().Get("x-test-header"), "header is not same")
}
