package movies

type Movie struct {
	Id int `json:"id"`
	Name string `json:"name" binding:"required"`
	Description string `json:"description" binding:"required"`	
}

func NewMovie(Id int, Name, Description string) Movie {
	return Movie{
		Id: Id,
		Name: Name,
		Description: Description,
	}
}
