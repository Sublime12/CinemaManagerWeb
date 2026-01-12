package main

import (
	"api/auth"
	"api/movies"
	"fmt"
	"net/http"
	"os"

	"api/core_db"

	"github.com/gin-contrib/sessions"
	// "github.com/gin-contrib/sessions/cookie"
	"github.com/gin-contrib/sessions/redis"
	"github.com/gin-gonic/gin"
	"gorm.io/driver/postgres"
	"gorm.io/gorm"
	// "github.com/gin-contrib/sessions/filesystem"
)

const dns = "host=db user=cinema_manager password=cinema_manager dbname=cinema_manager port=5432 sslmode=disable"

func main() {
	db, err := gorm.Open(postgres.Open(dns), &gorm.Config{})
	if err != nil {
		fmt.Fprintln(os.Stderr, "Error on initializing db: ", err)
	}

	// store := cookie.NewStore([]byte("secret"))
	// const sessionPath = "/tmp/"
	// store := filesystem.NewStore(sessionPath, []byte("secret"))
	store, err := redis.NewStore(10, "tcp", "redis:6379", "", "", []byte("secret"))

	if err != nil {
		fmt.Fprintln(os.Stderr, "error while connecting redis : ", err)
		return
	}

	router := gin.Default()
	router.Use(core_db.DBMiddleware(db))
	router.Use(sessions.Sessions("mysessions", store))
	api := router.Group("/api")
	movies.MapMoviesRoutes(api)
	auth.MapAuthRoutes(api)

	api.GET("/", func(c *gin.Context) {
		c.JSON(http.StatusOK, gin.H{
			"message": "hello world",
		})
	})

	router.Run(":8080")
}
