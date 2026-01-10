package main

import (
	"api/auth"
	"api/movies"
	"fmt"
	"net/http"
	"os"

	"github.com/gin-gonic/gin"
	"gorm.io/driver/postgres"
	"gorm.io/gorm"
)

const dns = "host=db user=cinema_manager password=cinema_manager dbname=cinema_manager port=5432 sslmode=disable"

type Product struct {
	gorm.Model
	Code  string
	Price uint
}

func main() {

	db, err := gorm.Open(postgres.Open(dns), &gorm.Config{})
	if err != nil {
		fmt.Fprintln(os.Stderr, "Error on initializing db: ", err)
	}
	// ctx := context.Background()

	db.AutoMigrate(&auth.User{})

	// err = gorm.G[Product](db).Create(ctx, &Product{Code: "D42", Price: 100})
	// if err != nil {
	// 	fmt.Fprintln(os.Stderr, "Error on create product", err)
	// }

	router := gin.Default()
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
