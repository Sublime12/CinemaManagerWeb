package movies

import (
	"api/auth"
	"api/core_db"
	"errors"
	"fmt"
	"net/http"
	"os"
	"path/filepath"
	"strconv"
	"time"

	"github.com/gin-gonic/gin"
	"gorm.io/gorm"
)


func MapMoviesRoutes(api *gin.RouterGroup) {
	router := api.Group("/movies")
	router.GET("", getMovies)
	router.GET("/:id", getMovie)

	authRouter := router.Group("")
	authRouter.Use(auth.AuthRequired)

	authRouter.POST("", postMovie)
	authRouter.POST("/upload", uploadPoster)
}


func getMovies(c *gin.Context) {
	ctx := c.Request.Context()
	db := core_db.GetSession(c)
	movies, err := gorm.G[Movie](db).Find(ctx)
	if err != nil {
		c.AbortWithError(http.StatusInternalServerError, err)
		return
	}

	moviesResponse := make([]MovieResponse, 0)
	for _, m := range movies {
		moviesResponse = append(moviesResponse, MovieResponseFrom(m))
	}
	c.JSON(http.StatusOK, moviesResponse)
}

func getMovie(c *gin.Context) {
	idInt, err := strconv.Atoi(c.Param("id"))
	id := uint(idInt)
	if err != nil {
		c.AbortWithError(http.StatusBadRequest, err)
		return
	}

	ctx := c.Request.Context()
	db := core_db.GetSession(c)
	movie, err := gorm.G[Movie](db).Where("id = ?", id).First(ctx)
	if errors.Is(err, gorm.ErrRecordNotFound) {
		c.AbortWithError(http.StatusNotFound, errors.New("movie not found"))
		return
	}
	if err != nil {
		c.AbortWithError(http.StatusInternalServerError, err)
	}


	c.JSON(http.StatusOK, MovieResponseFrom(movie))
}

func postMovie(c *gin.Context) {
	var movieRequest CreateMovieRequest
	err := c.ShouldBindJSON(&movieRequest)
	if err != nil {
		AbortWithError(c, err, http.StatusBadRequest)
		return
	}
	ctx := c.Request.Context()
	db := core_db.GetSession(c)

	movie := MovieFromRequest(movieRequest)
	err = gorm.G[Movie](db).Create(ctx, &movie)
	if err != nil {
		c.Error(err)
		AbortWithError(
			c, 
			errors.New("Internal server error"), 
			http.StatusInternalServerError,
		)
		return
	}
	
	c.JSON(http.StatusCreated, movie)
}

func uploadPoster(c *gin.Context) {
	file, err := c.FormFile("file")
	if err != nil {
		c.JSON(http.StatusBadRequest, gin.H{"error": "No file uploaded"})
		return
	}

	extension := filepath.Ext(file.Filename)
	filename := fmt.Sprintf("%d%s", time.Now().UnixNano(), extension)
	savePath := filepath.Join("uploads", "posters", filename)

	os.MkdirAll(filepath.Dir(savePath), os.ModePerm)

	if err := c.SaveUploadedFile(file, savePath); err != nil {
		c.JSON(http.StatusInternalServerError, gin.H{"error": "Failed to save file"})
		return
	}

	c.JSON(http.StatusOK, gin.H{
		"url": "/uploads/posters/" + filename,
	})
}

func AbortWithError(c *gin.Context, err error, status int) {
	c.Error(err)
	c.AbortWithStatusJSON(status, gin.H{
		"message": err.Error(),
	})
}
