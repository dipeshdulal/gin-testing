## Gin Framework Integration Test

Experiment for testing gin framework request and responses. For testing we require gin router, test request and test response recorder so that test cases can be tested against.

#### Example:

```go
req := httptest.NewRequest("GET", "/health-check", nil)
w := httptest.NewRecorder()
getRouter().ServeHTTP(w, req)
assert.Equal(t, http.StatusOK, response.Code, "status response is not same")
```
