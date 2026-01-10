package movies

import "time"

type Movie struct {
	Id int `json:"id"`
	Name string `json:"name" binding:"required"`
	Description string `json:"description" binding:"required"`	
	PublishedAt time.Time `json:"published_at" binding:"required"` 
	Length time.Duration `json:"length" binding:"required"`
	Language string `json:"language" binding:"required"`
	Genres []string `json:"genres" binding:"required"`
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
		Id: Id,
		Name: Name,
		Description: Description,
		PublishedAt: PublishedAt,
		Length: Length,
		Language: Language,
		Genres: Genres,
	}
}
