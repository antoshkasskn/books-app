package http

import (
	"github.com/gin-gonic/gin"
	"github.com/stretchr/testify/assert"
	"net/http"
	"net/http/httptest"
	"testing"
)

func TestHealthCheck(t *testing.T) {
	gin.SetMode(gin.TestMode)

	router := gin.New()
	group := router.Group("/health")
	registerHealthCheck(group)

	req, _ := http.NewRequest(http.MethodGet, "/health/", nil)
	w := httptest.NewRecorder()

	router.ServeHTTP(w, req)

	assert.Equal(t, http.StatusOK, w.Code)

	assert.Equal(t, "HEALTHY", w.Body.String())
	assert.Equal(t, "text/plain", w.Header().Get("Content-Type"))
}
