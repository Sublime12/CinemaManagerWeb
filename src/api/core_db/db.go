package core_db

import (
	"github.com/gin-gonic/gin"
	"gorm.io/gorm"
)

func DBMiddleware(db *gorm.DB) gin.HandlerFunc {
	return func(c *gin.Context) {
		c.Set("db", db)
		c.Next()
	}
}

func GetSession(c *gin.Context) *gorm.DB {
	db := c.MustGet("db").(*gorm.DB)
	return db
}
