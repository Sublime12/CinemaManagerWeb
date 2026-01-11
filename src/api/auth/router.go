package auth

import (
	"api/core_db"
	"fmt"
	"net/http"
	"os"

	"github.com/gin-contrib/sessions"
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

	db := core_db.GetSession(c)
	ctx := c.Request.Context()
	user, err := GetUserByUsername(ctx, db, loginForm.Username)
	if err != nil {
		fmt.Fprintln(os.Stderr, "User not found", err)
		c.AbortWithError(http.StatusUnauthorized, err)
		return
	}

	if !user.IsPasswordCorrect(loginForm.Password) {
		fmt.Fprintln(os.Stderr, "Incorrect password ")
		c.JSON(http.StatusUnauthorized, gin.H{
			"error": "Authentification failed",
		})
		return
	}
	session := sessions.Default(c)
	session.Set("userId", user.ID)
	if err := session.Save(); err != nil {
		c.AbortWithError(http.StatusInternalServerError, err)
		return
	}

	c.JSON(http.StatusOK, gin.H{
		"message": "Successfully authentificated user",
	})
}

func logout(c *gin.Context) {
	fmt.Println("In logout")
}
