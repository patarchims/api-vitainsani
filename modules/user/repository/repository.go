package repository

import (
	"context"
	"vincentcoreapi/modules/user"
	"vincentcoreapi/modules/user/entity"

	"gorm.io/gorm"
)

type userRepository struct {
	DB *gorm.DB
}

func NewUserRepository(db *gorm.DB) entity.UserRepository {
	return &userRepository{
		DB: db,
	}
}

func (ur *userRepository) GetByID(ctx context.Context, userID string) (user user.ApiUser, exist bool) {
	rs := ur.DB.WithContext(ctx).First(&user, userID).RowsAffected
	if rs > 0 {
		return user, true
	} else {
		return user, false
	}
}

func (ur *userRepository) GetByUser(ctx context.Context, userName string) (user user.ApiUser, exist bool) {
	rs := ur.DB.WithContext(ctx).Where("username = ? ", userName).First(&user).RowsAffected
	if rs > 0 {
		return user, true
	} else {
		return user, false
	}
}
