package movies

import "time"


type MovieResponse struct {
	ID          uint           `json:"id" binding:"required"`
	Name        string         `json:"name" binding:"required,min=1"`
	Description string         `json:"description" binding:"required,min=10"`
	PublishedAt time.Time      `json:"published_at" binding:"required"`
	Length      time.Duration  `json:"length" binding:"required,gt=0"` // keep as time.Duration
	Language    string         `json:"language" binding:"required,min=2"`
	Genres      []string       `json:"genres" binding:"required,dive,required"`
}

func MovieResponseFrom(m Movie) MovieResponse {
	return MovieResponse{
		ID:          m.ID,
		Name:        m.Name,
		Description: m.Description,
		PublishedAt: m.PublishedAt,
		Length:      m.Length,
		Language:    m.Language,
		Genres:      []string(m.Genres),
	}
}
