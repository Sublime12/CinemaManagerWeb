package auth_test

import (
	"api/auth"
	"api/core_router"
	auth_seeder "api/seeds/auth"
	"bytes"
	"encoding/json"
	"fmt"
	"log"
	"net/http"
	"net/http/httptest"
	"os"
	"testing"

	"github.com/gin-gonic/gin"
	"github.com/stretchr/testify/assert"
	"github.com/stretchr/testify/require"
	"gorm.io/driver/postgres"
	"gorm.io/gorm"
)

const testDbDns = "host=db-test user=cinema_manager password=cinema_manager dbname=cinema_manager port=5432 sslmode=disable"
const testRedisDns = "redis-test:6379"

var db *gorm.DB
func TestMain(m *testing.M) {
	gin.SetMode(gin.TestMode)
	db, err := gorm.Open(postgres.Open(testDbDns), &gorm.Config{})
	if err != nil {
		log.Fatal(err)
	}
	db.Migrator().DropTable(&auth.User{})
	err = db.AutoMigrate(&auth.User{})
	if err != nil {
		log.Fatal(err)
	}
	fmt.Println("BONJOUR")
	auth_seeder.SeedUsers(db)

	code := m.Run()
	os.Exit(code)
}

func TestLoginRouteBadBody(t *testing.T) {
	router, err := core_router.Engine(testDbDns, testRedisDns)

	require.NoError(t, err)
	require.NotNil(t, router != nil)
	w := httptest.NewRecorder()
	req, _ := http.NewRequest("POST", "/api/login", nil)

	router.ServeHTTP(w, req)

	require.Equal(t, http.StatusBadRequest, w.Code)
}

func TestLoginCorrectBodyInvalidCredentials(t *testing.T) {
	router, err := core_router.Engine(testDbDns, testRedisDns)
	require.NoError(t, err)
	require.NotNil(t, router)
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

func TestLoginValidCredentialsReturn200(t *testing.T) {
	router, err := core_router.Engine(testDbDns, testRedisDns)
	require.NoError(t, err)
	require.NotNil(t, router)

	body := gin.H {
		"username": "user01",
		"password": "password",
	}
	
	jsonBody, err := json.Marshal(body)
	require.NoError(t, err)

	req, err := http.NewRequest("POST", "/api/login", bytes.NewBuffer(jsonBody))
	require.NoError(t, err)
	req.Header.Set("Content-Type", "application/json")

	w := httptest.NewRecorder()
	router.ServeHTTP(w, req)
	assert.Equal(t, http.StatusOK, w.Code, "res body: ", w.Body.String())

}
