package auth_seeder

import (
	"api/auth"
	"context"
	"fmt"
	"os"

	"gorm.io/gorm"
)

// const dns = "host=db user=cinema_manager password=cinema_manager dbname=cinema_manager port=5432 sslmode=disable"

func SeedUsers(db *gorm.DB) {
	ctx := context.Background()

	password, err := auth.HashPassword("password")
	if err != nil {
		fmt.Fprintln(os.Stderr, "Can't hash password", err)
	}

	users := []auth.User{
		auth.NewUser("user01", "user01@test.com", password, "User One"),
		auth.NewUser("user02", "user02@test.com", password, "User Two"),
		auth.NewUser("user03", "user03@test.com", password, "User Three"),
		auth.NewUser("user04", "user04@test.com", password, "User Four"),
		auth.NewUser("user05", "user05@test.com", password, "User Five"),
		auth.NewUser("user06", "user06@test.com", password, "User Six"),
		auth.NewUser("user07", "user07@test.com", password, "User Seven"),
		auth.NewUser("user08", "user08@test.com", password, "User Eight"),
		auth.NewUser("user09", "user09@test.com", password, "User Nine"),
		auth.NewUser("user10", "user10@test.com", password, "User Ten"),
	}

	for _, u := range users {
		err := gorm.G[auth.User](db).Create(ctx, &u)
		if err != nil {
			fmt.Fprintln(os.Stderr, "can't create user ", u)
		}
	}
	fmt.Println("Users created")
}
