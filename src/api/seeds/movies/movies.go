package movies_seeder

import (
	m "api/movies"
	"context"
	"encoding/json"
	"fmt"
	"os"

	"gorm.io/gorm"
)

const moviesPath = "movies.json"
const threatresPath = "threatres.json"

func loadJSON[T any](path string, target *[]T) error {
	data, err := os.ReadFile(path)
	if err != nil {
		return err
	}
	return json.Unmarshal(data, target)
}

func SeedMovies(db *gorm.DB) error {
	ctx := context.Background()
	db.AutoMigrate(&m.Movie{})
	db.AutoMigrate(&m.Theatre{})

	var movies []m.Movie
	var threatres []m.Theatre
	err := loadJSON(moviesPath, &movies)
	if err != nil {
		return err
	}

	err = loadJSON(threatresPath, &threatres)
	if err != nil {
		return err
	}

	for _, movie := range movies {
		err := gorm.G[m.Movie](db).Create(ctx, &movie)
		if err != nil {
			fmt.Fprintln(os.Stderr, "Can't create movies ,", err)
			return err
		}
	}
	for _, theatre:= range threatres {
		err := gorm.G[m.Theatre](db).Create(ctx, &theatre)
		if err != nil {
			fmt.Fprintln(os.Stderr, "Can't create theatre", err)
			return err
		}
	}

	return nil
}
