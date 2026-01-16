package movies

import (
	"time"

	"gorm.io/gorm"
)

type Movie struct {
	Name string `json:"name" binding:"required"`
	Description string `json:"description" binding:"required"`	
	PublishedAt time.Time `json:"published_at" binding:"required"` 
	Length time.Duration `json:"length" binding:"required"`
	Language string `json:"language" binding:"required"`
	Genres []string `json:"genres" binding:"required"`
	gorm.Model
}

func NewMovie(
	Id int, 
	Name string, 
	Description string,
	PublishedAt	time.Time,
	Length time.Duration, 
	Language string,
	Genres []string,
) Movie {
	return Movie{
		Name: Name,
		Description: Description,
		PublishedAt: PublishedAt,
		Length: Length,
		Language: Language,
		Genres: Genres,
	}
}

type Theatre struct {
	Id int `json:"id"`
	Name string `json:"name" binding:"required"`
	Description *string `json:"description" binding:""`
	City string `json:"city" binding:"required"`
	Address string `json:"address" binding:"required"`
}

func NewThreatre(
	Id int, 
	Description *string,
	Name, City, Address string,
) Theatre {
	return Theatre{
		Id: Id,
		Description: Description,
		Name: Name,
		City: City,
		Address: Address,
	}
}
