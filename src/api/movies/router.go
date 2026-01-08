package movies

import (
	"fmt"
	"net/http"
	"strconv"

	"github.com/gin-gonic/gin"
)

var movies []Movie

func MapMoviesRoutes(api *gin.RouterGroup) {
	movies = make([]Movie, 0)
	for i := 0; i < 15; i++ {
		movie := NewMovie(i, 
			fmt.Sprintf("movie%d", i), 
			fmt.Sprintf("Description movie %d", i),
		)
		movies = append(movies, movie)
	}
	router := api.Group("/movies")	
	router.GET("", getMovies)
	router.GET("/:id", getMovie)
}


func getMovies(c *gin.Context) {
	c.JSON(http.StatusOK, movies)
}

func getMovie(c *gin.Context) {
	id, err := strconv.Atoi(c.Param("id"))
	if err != nil {
		c.AbortWithError(http.StatusBadRequest, err)
	}
	var movie *Movie = nil
	for _, m := range movies {
		if m.Id == id { movie = &m }
	}
	if movie == nil {
		c.AbortWithError(http.StatusNotFound, err)
	}
	c.JSON(http.StatusOK, *movie)
}
