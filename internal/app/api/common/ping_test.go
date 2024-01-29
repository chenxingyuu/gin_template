package common

import (
	"github.com/gin-gonic/gin"
	"github.com/stretchr/testify/assert"
	"net/http"
	"net/http/httptest"
	"testing"
)

func TestPingHandler(t *testing.T) {
	// Set Gin to Test Mode
	gin.SetMode(gin.TestMode)

	// Setup Routes
	router := gin.Default()
	router.GET("/ping", PingHandler)

	// Create a request to pass to our handler.
	req, err := http.NewRequest(http.MethodGet, "/ping", nil)
	if err != nil {
		t.Fatalf("Couldn't create request: %v\n", err)
	}

	// We create a ResponseRecorder (which satisfies http.ResponseWriter) to record the response.
	rr := httptest.NewRecorder()

	// Process the request
	router.ServeHTTP(rr, req)

	// Check the status code is what we expect.
	assert.Equal(t, http.StatusOK, rr.Code)

	// Check the response body is what we expect.
	expected := `{"code":0,"data":null,"message":"ok"}`
	assert.JSONEq(t, expected, rr.Body.String())
}
