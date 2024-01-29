package response

import (
	"github.com/gin-gonic/gin"
	"github.com/stretchr/testify/assert"
	"net/http"
	"net/http/httptest"
	"testing"
)

func TestJsonResponse(t *testing.T) {
	w := httptest.NewRecorder()
	c, _ := gin.CreateTestContext(w)

	json(c, http.StatusNotFound, jsonOption{
		Code:    1,
		Message: "not found",
		Data:    nil,
	})

	assert.Equal(t, http.StatusNotFound, w.Code)

	expectedBody := `{"code":1,"message":"not found","data":null}`
	assert.JSONEq(t, expectedBody, w.Body.String())
}

func TestErrorJsonResponse(t *testing.T) {
	w := httptest.NewRecorder()
	c, _ := gin.CreateTestContext(w)

	Error(c, 1, "not found")
	assert.Equal(t, http.StatusBadRequest, w.Code)
	expectedBody := `{"code":1,"message":"unknown error: not found","data":null}`
	assert.JSONEq(t, expectedBody, w.Body.String())
}

func TestSuccessJsonResponse(t *testing.T) {
	w := httptest.NewRecorder()
	c, _ := gin.CreateTestContext(w)

	Success(c, "test data")

	assert.Equal(t, http.StatusOK, w.Code)

	expectedBody := `{"code":0,"message":"ok","data":"test data"}`
	assert.JSONEq(t, expectedBody, w.Body.String())
}
