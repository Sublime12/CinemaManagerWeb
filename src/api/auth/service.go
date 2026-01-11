package auth

import (
	"context"

	"gorm.io/gorm"
)

func GetUserByUsername(ctx context.Context, db *gorm.DB, username string) (User, error) {
	user, err := gorm.G[User](db).
		Where("username = ?", username).
		First(ctx)
	if err != nil {
		return User{}, err
	}
	return user, nil
}
