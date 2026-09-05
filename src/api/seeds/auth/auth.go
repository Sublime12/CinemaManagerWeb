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
	db.Migrator().DropTable(&auth.User{})
	db.AutoMigrate(&auth.User{})
	ctx := context.Background()

	password, err := auth.HashPassword("password")
	if err != nil {
		fmt.Fprintln(os.Stderr, "Can't hash password", err)
	}

	test_password, err := auth.HashPassword("testpass")
	users := []auth.User{
		auth.NewUser("testuser", "testuser@test.com", test_password, "Test User", false),
		auth.NewUser("user01", "user01@test.com", password, "User One", true),
		auth.NewUser("user02", "user02@test.com", password, "User Two", false),
		auth.NewUser("user03", "user03@test.com", password, "User Three", false),
		auth.NewUser("user04", "user04@test.com", password, "User Four", false),
		auth.NewUser("user05", "user05@test.com", password, "User Five", false),
		auth.NewUser("user06", "user06@test.com", password, "User Six", false),
		auth.NewUser("user07", "user07@test.com", password, "User Seven", false),
		auth.NewUser("user08", "user08@test.com", password, "User Eight", false),
		auth.NewUser("user09", "user09@test.com", password, "User Nine", false),
		auth.NewUser("user10", "user10@test.com", password, "User Ten", false),
	}

	for _, u := range users {
		err := gorm.G[auth.User](db).Create(ctx, &u)
		if err != nil {
			fmt.Fprintln(os.Stderr, "can't create user ", u)
		}
	}
	fmt.Println("Users created")
}
