package movies

import (
	"errors"
	"fmt"
	"net/http"
	"strconv"
	"time"

	"github.com/gin-gonic/gin"
)

var movies []Movie

var movieDescription string = `Jake and Neytiri's family grapples with 
	grief after Neteyam's death, encountering a new, aggressive 
	Na'vi tribe, the Ash People, who are led 
	by the fiery Varang, as the conflict on Pandora escalates and a
	new moral focus emerges.`

func MapMoviesRoutes(api *gin.RouterGroup) {
	movies = make([]Movie, 0)
	for i := 0; i < 15; i++ {
		movie := NewMovie(i, 
			fmt.Sprintf("Avatar: Fire and Ash %dII", i), 
			movieDescription,
			time.Now(),
			time.Hour * 2 + (time.Minute * 17),
			"English",
			[]string {"Action", "Adventure", "Fiction"},
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
	idInt, err := strconv.Atoi(c.Param("id"))
	id := uint(idInt)
	if err != nil {
		c.AbortWithError(http.StatusBadRequest, err)
		return
	}
	var movie *Movie = nil
	for _, m := range movies {
		if m.ID == id { movie = &m }
	}
	if movie == nil {
		c.AbortWithError(http.StatusNotFound, errors.New("movie not found"))
		return
	}
	c.JSON(http.StatusOK, *movie)
}
