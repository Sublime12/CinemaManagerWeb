package auth_test

import (
	"api/core_router"
	"bytes"
	"encoding/json"
	"net/http"
	"net/http/httptest"
	"testing"

	"github.com/gin-gonic/gin"
	"github.com/stretchr/testify/assert"
	"github.com/stretchr/testify/require"
)

const testDbDns = "host=db-test user=cinema_manager password=cinema_manager dbname=cinema_manager port=5432 sslmode=disable"
const testRedisDns = "redis-test:6379"

func TestLoginRouteBadBody(t *testing.T) {
	router, err := core_router.Engine(testDbDns, testRedisDns)
	require.NoError(t, err)
	require.NotNil(t, router != nil)
	w := httptest.NewRecorder()
	req, _ := http.NewRequest("POST", "/api/login", nil)

	router.ServeHTTP(w, req)

	assert.Equal(t, http.StatusBadRequest, w.Code)
}

func TestLoginCorrectBodyInvalidCredentials(t *testing.T) {
	router, err := core_router.Engine(testDbDns, testRedisDns)
	require.NoError(t, err)
	require.NotNil(t, router != nil)
	body := gin.H {
		"username": "testuser",
		"password": "testpass",
	}
	
	jsonBody, err := json.Marshal(body)
	require.NoError(t, err)

	req, err := http.NewRequest("POST", "/api/login", bytes.NewBuffer(jsonBody))
	require.NoError(t, err)
	req.Header.Set("Content-Type", "application/json")

	w := httptest.NewRecorder()
	router.ServeHTTP(w, req)
	assert.Equal(t, http.StatusUnauthorized, w.Code, "res body: ", w.Body.String())
}
