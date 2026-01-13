package core_router

import (
	"api/auth"
	"api/core_db"
	"api/movies"
	"fmt"
	"net/http"
	"os"

	"github.com/gin-contrib/sessions"
	"github.com/gin-contrib/sessions/redis"
	"github.com/gin-gonic/gin"
	"gorm.io/driver/postgres"
	"gorm.io/gorm"
)

func Engine(dbDns, redisDns string) (*gin.Engine, error) {
	db, err := gorm.Open(postgres.Open(dbDns), &gorm.Config{})
	if err != nil {
		fmt.Fprintln(os.Stderr, "Error on initializing db: ", err)
		return nil, err
	}

	store, err := redis.NewStore(100, "tcp", redisDns, "", "", []byte("secret"))

	if err != nil {
		fmt.Fprintln(os.Stderr, "error while connecting redis : ", err)
		return nil, err
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

	return router, nil
}
