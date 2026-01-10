package auth

import (
	"fmt"
	"net/http"

	"github.com/gin-gonic/gin"
)

func MapAuthRoutes(router *gin.RouterGroup) {
	router.POST("/login", login)
	router.POST("/logout", logout)
}

func login(c *gin.Context) {
	fmt.Println("In login")

	var loginForm LoginForm
	err := c.BindJSON(&loginForm)
	if err != nil {
		c.AbortWithError(
			http.StatusBadRequest,
			err,
		)
		return
	}


	c.JSON(http.StatusOK, "Success response")
}

func logout(c *gin.Context) {
	fmt.Println("In logout")
}
