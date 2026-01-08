package movies

import (
	"fmt"
	"net/http"

	"github.com/gin-gonic/gin"
)

func MapMoviesRoutes(api *gin.RouterGroup) {
	router := api.Group("/movies")	
	router.GET("", getMovies)
}

func getMovies(c *gin.Context) {
	movies := make([]Movie, 0)
	
	for i := 0; i < 15; i++ {
		movie := NewMovie(i, 
			fmt.Sprintf("movie%d", i), 
			fmt.Sprintf("Description movie %d", i),
		)
		movies = append(movies, movie)
	}

	c.JSON(http.StatusOK, movies)
}
