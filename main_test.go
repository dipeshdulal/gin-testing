package main

import (
	"encoding/json"
	"net/http"
	"net/http/httptest"
	"strings"
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

func TestDoubleRequestAPI(t *testing.T) {
	testPath := "/double"
	testMethod := "POST"

	t.Run("test-success-passing-proper-input", func(t *testing.T) {
		testRequest := DoubleRequest{Number: 2}.String()
		testResponse := DataResponse{Data: 4}.String()
		req := httptest.NewRequest(testMethod, testPath, strings.NewReader(testRequest))
		response := runRequest(req)
		assert.Equal(t, http.StatusOK, response.Code, "status code should be 200")
		assert.Equal(t, testResponse, response.Body.String(), "body is not same")
	})

	t.Run("test-failure-passing-bad-input", func(t *testing.T) {
		testRequest := DoubleRequest{Number: -1}.String()
		testResponse := ErrorResponse{Err: ErrNumberPositive.Error()}.String()
		req := httptest.NewRequest(testMethod, testPath, strings.NewReader(testRequest))
		response := runRequest(req)
		assert.Equal(t, http.StatusBadRequest, response.Code, "status not matching")
		assert.Equal(t, testResponse, response.Body.String(), "error return is not matching")
	})

	t.Run("test-failure-passing-no-input", func(t *testing.T) {
		req := httptest.NewRequest(testMethod, testPath, nil)
		response := runRequest(req)
		assert.Equal(t, http.StatusBadRequest, response.Code, "status not matching")
	})
}
