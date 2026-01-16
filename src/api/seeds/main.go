package main

import (
	auth_seeder "api/seeds/auth"
	movies_seeder "api/seeds/movies"
	"fmt"
	"os"

	"gorm.io/driver/postgres"
	"gorm.io/gorm"
)

const dns = "host=db user=cinema_manager password=cinema_manager dbname=cinema_manager port=5432 sslmode=disable"
func main() {
	db, err := gorm.Open(postgres.Open(dns), &gorm.Config{})
	if err != nil {
		fmt.Fprintln(os.Stderr, "Error on initializing db: ", err)
	}
	auth_seeder.SeedUsers(db)
	movies_seeder.SeedMovies(db)
}
